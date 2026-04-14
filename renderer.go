package main

type ChartRenderer interface {
	RenderCandles(candles []Candle)

	DrawHorizontalLine(level float64, fromIndex, toIndex int, label string)

	DrawTrendLine(slope, intercept float64, fromIndex, toIndex int, label string)

	MarkBreakout(candleIndex int)

	DrawTargetLine(price float64)

	Export(filename string) error
}
