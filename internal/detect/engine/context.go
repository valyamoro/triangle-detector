package engine

import (
	"github.com/gopherchan2006/go-triangle-detector/internal/detect/spec"
	"github.com/gopherchan2006/go-triangle-detector/internal/domain"
)

type pipeCtx struct {
	candles []domain.Candle
	p       spec.Params
	counter spec.RejectCounter
	dbg     DebugInfo

	avgPrice              float64
	vol                   float64
	swingHighs            []SwingPoint
	resistanceLevel       float64
	resistanceTouches     int
	resistanceTouchPoints []SwingPoint
	valleys               []SwingPoint
	supportSlope          float64
	supportIntercept      float64
	patternStart          int
	patternEnd            int
	xIntersect            float64
	lastX                 float64

	rejected *Result
}

func (ctx *pipeCtx) reject(reason spec.RejectReason) {
	ctx.counter.Inc(reason)
	ctx.rejected = &Result{RejectReason: reason}
}

func (ctx *pipeCtx) volFloor(min float64, mult float64) float64 {
	return max(min, ctx.vol*mult)
}
