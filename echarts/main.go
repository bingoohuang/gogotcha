package main

import (
	"fmt"
	"os"

	"github.com/go-echarts/go-echarts/charts"
)

func main() {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-示例图"}, charts.ToolboxOpts{Show: true})
	bar.AddXAxis([]string{"衬衫", "牛仔裤"}).
		AddYAxis("商家A", []int{30, 20}).
		AddYAxis("商家B", []int{35, 14})

	f, err := os.Create("bar.html")
	if err != nil {
		fmt.Println(err)
	}

	err = bar.Render(f)
	if err != nil {
		fmt.Println(err)
	}
}
