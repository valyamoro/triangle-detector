package engine

import (
	"github.com/gopherchan2006/go-triangle-detector/internal/detect/spec"
	"github.com/gopherchan2006/go-triangle-detector/internal/domain"
)

type RunOpts struct {
	Params  spec.Params
	Counter spec.RejectCounter
}

type pipelineStep struct {
	fn func(*pipeCtx)
}

var pipelineSteps = []pipelineStep{
	{stepCalcATR},
	{stepFindSwingHighs},
	{stepFindResistance},
	{stepCheckTimingAndHighs},
	{stepFindValleys},
	{stepValidateValleys},
	{stepFitSupportLine},
	{stepCheckGeometry},
	{stepCheckVolume},
}

func Detect(candles []domain.Candle, ro RunOpts) Result {
	ctx := &pipeCtx{candles: candles, p: ro.Params, counter: ro.Counter}
	for _, step := range pipelineSteps {
		step.fn(ctx)
		if ctx.rejected != nil {
			return *ctx.rejected
		}
	}
	return buildDetectResult(ctx)
}

func buildDetectResult(ctx *pipeCtx) Result {
	p := ctx.p
	candles := ctx.candles
	n := len(candles)

	targetPrice := ctx.resistanceLevel + (ctx.resistanceLevel - ctx.valleys[0].Value)

	breakoutDetected := candles[n-1].Close > ctx.resistanceLevel*(1+p.Breakout.BreakoutConfirm)
	breakoutVolumeRatio := 0.0
	if breakoutDetected {
		volStart := max(n-p.Breakout.VolAvgWindow, 0)
		sum := 0.0
		count := 0
		for i := volStart; i < n-1; i++ {
			sum += candles[i].Volume
			count++
		}
		if count > 0 && sum > 0 {
			avgVol := sum / float64(count)
			breakoutVolumeRatio = candles[n-1].Volume / avgVol
		}
	}

	return Result{
		Found:                 true,
		ResistanceLevel:       ctx.resistanceLevel,
		ResistanceTouches:     ctx.resistanceTouches,
		ResistanceTouchPoints: ctx.resistanceTouchPoints,
		SupportSlope:          ctx.supportSlope,
		SupportIntercept:      ctx.supportIntercept,
		SupportTouchPoints:    ctx.valleys,
		Debug:                 ctx.dbg,
		TargetPrice:           targetPrice,
		BreakoutDetected:      breakoutDetected,
		BreakoutVolumeRatio:   breakoutVolumeRatio,
	}
}
