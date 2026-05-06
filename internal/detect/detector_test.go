package detect_test

import (
	"testing"

	"github.com/gopherchan2006/go-triangle-detector/internal/detect"
	"github.com/gopherchan2006/go-triangle-detector/internal/domain"
)

func makeCandle(close float64) domain.Candle {
	return domain.Candle{
		Open:   close,
		High:   close * 1.01,
		Low:    close * 0.99,
		Close:  close,
		Volume: 1000,
	}
}

func TestDetectAscendingTriangle_TooFewCandles(t *testing.T) {
	candles := make([]domain.Candle, 5)
	result := detect.DetectAscendingTriangle(candles)
	if result.Found {
		t.Error("expected not found with too few candles")
	}
}

func TestDetectAscendingTriangle_EmptyCandles(t *testing.T) {
	result := detect.DetectAscendingTriangle(nil)
	if result.Found {
		t.Error("expected not found with nil candles")
	}
}

func TestDetectAscendingTriangle_RejectReasonSet(t *testing.T) {

	candles := make([]domain.Candle, 50)
	for i := range candles {
		candles[i] = makeCandle(100.0)
	}
	result := detect.DetectAscendingTriangle(candles)
	if result.Found {
		t.Error("expected not found for flat candles")
	}
	if result.RejectReason == "" {
		t.Error("expected a reject reason to be set")
	}
}

func makeCandleWick(close, high, low float64) domain.Candle {
	return domain.Candle{
		Open:   close,
		High:   high,
		Low:    low,
		Close:  close,
		Volume: 1000,
	}
}

func TestRejectValleyFlatNotRising(t *testing.T) {
	candles := make([]domain.Candle, 50)
	resistance := 100.0
	for i := range candles {
		candles[i] = makeCandleWick(resistance-0.1, resistance, resistance-0.5)
	}
	candles[3] = makeCandleWick(resistance, resistance, resistance-0.5)
	candles[18] = makeCandleWick(resistance, resistance, resistance-0.5)
	candles[35] = makeCandleWick(resistance, resistance, resistance-0.5)

	result := detect.DetectAscendingTriangle(candles)
	if result.Found {
		t.Error("expected not found for flat resistance + flat support (no ascending)")
	}
}

func TestRejectSupportSlopeTooFlat(t *testing.T) {
	candles := make([]domain.Candle, 50)
	resistance := 100.0
	support := 98.0
	for i := range candles {
		v := support + float64(i)*0.000001
		candles[i] = makeCandleWick(resistance-0.1, resistance*(1+0.003), v*(1-0.001))
	}
	candles[3] = makeCandleWick(resistance, resistance, support)
	candles[18] = makeCandleWick(resistance, resistance, support+0.000018)
	candles[35] = makeCandleWick(resistance, resistance, support+0.000035)
	candles[45] = makeCandleWick(resistance, resistance, support+0.000045)

	result := detect.DetectAscendingTriangle(candles)
	if result.Found && result.SupportSlope < 0.0001 {
		t.Error("expected flat support slope to be rejected")
	}
}
