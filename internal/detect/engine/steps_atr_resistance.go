package engine

import "github.com/gopherchan2006/go-triangle-detector/internal/detect/spec"

func stepCalcATR(ctx *pipeCtx) {
	sum := 0.0
	for _, c := range ctx.candles {
		sum += c.Close
	}
	ctx.avgPrice = sum / float64(len(ctx.candles))

	atrSnap := collectCalcATRDebug(ctx.candles)
	ctx.vol = atrSnap.ATR / ctx.avgPrice
	ctx.dbg.ATR.AvgPrice = ctx.avgPrice
	ctx.dbg.ATR.ATRValue = atrSnap.ATR
	ctx.dbg.ATR.Vol = ctx.vol
	ctx.dbg.Logs.CalcATR = formatCalcATRDebug(atrSnap)
}

func stepFindSwingHighs(ctx *pipeCtx) {
	snap := collectFindSwingHighsDebug(ctx.candles, ctx.p.Swing.Radius)
	ctx.swingHighs = snap.SwingHighs
	ctx.dbg.Logs.FindSwingHighs = formatFindSwingHighsDebug(snap)
	ctx.dbg.Swing.SwingHighsCount = len(ctx.swingHighs)
	if len(ctx.swingHighs) < 2 {
		ctx.reject(spec.ReasonFewSwingHighs)
	}
}

func stepFindResistance(ctx *pipeCtx) {
	snap := collectFindHorizontalResistanceDebug(ctx.candles, ctx.swingHighs, ctx.vol, ctx.p)
	ctx.dbg.Logs.FindHorizontalResistance = formatFindHorizontalResistanceDebug(snap)
	ctx.resistanceLevel = snap.Level
	ctx.resistanceTouches = snap.Touches
	ctx.resistanceTouchPoints = snap.TouchPoints
	ctx.dbg.Resistance.ResistanceLevel = snap.Level
	ctx.dbg.Resistance.ResistanceTouches = snap.Touches
	if snap.Touches < 3 {
		ctx.reject(spec.ReasonResistanceLt3Touches)
	}
}

func stepCheckTimingAndHighs(ctx *pipeCtx) {
	p := ctx.p
	td := collectCheckTimingDebug(ctx)
	ctx.dbg.Logs.CheckTimingAndHighs = formatCheckTimingDebug(td)

	firstTouchIdx := ctx.resistanceTouchPoints[0].Index
	highAboveThreshold := ctx.resistanceLevel * (1 + ctx.vol*p.Timing.HighAboveVolMult)
	crashThreshold := ctx.resistanceLevel * (1 - max(p.Timing.CrashVolMin, ctx.vol*8))
	ctx.dbg.Resistance.FirstTouchIdx = firstTouchIdx
	ctx.dbg.Resistance.HighAboveThreshold = highAboveThreshold
	ctx.dbg.Resistance.CrashThreshold = crashThreshold

	for i := 0; i < firstTouchIdx; i++ {
		if ctx.candles[i].High > highAboveThreshold {
			ctx.reject(spec.ReasonHighBeforeFirstTouch)
			return
		}
		if ctx.candles[i].Low < crashThreshold {
			ctx.reject(spec.ReasonCrashBeforeFirstTouch)
			return
		}
	}

	if float64(firstTouchIdx) > float64(len(ctx.candles))*p.Timing.FirstTouchMaxRatio {
		ctx.reject(spec.ReasonFirstTouchTooLate)
		return
	}

	lastTouchIdx := ctx.resistanceTouchPoints[len(ctx.resistanceTouchPoints)-1].Index
	if float64(lastTouchIdx) < float64(len(ctx.candles))*p.Horizontal.MinLastTouchRatio {
		ctx.reject(spec.ReasonResistanceLastTouchEarly)
	}
}

func stepFindValleys(ctx *pipeCtx) {
	valleys, snap := collectFindValleysDebug(ctx.candles, ctx.resistanceTouchPoints)
	ctx.valleys = valleys
	ctx.dbg.Logs.FindValleys = formatFindValleysDebug(snap)
	ctx.dbg.Support.ValleysCount = len(ctx.valleys)
	if len(ctx.valleys) < 2 {
		ctx.reject(spec.ReasonFewValleys)
	}
}
