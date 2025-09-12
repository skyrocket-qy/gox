package pkg

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type LineData struct {
	Name string        `json:"Name"`
	Data []interface{} `json:"Data"`
}

// chartJSON is the top-level structure for the JSON data sent to the Python script.
type chartJSON struct {
	X     []string   `json:"x"`
	Datas []LineData `json:"datas"`
}

func WriteChartToFile(filename string, x []string, datas ...LineData) error {
	// Prepare data for JSON marshaling
	chartData := chartJSON{
		X:     x,
		Datas: datas,
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

	go func() {
		defer stdin.Close()
		stdin.Write(jsonData)
	}()

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to execute python script: %w\nOutput:\n%s", err, string(output))
	}

	return nil
}
