package pkg

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type LineData struct {
	Name string `json:"Name"`
	Data []any  `json:"Data"`
}

// chartJSON is the top-level structure for the JSON data sent to the Python script.
type chartJSON struct {
	X     []string   `json:"x"`
	Datas []LineData `json:"datas"`
}

func WriteChartToFile(filename string, x []string, data LineData) error {
	// Prepare data for JSON marshaling
	// The python script expects a list of series, so we create a list with one item.
	chartData := chartJSON{
		X:     x,
		Datas: []LineData{data},
	}

	// Marshal data to JSON
	jsonData, err := json.Marshal(chartData)
	if err != nil {
		return err
	}

	// Execute Python script
	cmd := exec.Command("python3", "bench/pkg/plot.py", filename)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	errChan := make(chan error, 1)

	go func() {
		defer stdin.Close()

		if _, err := stdin.Write(jsonData); err != nil {
			errChan <- fmt.Errorf("failed to write JSON data to stdin: %w", err)
		}

		close(errChan)
	}()

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to execute python script: %w\nOutput:\n%s", err, string(output))
	}

	// Check for errors from the goroutine
	if writeErr := <-errChan; writeErr != nil {
		return writeErr
	}

	return nil
}
