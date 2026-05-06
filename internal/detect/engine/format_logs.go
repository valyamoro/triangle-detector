package engine

import (
	"fmt"
	"strings"
)

func formatCheckTimingDebug(s CheckTimingDebugSnapshot) string {
	var b strings.Builder
	fmt.Fprintf(&b, "checkTimingAndHighs trace\n")
	fmt.Fprintf(&b, "firstTouchIdx=%d n=%d\n", s.FirstTouchIdx, s.N)
	fmt.Fprintf(&b, "highAboveThreshold=%s crashThreshold=%s\n", atrFmt(s.HighAboveThreshold), atrFmt(s.CrashThreshold))
	fmt.Fprintf(&b, "firstTouchMaxRatio=%s maxAllowedFirstTouchIdx=%s\n", atrFmt(s.FirstTouchMaxRatio), atrFmt(s.MaxFirstTouchIdx))
	fmt.Fprintf(&b, "\npreFirstTouch bars 0..%d:\n", s.FirstTouchIdx-1)
	for _, row := range s.BarChecks {
		fmt.Fprintf(&b, "i=%d High=%s Low=%s highOK=%v lowOK=%v", row.Index, atrFmt(row.High), atrFmt(row.Low), row.HighOK, row.LowOK)
		if row.FailHigh {
			fmt.Fprintf(&b, " FAIL_HIGH")
		}
		if row.FailLow {
			fmt.Fprintf(&b, " FAIL_LOW")
		}
		fmt.Fprintf(&b, "\n")
	}
	fmt.Fprintf(&b, "lastTouchIdx=%d minLastTouchIdx=%d ok=%v\n", s.LastTouchIdx, s.MinLastTouchIdx, s.LastTouchOK)
	return b.String()
}

func formatFindValleysDebug(s FindValleysDebugSnapshot) string {
	var b strings.Builder
	fmt.Fprintf(&b, "findValleysBetweenTouches trace\n")
	fmt.Fprintf(&b, "resistance touches: %d\n", len(s.Touches))
	for i, t := range s.Touches {
		fmt.Fprintf(&b, "  touch[%d] index=%d value=%s\n", i, t.Index, atrFmt(t.Value))
	}
	fmt.Fprintf(&b, "\nsegments:\n")
	for i, seg := range s.Segments {
		fmt.Fprintf(&b, "segment[%d] fromTouch=%d toTouch=%d start=%d end=%d skipped=%v\n",
			i, seg.FromTouch, seg.ToTouch, seg.Start, seg.End, seg.Skipped)
		if seg.SkipNote != "" {
			fmt.Fprintf(&b, "  note: %s\n", seg.SkipNote)
		}
		if seg.HasValley {
			fmt.Fprintf(&b, "  valley index=%d value=%s\n", seg.Valley.Index, atrFmt(seg.Valley.Value))
		}
	}
	fmt.Fprintf(&b, "\nvalleys count=%d\n", len(s.Valleys))
	for i, v := range s.Valleys {
		fmt.Fprintf(&b, "  [%d] index=%d value=%s\n", i, v.Index, atrFmt(v.Value))
	}
	return b.String()
}

func formatValidateValleysDebug(s ValidateValleysDebugSnapshot) string {
	var b strings.Builder
	fmt.Fprintf(&b, "validateValleys trace\n")
	fmt.Fprintf(&b, "avgPrice=%s resistance=%s firstVIdx=%d\n", atrFmt(s.AvgPrice), atrFmt(s.ResistanceLevel), s.FirstVIdx)
	fmt.Fprintf(&b, "allowedFlat=%s maxValleyDepth=%s\n", atrFmt(s.AllowedFlat), atrFmt(s.MaxValleyDepth))
	fmt.Fprintf(&b, "\npair rising checks:\n")
	for _, r := range s.PairChecks {
		fmt.Fprintf(&b, "i=%d prev=%s curr=%s minAllowed=%s ok=%v\n", r.I, atrFmt(r.PrevVal), atrFmt(r.CurrVal), atrFmt(r.MinAllowed), r.OK)
	}
	fmt.Fprintf(&b, "\ndepth vs resistance:\n")
	for _, r := range s.DepthChecks {
		fmt.Fprintf(&b, "index=%d value=%s minAllowed=%s ok=%v\n", r.Index, atrFmt(r.Value), atrFmt(r.MinAllowed), r.OK)
	}
	return b.String()
}

func formatFitSupportDebug(s FitSupportDebugSnapshot) string {
	var b strings.Builder
	fmt.Fprintf(&b, "fitSupportLine trace\n")
	fmt.Fprintf(&b, "slope=%s intercept=%s\n", atrFmt(s.Slope), atrFmt(s.Intercept))
	if s.SlopeRiseChecked {
		fmt.Fprintf(&b, "slopeRise=%s threshold=%s ok=%v\n", atrFmt(s.SlopeRise), atrFmt(s.SlopeRiseThreshold), s.SlopeRise >= s.SlopeRiseThreshold)
	}
	fmt.Fprintf(&b, "minRSquared=%s rSquared=%s checked=%v\n", atrFmt(s.MinRSquared), atrFmt(s.RSquared), s.RSquaredChecked)
	fmt.Fprintf(&b, "valleyDeviationMax=%s\n", atrFmt(s.ValleyDeviation))
	for _, v := range s.Valleys {
		fmt.Fprintf(&b, "  valley index=%d value=%s\n", v.Index, atrFmt(v.Value))
	}
	fmt.Fprintf(&b, "\npoint deviations from line:\n")
	for _, r := range s.DeviationRows {
		fmt.Fprintf(&b, "index=%d value=%s expected=%s relErr=%s maxOK=%s ok=%v\n",
			r.Index, atrFmt(r.Value), atrFmt(r.Expected), atrFmt(r.RelErr), atrFmt(r.MaxOK), r.OK)
	}
	return b.String()
}

func formatCheckGeometryDebug(s CheckGeometryDebugSnapshot) string {
	var b strings.Builder
	fmt.Fprintf(&b, "checkGeometry trace\n")
	fmt.Fprintf(&b, "patternStart=%d patternEnd=%d\n", s.PatternStart, s.PatternEnd)
	fmt.Fprintf(&b, "resistance=%s supportSlope=%s supportIntercept=%s\n", atrFmt(s.ResistanceLevel), atrFmt(s.SupportSlope), atrFmt(s.SupportIntercept))
	fmt.Fprintf(&b, "xIntersect=%s lastX=%s\n", atrFmt(s.XIntersect), atrFmt(s.LastX))
	fmt.Fprintf(&b, "ceilingTol=%s ceiling=%s ceilingScanEnd=%d\n", atrFmt(s.CeilingTol), atrFmt(s.Ceiling), s.CeilingEnd)
	fmt.Fprintf(&b, "floorTol=%s\n", atrFmt(s.FloorTol))
	fmt.Fprintf(&b, "heightAtStart=%s heightAtEnd=%s maxNarrowingRatio=%s\n", atrFmt(s.HeightAtStart), atrFmt(s.HeightAtEnd), atrFmt(s.MaxNarrowingRatio))
	fmt.Fprintf(&b, "minPatternHeight=%s minPatternWidth=%d maxApexFactor=%s\n", atrFmt(s.MinPatternHeight), s.MinPatternWidth, atrFmt(s.MaxApexFactor))
	fmt.Fprintf(&b, "lastResistanceIdx=%d lastValleyIdx=%d pEnd=%d patternWidth=%s\n",
		s.LastResistanceIdx, s.LastValleyIdx, s.PEnd, atrFmt(s.PatternWidth))
	if s.MaxResistanceTrailingGap > 0 {
		fmt.Fprintf(&b, "resistanceGap=%d gapRatio=%s maxGapRatio=%s\n",
			s.ResistanceGap, atrFmt(s.ResistanceGapRatio), atrFmt(s.MaxResistanceTrailingGap))
	}
	if s.StageNote != "" {
		fmt.Fprintf(&b, "stage: %s\n", s.StageNote)
	}
	if s.CeilingBreakBar >= 0 {
		fmt.Fprintf(&b, "ceilingBreakBar=%d\n", s.CeilingBreakBar)
	}
	if s.FloorBreakBar >= 0 {
		fmt.Fprintf(&b, "floorBreakBar=%d\n", s.FloorBreakBar)
	}
	if s.SupportCrossBar >= 0 {
		fmt.Fprintf(&b, "supportCrossBar=%d\n", s.SupportCrossBar)
	}
	return b.String()
}

func formatCheckVolumeDebug(s CheckVolumeDebugSnapshot) string {
	var b strings.Builder
	fmt.Fprintf(&b, "checkVolume trace\n")
	fmt.Fprintf(&b, "patternStart=%d pEnd=%d width=%d minWidth=%d\n", s.PatternStart, s.PEnd, s.Width, s.MinWidth)
	if s.Skipped {
		fmt.Fprintf(&b, "skipped: %s\n", s.SkipNote)
		return b.String()
	}
	fmt.Fprintf(&b, "pointCount=%d avgVol=%s volSlope=%s normalizedSlope=%s slopeMax=%s\n",
		s.PointCount, atrFmt(s.AvgVol), atrFmt(s.VolSlope), atrFmt(s.NormalizedSlope), atrFmt(s.SlopeMax))
	for _, pt := range s.Points {
		fmt.Fprintf(&b, "  i=%d volume=%s\n", pt.Index, atrFmt(pt.Value))
	}
	return b.String()
}

func collectCheckTimingDebug(ctx *pipeCtx) CheckTimingDebugSnapshot {
	p := ctx.p
	firstTouchIdx := ctx.resistanceTouchPoints[0].Index
	highAbove := ctx.resistanceLevel * (1 + ctx.vol*p.Timing.HighAboveVolMult)
	crash := ctx.resistanceLevel * (1 - max(p.Timing.CrashVolMin, ctx.vol*8))
	n := len(ctx.candles)
	s := CheckTimingDebugSnapshot{
		FirstTouchIdx:      firstTouchIdx,
		N:                  n,
		HighAboveThreshold: highAbove,
		CrashThreshold:     crash,
		FirstTouchMaxRatio: p.Timing.FirstTouchMaxRatio,
		MaxFirstTouchIdx:   float64(n) * p.Timing.FirstTouchMaxRatio,
	}
	for i := 0; i < firstTouchIdx; i++ {
		row := TimingBarCheckRow{
			Index: i,
			High:  ctx.candles[i].High,
			Low:   ctx.candles[i].Low,
		}
		row.HighOK = ctx.candles[i].High <= highAbove
		row.LowOK = ctx.candles[i].Low >= crash
		row.FailHigh = !row.HighOK
		row.FailLow = !row.LowOK
		s.BarChecks = append(s.BarChecks, row)
	}
	s.LastTouchIdx = ctx.resistanceTouchPoints[len(ctx.resistanceTouchPoints)-1].Index
	s.MinLastTouchIdx = int(float64(n) * p.Horizontal.MinLastTouchRatio)
	s.LastTouchOK = s.LastTouchIdx >= s.MinLastTouchIdx
	return s
}
