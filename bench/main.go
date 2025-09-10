package main

import "log"

func main() {
	if err := WriteChartToFile("chart.html", []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}); err != nil {
		log.Printf("Error writing chart to file: %v", err)
	}
}
