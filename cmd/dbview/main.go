package main

import (
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/spf13/viper"

	"github.com/bingoohuang/sqlx"
	"github.com/gdamore/tcell"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rivo/tview"
	"github.com/spf13/pflag"
)

const (
	batchSize  = 80         // The number of rows loaded per batch.
	finderPage = "*finder*" // The name of the Finder page.
)

// nolint:gochecknoglobals
var (
	app         *tview.Application // The tview application.
	pages       *tview.Pages       // The application pages.
	finderFocus tview.Primitive    // The primitive in the Finder that last had focus.
)

// Main entry point.
func main() {
	pflag.StringP("ds", "d", "", `eg.
	user:pass@localhost
	user:pass@localhost:3306/dbname
	MYSQL_PWD=8BE4 mysql -h 127.0.0.1 -P 9633 -u root
	mysql -h 127.0.0.1 -P 9633 -u root -p8BE4
	mysql -h 127.0.0.1 -P 9633 -u root -p8BE4 -Dtest
	mysql -h127.0.0.1 -u root -p8BE4 -Dtest
	127.0.0.1:9633 root/8BE4
	127.0.0.1 root/8BE4
	127.0.0.1:9633 root/8BE4 db=test
	root:8BE4@tcp(127.0.0.1:9633)/?charset=utf8mb4&parseTime=true&loc=Local
`)

	pflag.Parse()

	viper.SetEnvPrefix("DBVIEW")
	viper.AutomaticEnv()
	_ = viper.BindPFlags(pflag.CommandLine)

	dataSourceName := viper.GetString("ds")
	ds := sqlx.CompatibleMySQLDs(dataSourceName)
	more := sqlx.NewSQLMore("mysql", ds)

	sqlx.DB = more.Open()

	if err := sqlx.CreateDao(&dao); err != nil {
		panic(err)
	}

	// Start the application.
	app = tview.NewApplication()

	finder()

	if err := app.Run(); err != nil {
		fmt.Printf("Error running application: %s\n", err)
	}
}

// nolint:gochecknoglobals
var dao mysqlSchemaDao

// nolint:lll
type mysqlSchemaDao struct {
	ShowDatabases    func() []MyDatabase                   `sql:"show databases"`
	GetTables        func(schema string) []MySQLTable      `sql:"SELECT * FROM information_schema.TABLES WHERE TABLE_SCHEMA = :1"`
	GetTableCols     func(schema, t string) []MyTableCol   `sql:"SELECT * FROM information_schema.COLUMNS WHERE TABLE_SCHEMA = :1 AND TABLE_NAME = :2 ORDER BY ORDINAL_POSITION"`
	GetMyConstraints func(schema, t string) []MyConstraint `sql:"SELECT * FROM information_schema.TABLE_CONSTRAINTS t JOIN information_schema.KEY_COLUMN_USAGE k USING(CONSTRAINT_NAME,TABLE_SCHEMA,TABLE_NAME) WHERE TABLE_SCHEMA = :1 AND TABLE_NAME = :2 ORDER BY ORDINAL_POSITION"`
}

// MyDatabase ...
type MyDatabase struct {
	Database string `name:"Database"`
}

// MySQLTable ...
type MySQLTable struct {
	Name    string `name:"TABLE_NAME"`
	Comment string `name:"TABLE_COMMENT"`
}

// MyConstraint ...
type MyConstraint struct {
	ColumnName     string `name:"COLUMN_NAME"`
	ConstraintType string `name:"CONSTRAINT_TYPE"`
}

// MyTableCol ...
type MyTableCol struct {
	ColumnName   string `name:"COLUMN_NAME"`
	Type         string `name:"COLUMN_TYPE"`
	Extra        string `name:"EXTRA"` // auto_increment
	Comment      string `name:"COLUMN_COMMENT"`
	DataType     string `name:"DATA_TYPE"`   // char
	ColumnType   string `name:"COLUMN_TYPE"` // char(20)
	MaxLength    int    `name:"CHARACTER_MAXIMUM_LENGTH"`
	Nullable     string `name:"IS_NULLABLE"` // YES NO
	Default      string `name:"COLUMN_DEFAULT"`
	CharacterSet string `name:"CHARACTER_SET_NAME"`

	NumericPrecision int `name:"NUMERIC_PRECISION"`
	NumericScale     int `name:"NUMERIC_SCALE"`
	OrdinalPosition  int `name:"ORDINAL_POSITION"`
}

// Sets up a "Finder" used to navigate the databases, tables, and columns.
// nolint:funlen
func finder() {
	// Create the basic objects.
	databases := tview.NewList()
	databases.ShowSecondaryText(false).SetBorder(true)
	databases.SetTitle("Databases")

	columns := tview.NewTable()
	columns.SetBorder(true).SetTitle("Columns")

	tables := tview.NewList()
	tables.SetDoneFunc(func() {
		tables.Clear()
		columns.Clear()
		app.SetFocus(databases)
	})
	tables.ShowSecondaryText(true).SetBorder(true).SetTitle("Tables")

	// Create the layout.
	flex := tview.NewFlex().
		AddItem(databases, 0, 1, true).
		AddItem(tables, 0, 1, false).
		AddItem(columns, 0, 3, false)

	schemas := dao.ShowDatabases()
	for _, _schema := range schemas {
		schema := _schema
		databases.AddItem(schema.Database, "", 0, func() {
			// A database was selected. Show all of its tables.
			columns.Clear()
			tables.Clear()

			schemaTables := dao.GetTables(schema.Database)
			for _, schemaTable := range schemaTables {
				tables.AddItem(schemaTable.Name, schemaTable.Comment, 0, nil)
			}

			app.SetFocus(tables)

			// When the user navigates to a table, show its columns.
			tables.SetChangedFunc(func(i int, tableName string, t string, s rune) {
				// A table was selected. Show its columns.
				columns.Clear()

				columns.SetCell(0, 0, &tview.TableCell{Text: "Name", Align: tview.AlignCenter, Color: tcell.ColorYellow}).
					SetCell(0, 1, &tview.TableCell{Text: "Type", Align: tview.AlignCenter, Color: tcell.ColorYellow}).
					SetCell(0, 2, &tview.TableCell{Text: "Null", Align: tview.AlignCenter, Color: tcell.ColorYellow}).
					SetCell(0, 3, &tview.TableCell{Text: "Default", Align: tview.AlignCenter, Color: tcell.ColorYellow}).
					SetCell(0, 4, &tview.TableCell{Text: "Comment", Align: tview.AlignCenter, Color: tcell.ColorYellow})

				tableCols := dao.GetTableCols(schema.Database, tableName)
				tableConstraints := dao.GetMyConstraints(schema.Database, tableName)
				tableConstraintsMap := make(map[string]MyConstraint)
				for _, tableConstraint := range tableConstraints {
					tableConstraintsMap[tableConstraint.ColumnName] = tableConstraint
				}

				for _, col := range tableCols {
					color := parseColor(tableConstraintsMap, col)

					columns.SetCell(col.OrdinalPosition, 0, &tview.TableCell{Text: col.ColumnName, Color: color}).
						SetCell(col.OrdinalPosition, 1, &tview.TableCell{Text: col.ColumnType, Color: color}).
						SetCell(col.OrdinalPosition, 2, &tview.TableCell{Text: col.Nullable, Align: tview.AlignRight, Color: color}).
						SetCell(col.OrdinalPosition, 3, &tview.TableCell{Text: col.Default, Align: tview.AlignRight, Color: color}).
						SetCell(col.OrdinalPosition, 4, &tview.TableCell{Text: col.Comment, Align: tview.AlignLeft, Color: color})
				}
			})

			tables.SetCurrentItem(0) // Trigger the initial selection.

			// When the user selects a table, show its content.
			tables.SetSelectedFunc(func(i int, tableName string, t string, s rune) {
				content(schema.Database, tableName)
			})
		})
	}

	// Set up the pages and show the Finder.
	pages = tview.NewPages().AddPage(finderPage, flex, true, true)
	app.SetRoot(pages, true)
}

func parseColor(tableConstraintsMap map[string]MyConstraint, col MyTableCol) tcell.Color {
	color := tcell.ColorWhite
	myConstraint, ok := tableConstraintsMap[col.ColumnName]

	if ok {
		color, ok = map[string]tcell.Color{
			"CHECK":       tcell.ColorGreen,
			"FOREIGN KEY": tcell.ColorDarkMagenta,
			"PRIMARY KEY": tcell.ColorRed,
			"UNIQUE":      tcell.ColorDarkCyan,
		}[myConstraint.ConstraintType]
	}

	if !ok {
		color = tcell.ColorWhite
	}

	return color
}

// Shows the contents of the given table.
// nolint:lll,funlen,gocognit
func content(schema, tableName string) {
	finderFocus = app.GetFocus()

	// If this page already exists, just show it.
	schemaTableName := schema + "." + tableName
	if pages.HasPage(schemaTableName) {
		pages.SwitchToPage(schemaTableName)
		return
	}

	// We display the data in a table embedded in a frame.
	table := tview.NewTable().
		//SetFixed(1, 0).
		SetSeparator(tview.Borders.Vertical).
		SetBordersColor(tcell.ColorYellow)
	frame := tview.NewFrame(table).
		SetBorders(0, 0, 0, 0, 0, 0)
	frame.SetBorder(true).
		SetTitle(fmt.Sprintf(`Contents of table "%s"`, schemaTableName))

	// How many rows does this table have?
	var rowCount int

	err := sqlx.DB.QueryRow(fmt.Sprintf("select count(*) from %s", schemaTableName)).Scan(&rowCount)
	if err != nil {
		panic(err)
	}

	// Load a batch of rows.
	loadRows := func(offset int) {
		rows, err := sqlx.DB.Query(fmt.Sprintf("select * from %s limit ? offset ?", schemaTableName), batchSize, offset)
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		// The first row in the table is the list of column names.
		columnNames, err := rows.Columns()
		if err != nil {
			panic(err)
		}

		for index, name := range columnNames {
			table.SetCell(0, index, &tview.TableCell{Text: name, Align: tview.AlignCenter, Color: tcell.ColorYellow})
		}

		// Read the rows.
		columns := make([]interface{}, len(columnNames))
		columnPointers := make([]interface{}, len(columns))

		for index := range columnPointers {
			columnPointers[index] = &columns[index]
		}

		for rows.Next() {
			// Read the columns.
			err := rows.Scan(columnPointers...)
			if err != nil {
				panic(err)
			}

			// Transfer them to the table.
			row := table.GetRowCount()

			for index, column := range columns {
				switch value := column.(type) {
				case int64:
					table.SetCell(row, index, &tview.TableCell{Text: strconv.Itoa(int(value)), Align: tview.AlignRight, Color: tcell.ColorDarkCyan})
				case float64:
					table.SetCell(row, index, &tview.TableCell{Text: strconv.FormatFloat(value, 'f', 2, 64), Align: tview.AlignRight, Color: tcell.ColorDarkCyan})
				case string:
					table.SetCellSimple(row, index, value)
				case time.Time:
					t := value.Format("2006-01-02")

					table.SetCell(row, index, &tview.TableCell{Text: t, Align: tview.AlignRight, Color: tcell.ColorDarkMagenta})
				case []uint8:
					str := make([]byte, len(value))
					copy(str, value)

					table.SetCell(row, index, &tview.TableCell{Text: string(str), Align: tview.AlignRight, Color: tcell.ColorGreen})
				case nil:
					table.SetCell(row, index, &tview.TableCell{Text: "NULL", Align: tview.AlignCenter, Color: tcell.ColorRed})
				default:
					// We've encountered a type that we don't know yet.
					t := reflect.TypeOf(value)
					str := "?nil?"

					if t != nil {
						str = "?" + t.String() + "?"
					}

					table.SetCellSimple(row, index, str)
				}
			}
		}

		if err := rows.Err(); err != nil {
			panic(err)
		}

		// Show how much we've loaded.
		frame.Clear()

		loadMore := ""
		if table.GetRowCount()-1 < rowCount {
			loadMore = " - press Enter to load more"
		}

		loadMore = fmt.Sprintf("Loaded %d of %d rows%s", table.GetRowCount()-1, rowCount, loadMore)
		frame.AddText(loadMore, false, tview.AlignCenter, tcell.ColorYellow)
	}

	// Load the first batch of rows.
	loadRows(0)

	// Handle key presses.
	table.SetDoneFunc(func(key tcell.Key) {
		switch key {
		case tcell.KeyEscape:
			// Go back to Finder.
			pages.SwitchToPage(finderPage)
			if finderFocus != nil {
				app.SetFocus(finderFocus)
			}
		case tcell.KeyEnter:
			// Load the next batch of rows.
			loadRows(table.GetRowCount() - 1)
			table.ScrollToEnd()
		}
	})

	// Add a new page and show it.
	pages.AddPage(schemaTableName, frame, true, true)
}
