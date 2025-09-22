package pkg

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
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

	// Create the 'tmp' directory if it doesn't exist
	if err := os.MkdirAll("tmp", 0o750); err != nil {
		return err
	}

	for _, data := range datas {
		ext := filepath.Ext(baseOutputFile)
		baseName := baseOutputFile[:len(baseOutputFile)-len(ext)]
		sanitizedSeriesName := reg.ReplaceAllString(data.Name, "_")
		outputFile := filepath.Join("tmp", baseName+"_"+sanitizedSeriesName+ext)

		if err := WriteChartToFile(outputFile, x, data); err != nil {
			// Continue to generate other charts even if one fails
			// log.Printf("Failed to generate chart for %s: %v", data.Name, err)
			return err // Or return immediately on first error
		}
	}

	if err := WriteResultsToFile(baseOutputFile, datas); err != nil {
		return err
	}

	return nil
}

func WriteResultsToFile(baseOutputFile string, datas []LineData) error {
	ext := filepath.Ext(baseOutputFile)
	baseName := baseOutputFile[:len(baseOutputFile)-len(ext)]
	outputFile := filepath.Join("tmp", baseName+"_results.json")

	// G304 (CWE-22): Ensure outputFile is within the 'tmp' directory.
	tmpDir, err := filepath.Abs("tmp")
	if err != nil {
		return fmt.Errorf("failed to get absolute path for tmp directory: %w", err)
	}

	absOutputFile, err := filepath.Abs(outputFile)
	if err != nil {
		return fmt.Errorf("failed to get absolute path for output file: %w", err)
	}

	if !strings.HasPrefix(absOutputFile, tmpDir) {
		return fmt.Errorf(
			"output file path %s is outside of the allowed directory %s",
			absOutputFile,
			tmpDir,
		)
	}

	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // For pretty printing

	if err := encoder.Encode(datas); err != nil {
		return err
	}

	return nil
}
