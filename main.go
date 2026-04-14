package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	symbol := flag.String("symbol", "", "Trading pair symbol, e.g. BTCUSDT")
	interval := flag.String("interval", "", "Candle interval, e.g. 15m")
	startDate := flag.String("start", "", "Start time in RFC3339 or YYYY-MM-DD")
	endDate := flag.String("end", "", "End time in RFC3339 or YYYY-MM-DD")
	flag.Parse()

	candles, err := LoadCandles(CandleRequestParams{
		Symbol:    *symbol,
		Interval:  *interval,
		StartTime: *startDate,
		EndTime:   *endDate,
	}, "candles.json")
	if err != nil {
		log.Fatalf("Failed to load candles: %v", err)
	}

	fmt.Printf("Loaded %d candles\n", len(candles))

	if len(candles) > 50 {
		candles = candles[len(candles)-50:]
		fmt.Println("Working with the last 50 candles.")
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
