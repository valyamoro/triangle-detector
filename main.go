package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	jsonPath := "candles.json"
	if len(os.Args) > 1 {
		jsonPath = os.Args[1]
	}

	data, err := os.ReadFile(jsonPath)
	if err != nil {
		log.Fatalf("Error reading file %s: %v", jsonPath, err)
	}

	var candles []Candle
	if err := json.Unmarshal(data, &candles); err != nil {
		log.Fatalf("Error parsing JSON: %v\n"+
			"Expected format: [{\"open\":100,\"high\":105,\"low\":95,\"close\":102,"+
			"\"volume\":1000,\"timestamp\":\"2024-01-01T10:00:00Z\"}, ...]", err)
	}

	fmt.Printf("Loaded %d candles\n", len(candles))

	if len(candles) > 50 {
		candles = candles[len(candles)-50:]
		fmt.Printf("Working with the last 50 candles.\n")
	}

	result := DetectAscendingTriangle(candles, 15)

	renderer := NewEChartsRenderer()

	if err := RenderTriangleDetection(candles, result, renderer); err != nil {
		log.Fatalf("Rendering error: %v", err)
	}

	outputFile := "chart.html"
	if !result.Found {
		outputFile = "chart_no_pattern.html"
	}
	fmt.Printf("\nChart saved: %s\n", outputFile)
	fmt.Println("Open the file in your browser to view.")
}
