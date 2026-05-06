package detect

import (
	"github.com/gopherchan2006/go-triangle-detector/internal/domain"
)

func collectFindValleysDebug(candles []domain.Candle, touches []SwingPoint) ([]SwingPoint, FindValleysDebugSnapshot) {
	snap := FindValleysDebugSnapshot{Touches: append([]SwingPoint(nil), touches...)}
	var valleys []SwingPoint

	for i := 0; i < len(touches)-1; i++ {
		start := touches[i].Index + 1
		end := touches[i+1].Index
		seg := FindValleysSegmentDebug{
			FromTouch: touches[i].Index,
			ToTouch:   touches[i+1].Index,
			Start:     start,
			End:       end,
		}
		if end-start < 2 {
			seg.Skipped = true
			seg.SkipNote = "segment width < 2"
			snap.Segments = append(snap.Segments, seg)
			continue
		}
		v := findLowestLow(candles, start, end)
		seg.Valley = v
		seg.HasValley = true
		snap.Segments = append(snap.Segments, seg)
		valleys = append(valleys, v)
	}

	lastTouch := touches[len(touches)-1].Index
	seg := FindValleysSegmentDebug{
		FromTouch: lastTouch,
		ToTouch:   -1,
		Start:     lastTouch + 1,
		End:       len(candles),
	}
	if len(candles)-1-lastTouch >= 5 {
		v := findLowestLow(candles, lastTouch+1, len(candles))
		seg.Valley = v
		seg.HasValley = true
		snap.Segments = append(snap.Segments, seg)
		valleys = append(valleys, v)
	} else {
		seg.Skipped = true
		seg.SkipNote = "tail after last touch shorter than 5 bars"
		snap.Segments = append(snap.Segments, seg)
	}

	snap.Valleys = append([]SwingPoint(nil), valleys...)
	return valleys, snap
}

func findValleysBetweenTouches(candles []domain.Candle, touches []SwingPoint) []SwingPoint {
	v, _ := collectFindValleysDebug(candles, touches)
	return v
}

func findLowestLow(candles []domain.Candle, from, to int) SwingPoint {
	minIdx := from
	minVal := candles[from].Low
	for i := from + 1; i < to; i++ {
		if candles[i].Low < minVal {
			minVal = candles[i].Low
			minIdx = i
		}
	}
	return SwingPoint{Index: minIdx, Value: minVal}
}

func rSquared(points []SwingPoint, slope, intercept float64) float64 {
	n := float64(len(points))
	sumY := 0.0
	for _, p := range points {
		sumY += p.Value
	}
	meanY := sumY / n
	ssTot, ssRes := 0.0, 0.0
	for _, p := range points {
		predicted := slope*float64(p.Index) + intercept
		ssTot += (p.Value - meanY) * (p.Value - meanY)
		ssRes += (p.Value - predicted) * (p.Value - predicted)
	}
	if ssTot == 0 {
		return 1.0
	}
	return 1.0 - ssRes/ssTot
}

func linearRegression(points []SwingPoint) (slope, intercept float64) {
	n := float64(len(points))
	sumX, sumY, sumXY, sumX2 := 0.0, 0.0, 0.0, 0.0
	for _, p := range points {
		x := float64(p.Index)
		y := p.Value
		sumX += x
		sumY += y
		sumXY += x * y
		sumX2 += x * x
	}
	denom := n*sumX2 - sumX*sumX
	if denom == 0 {
		return 0, 0
	}
	slope = (n*sumXY - sumX*sumY) / denom
	intercept = (sumY - slope*sumX) / n
	return
}
