package main

import "fmt"

func RenderTriangleDetection(
	candles []Candle,
	result AscendingTriangleResult,
	renderer ChartRenderer,
	outputPath string,
) error {
	renderer.RenderCandles(candles)

	if !result.Found {
		fmt.Println("Pattern not found. Saving clean chart.")
		return renderer.Export(outputPath)
	}

	renderer.DrawHorizontalLine(
		result.ResistanceLevel,
		0,
		len(candles)-1,
		fmt.Sprintf("Resistance %.2f", result.ResistanceLevel),
	)

	renderer.DrawTrendLine(
		result.SupportSlope,
		result.SupportIntercept,
		0,
		len(candles)-1,
		"Support",
	)

	fmt.Println("Pattern found!")
	fmt.Printf("  Resistance : %.2f\n", result.ResistanceLevel)
	fmt.Printf("  Support slope : %.4f\n", result.SupportSlope)

	return renderer.Export(outputPath)
}
