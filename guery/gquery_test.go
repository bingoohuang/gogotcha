package guery

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/bingoohuang/gou/enc"
	"log"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

type PipeLine struct {
	Seq              string `col:"序号"`
	Name             string `col:"Pipeline名字"`
	State            string `col:"mainstem状态"`
	Delay            string `col:"延迟时间"`
	LastSyncTime     string `col:"最后同步时间"`
	LastPositionTime string `col:"最后位点时间"`
}

type StructCreator struct {
	SlicePtr       reflect.Value
	SliceValue     reflect.Value
	ItemStructType reflect.Type
	Columns        []itemStructField
	ColumnsMap     map[int]itemStructField
}

func (c *StructCreator) PrepareColumns(columns []string) {
	m := make(map[int]itemStructField)

	for _, structCol := range c.Columns {
		for j, col := range columns {
			if strings.Contains(structCol.ColumnName, col) {
				m[j] = structCol
				break
			}
		}
	}

	c.ColumnsMap = m
}

func (c *StructCreator) CreateSliceItem(columns []string) {
	v := reflect.New(c.ItemStructType).Elem()

	for j, col := range columns {
		if structCol, ok := c.ColumnsMap[j]; ok {
			v.Field(structCol.FiledIndex).Set(reflect.ValueOf(col))
		}
	}

	c.SliceValue = reflect.Append(c.SliceValue, v)
	c.SlicePtr.Elem().Set(c.SliceValue)
}

type itemStructField struct {
	FiledIndex int
	ColumnName string
}

func NewStructCreator(slicePtr interface{}) *StructCreator {
	SlicePtr := reflect.ValueOf(slicePtr)
	sliceValue := SlicePtr.Elem()

	s := &StructCreator{
		SlicePtr:       SlicePtr,
		SliceValue:     sliceValue,
		ItemStructType: sliceValue.Type().Elem(),
	}

	columns := make([]itemStructField, 0)
	for i := 0; i < s.ItemStructType.NumField(); i++ {
		fi := s.ItemStructType.Field(i)
		if fi.PkgPath != "" {
			continue
		}

		colName := fi.Tag.Get("col")
		if colName == "" {
			colName = fi.Name
		}

		columns = append(columns, itemStructField{FiledIndex: i, ColumnName: colName})
	}

	s.Columns = columns

	return s
}

func TestGquery(t *testing.T) {
	// Request the HTML page.
	res, err := http.Get("http://127.0.0.1:2901/pipeline_list.htm?channelId=1")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	pipelines := make([]PipeLine, 0)
	creator := NewStructCreator(&pipelines)

	doc.Find("table.list tr").Each(func(i int, s *goquery.Selection) {
		columns := make([]string, 0)
		s.Children().Each(func(j int, s *goquery.Selection) {
			columns = append(columns, strings.TrimSpace(s.Text()))
		})

		if i == 0 {
			creator.PrepareColumns(columns)
		} else {
			creator.CreateSliceItem(columns)
		}
	})

	fmt.Println(enc.JSONPretty(pipelines))
	/*
	   [
	      	{
	      		"Seq": "3",
	      		"Name": "pipeb",
	      		"State": "工作中",
	      		"Delay": "1.213 s",
	      		"LastSyncTime": "2020-04-14 08:27:35",
	      		"LastPositionTime": "2020-04-17 07:22:04"
	      	},
	      	{
	      		"Seq": "1",
	      		"Name": "pipea",
	      		"State": "工作中",
	      		"Delay": "0.75 s",
	      		"LastSyncTime": "2020-04-17 05:32:23",
	      		"LastPositionTime": "2020-04-17 07:22:48"
	      	}
	   ]
	*/
}
