package pkg

import (
	"strconv"

	"github.com/go-echarts/go-echarts/v2/opts"
)

func toLineData[T ~int | ~int32 | ~int64 | ~uint | ~uint32 | ~uint64](data []T) []opts.LineData {
	ret := make([]opts.LineData, 0, len(data))
	for _, v := range data {
		ret = append(ret, opts.LineData{Value: v})
	}

	return ret
}

func Bench(fn func(), repeat int, outputFile string) error {
	x := make([]string, 0, repeat)
	for i := 1; i <= repeat; i++ {
		x = append(x, strconv.Itoa(i))
	}

	datas := []LineData{
		{"Timings (ns)", toLineData(CollectTimings(fn, repeat))},
		{"Heap Delta (bytes)", toLineData(CollectMems(fn, repeat))},
		{"NumGCs", toLineData(CollectNumGCs(fn, repeat))},
		{"Total Alloc (bytes)", toLineData(CollectAllocs(fn, repeat))},
		{"PauseNs", toLineData(CollectPauseNs(fn, repeat))},
		{"Goroutines", toLineData(CollectGoroutines(fn, repeat))},
	}

	return WriteChartToFile(outputFile, x, datas...)
}
