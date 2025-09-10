package main

import (
	"math/rand"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

// generate random data for line chart.
func generateLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for range 7 {
		items = append(items, opts.LineData{Value: rand.Intn(300)})
	}

	return items
}

type LineData struct {
	Name string
	Data []opts.LineData
}

func WriteChartToFile(filename string, x []string, datas ...LineData) error {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Line example in Westeros theme",
			Subtitle: "Line chart rendered by the http server this time",
		}))

	line.SetXAxis(x).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(true)}))

	for _, data := range datas {
		line.AddSeries(data.Name, data.Data)
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return line.Render(f)
}
