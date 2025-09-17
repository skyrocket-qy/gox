package pkg

import (
	"path/filepath"
	"regexp"
	"strconv"
)

func toInterfaceSlice[T ~int | ~int32 | ~int64 | ~uint | ~uint32 | ~uint64](data []T) []any {
	ret := make([]any, len(data))
	for i, v := range data {
		ret[i] = v
	}

	return ret
}

func Bench(fn func(), repeat int, baseOutputFile string) error {
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

	// Not safe for concurrent execution
	reg := regexp.MustCompile("[^a-zA-Z0-9]+")

	for _, data := range datas {
		ext := filepath.Ext(baseOutputFile)
		baseName := baseOutputFile[:len(baseOutputFile)-len(ext)]
		sanitizedSeriesName := reg.ReplaceAllString(data.Name, "_")
		outputFile := baseName + "_" + sanitizedSeriesName + ext

		if err := WriteChartToFile(outputFile, x, data); err != nil {
			// Continue to generate other charts even if one fails
			// log.Printf("Failed to generate chart for %s: %v", data.Name, err)
			return err // Or return immediately on first error
		}
	}

	return nil
}
