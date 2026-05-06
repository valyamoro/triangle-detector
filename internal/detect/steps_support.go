package detect

import "math"

func stepValidateValleys(ctx *pipeCtx) {
	p := ctx.p
	candles := ctx.candles
	valleys := ctx.valleys

	snap := ValidateValleysDebugSnapshot{
		AvgPrice:        ctx.avgPrice,
		ResistanceLevel: ctx.resistanceLevel,
		Valleys:         append([]SwingPoint(nil), valleys...),
	}

	firstVIdx := valleys[0].Index
	maxCrashRange := 0.0
	for k := firstVIdx - 2; k <= firstVIdx; k++ {
		if k >= 0 {
			r := (candles[k].High - candles[k].Low) / ctx.avgPrice
			if r > maxCrashRange {
				maxCrashRange = r
			}
		}
	}
	snap.FirstVIdx = firstVIdx
	snap.MaxCrashRange = maxCrashRange
	crashLimit := max(p.Support.MaxFirstValleyCrash, ctx.vol*4)
	snap.CrashLimit = crashLimit
	ctx.dbg.Support.FirstVIdx = firstVIdx
	ctx.dbg.Support.MaxCrashRange = maxCrashRange

	if maxCrashRange > crashLimit {
		ctx.dbg.Logs.ValidateValleys = formatValidateValleysDebug(snap)
		ctx.reject(ReasonFirstValleyCrash)
		return
	}

	allowedFlat := ctx.vol * p.Support.AllowedFlatVolMult
	snap.AllowedFlat = allowedFlat
	ctx.dbg.Support.AllowedFlat = allowedFlat
	for i := 1; i < len(valleys); i++ {
		minAl := valleys[i-1].Value * (1 - allowedFlat)
		snap.PairChecks = append(snap.PairChecks, ValleyPairCheckRow{
			I: i, PrevVal: valleys[i-1].Value, CurrVal: valleys[i].Value, MinAllowed: minAl,
			OK: valleys[i].Value >= minAl,
		})
		if valleys[i].Value < minAl {
			ctx.dbg.Logs.ValidateValleys = formatValidateValleysDebug(snap)
			ctx.reject(ReasonValleyNotRising)
			return
		}
	}

	floorTolerance := max(p.Support.FloorTolerance, ctx.vol)
	snap.FloorTolerance = floorTolerance
	for i := 1; i < len(valleys); i++ {
		floorMin := valleys[0].Value * (1 - floorTolerance)
		snap.FloorChecks = append(snap.FloorChecks, ValleyFloorCheckRow{
			I: i, CurrVal: valleys[i].Value, FloorMin: floorMin, OK: valleys[i].Value >= floorMin,
		})
		if valleys[i].Value < floorMin {
			ctx.dbg.Logs.ValidateValleys = formatValidateValleysDebug(snap)
			ctx.reject(ReasonFirstValleyNotFloor)
			return
		}
	}

	maxValleyDepth := max(p.Support.MaxValleyDepthMin, ctx.vol*5)
	snap.MaxValleyDepth = maxValleyDepth
	ctx.dbg.Support.MaxValleyDepth = maxValleyDepth
	for _, v := range valleys {
		minAl := ctx.resistanceLevel * (1 - maxValleyDepth)
		snap.DepthChecks = append(snap.DepthChecks, ValleyDepthCheckRow{
			Index: v.Index, Value: v.Value, MinAllowed: minAl, OK: v.Value >= minAl,
		})
		if v.Value < minAl {
			ctx.dbg.Logs.ValidateValleys = formatValidateValleysDebug(snap)
			ctx.reject(ReasonValleyTooDeep)
			return
		}
	}

	ctx.dbg.Logs.ValidateValleys = formatValidateValleysDebug(snap)
}

func stepFitSupportLine(ctx *pipeCtx) {
	p := ctx.p
	valleys := ctx.valleys
	snap := FitSupportDebugSnapshot{Valleys: append([]SwingPoint(nil), valleys...)}
	slope, intercept := linearRegression(valleys)
	snap.Slope = slope
	snap.Intercept = intercept
	ctx.supportSlope = slope
	ctx.supportIntercept = intercept
	ctx.dbg.Support.SupportSlope = slope
	ctx.dbg.Support.SupportIntercept = intercept

	if slope <= 0 {
		ctx.dbg.Logs.FitSupportLine = formatFitSupportDebug(snap)
		ctx.reject(ReasonNegativeSlope)
		return
	}

	if len(valleys) >= 3 {
		r2 := rSquared(valleys, slope, intercept)
		snap.RSquared = r2
		snap.MinRSquared = p.Support.MinRSquared
		snap.RSquaredChecked = true
		if r2 < p.Support.MinRSquared {
			ctx.dbg.Logs.FitSupportLine = formatFitSupportDebug(snap)
			ctx.reject(ReasonLowRSquared)
			return
		}
	}

	valleyDeviation := max(p.Support.ValleyDeviationMin, ctx.vol*1.0)
	snap.ValleyDeviation = valleyDeviation
	ctx.dbg.Support.ValleyDeviation = valleyDeviation
	for _, v := range valleys {
		expected := slope*float64(v.Index) + intercept
		relErr := 0.0
		ok := true
		if expected > 0 {
			relErr = math.Abs(v.Value-expected) / expected
			ok = relErr <= valleyDeviation
		}
		snap.DeviationRows = append(snap.DeviationRows, SupportDeviationRow{
			Index: v.Index, Value: v.Value, Expected: expected, RelErr: relErr, MaxOK: valleyDeviation, OK: ok,
		})
		if expected > 0 && relErr > valleyDeviation {
			ctx.dbg.Logs.FitSupportLine = formatFitSupportDebug(snap)
			ctx.reject(ReasonValleyOffSupportLine)
			return
		}
	}
	ctx.dbg.Logs.FitSupportLine = formatFitSupportDebug(snap)
}
