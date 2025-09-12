package pkg

import (
	"strconv"
)

func toInterfaceSlice[T ~int | ~int32 | ~int64 | ~uint | ~uint32 | ~uint64](data []T) []interface{} {
	ret := make([]interface{}, len(data))
	for i, v := range data {
		ret[i] = v
	}
	return ret
}

func Bench(fn func(), repeat int, outputFile string) error {
	x := make([]string, 0, repeat)
	for i := 1; i <= repeat; i++ {
		x = append(x, strconv.Itoa(i))
	}

	datas := []LineData{
		{"Timings (ns)", toInterfaceSlice(CollectTimings(fn, repeat))},
		{"Heap Delta (bytes)", toInterfaceSlice(CollectMems(fn, repeat))},
		{"NumGCs", toInterfaceSlice(CollectNumGCs(fn, repeat))},
		{"Total Alloc (bytes)", toInterfaceSlice(CollectAllocs(fn, repeat))},
		{"PauseNs", toInterfaceSlice(CollectPauseNs(fn, repeat))},
		{"Goroutines", toInterfaceSlice(CollectGoroutines(fn, repeat))},
	}

	return WriteChartToFile(outputFile, x, datas...)
}
