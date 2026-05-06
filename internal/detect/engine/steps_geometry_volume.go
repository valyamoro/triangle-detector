package engine

import "github.com/gopherchan2006/go-triangle-detector/internal/detect/spec"

func geometrySnapshotFromCtx(ctx *pipeCtx, pEnd, ceilingEnd int, note string, cBreak, fBreak, sCross int) CheckGeometryDebugSnapshot {
	p := ctx.p
	g := ctx.dbg.Geometry
	return CheckGeometryDebugSnapshot{
		PatternStart:      ctx.patternStart,
		PatternEnd:        ctx.patternEnd,
		ResistanceLevel:   ctx.resistanceLevel,
		SupportSlope:      ctx.supportSlope,
		SupportIntercept:  ctx.supportIntercept,
		XIntersect:        ctx.xIntersect,
		LastX:             ctx.lastX,
		CeilingTol:        g.CeilingTol,
		Ceiling:           g.Ceiling,
		CeilingEnd:        ceilingEnd,
		FloorTol:          g.FloorTol,
		HeightAtStart:     g.HeightAtStart,
		HeightAtEnd:       g.HeightAtEnd,
		MaxNarrowingRatio: p.Geometry.MaxNarrowingRatio,
		MinPatternHeight:  p.Geometry.MinPatternHeight,
		MinPatternWidth:   p.Geometry.MinPatternWidth,
		MaxApexFactor:     p.Geometry.MaxApexFactor,
		LastResistanceIdx: g.LastResistanceIdx,
		LastValleyIdx:     g.LastValleyIdx,
		PEnd:              pEnd,
		PatternWidth:      g.PatternWidth,
		StageNote:         note,
		CeilingBreakBar:   cBreak,
		FloorBreakBar:     fBreak,
		SupportCrossBar:   sCross,
	}
}

func stepCheckGeometry(ctx *pipeCtx) {
	p := ctx.p
	candles := ctx.candles
	valleys := ctx.valleys

	patternStart := ctx.resistanceTouchPoints[0].Index
	if valleys[0].Index < patternStart {
		patternStart = valleys[0].Index
	}
	patternEnd := len(candles) - 1
	ctx.patternStart = patternStart
	ctx.patternEnd = patternEnd
	ctx.dbg.Geometry.PatternStart = patternStart
	ctx.dbg.Geometry.PatternEnd = patternEnd

	xIntersect := (ctx.resistanceLevel - ctx.supportIntercept) / ctx.supportSlope
	lastX := float64(len(candles) - 1)
	ctx.xIntersect = xIntersect
	ctx.lastX = lastX
	ctx.dbg.Geometry.XIntersect = xIntersect
	ctx.dbg.Geometry.LastX = lastX
	if xIntersect <= lastX {
		ctx.dbg.Logs.CheckGeometry = formatCheckGeometryDebug(geometrySnapshotFromCtx(ctx, 0, patternEnd, "no convergence (apex not right of window)", -1, -1, -1))
		ctx.reject(spec.ReasonNoConvergence)
		return
	}

	ceilingTol := max(p.Geometry.CeilingTolMin, ctx.vol*0.7)
	ceiling := ctx.resistanceLevel * (1 + ceilingTol)
	ctx.dbg.Geometry.CeilingTol = ceilingTol
	ctx.dbg.Geometry.Ceiling = ceiling
	ceilingEnd := patternEnd
	if ceilingEnd == len(candles)-1 {
		ceilingEnd = patternEnd - 1
	}
	for i := patternStart; i <= ceilingEnd; i++ {
		if candles[i].High > ceiling {
			ctx.dbg.Logs.CheckGeometry = formatCheckGeometryDebug(geometrySnapshotFromCtx(ctx, 0, ceilingEnd, "breaks ceiling", i, -1, -1))
			ctx.reject(spec.ReasonBreaksCeiling)
			return
		}
	}

	floorTol := max(p.Geometry.FloorTolMin, ctx.vol*0.5)
	ctx.dbg.Geometry.FloorTol = floorTol
	for i := patternStart; i <= patternEnd; i++ {
		supportVal := ctx.supportSlope*float64(i) + ctx.supportIntercept
		if candles[i].Low < supportVal*(1-floorTol) {
			ctx.dbg.Logs.CheckGeometry = formatCheckGeometryDebug(geometrySnapshotFromCtx(ctx, 0, ceilingEnd, "breaks support floor", -1, i, -1))
			ctx.reject(spec.ReasonBreaksSupportFloor)
			return
		}
	}

	for i := patternStart; i <= patternEnd; i++ {
		if ctx.resistanceLevel <= ctx.supportSlope*float64(i)+ctx.supportIntercept {
			ctx.dbg.Logs.CheckGeometry = formatCheckGeometryDebug(geometrySnapshotFromCtx(ctx, 0, ceilingEnd, "support above resistance", -1, -1, i))
			ctx.reject(spec.ReasonSupportAboveResistance)
			return
		}
	}

	heightAtStart := ctx.resistanceLevel - (ctx.supportSlope*float64(patternStart) + ctx.supportIntercept)
	heightAtEnd := ctx.resistanceLevel - (ctx.supportSlope*float64(patternEnd) + ctx.supportIntercept)
	ctx.dbg.Geometry.HeightAtStart = heightAtStart
	ctx.dbg.Geometry.HeightAtEnd = heightAtEnd
	if heightAtEnd <= 0 || heightAtEnd >= heightAtStart*p.Geometry.MaxNarrowingRatio {
		ctx.dbg.Logs.CheckGeometry = formatCheckGeometryDebug(geometrySnapshotFromCtx(ctx, 0, ceilingEnd, "not narrowing", -1, -1, -1))
		ctx.reject(spec.ReasonNotNarrowing)
		return
	}

	if heightAtStart < ctx.resistanceLevel*p.Geometry.MinPatternHeight {
		ctx.dbg.Logs.CheckGeometry = formatCheckGeometryDebug(geometrySnapshotFromCtx(ctx, 0, ceilingEnd, "too flat", -1, -1, -1))
		ctx.reject(spec.ReasonTooFlat)
		return
	}

	lastResistanceIdx := ctx.resistanceTouchPoints[len(ctx.resistanceTouchPoints)-1].Index
	lastValleyIdx := valleys[len(valleys)-1].Index
	pEnd := lastResistanceIdx
	if lastValleyIdx > pEnd {
		pEnd = lastValleyIdx
	}
	ctx.dbg.Geometry.LastResistanceIdx = lastResistanceIdx
	ctx.dbg.Geometry.LastValleyIdx = lastValleyIdx
	ctx.dbg.Geometry.PEnd = pEnd
	if pEnd-patternStart < p.Geometry.MinPatternWidth {
		ctx.dbg.Logs.CheckGeometry = formatCheckGeometryDebug(geometrySnapshotFromCtx(ctx, pEnd, ceilingEnd, "too narrow", -1, -1, -1))
		ctx.reject(spec.ReasonTooNarrow)
		return
	}

	patternWidth := float64(pEnd - patternStart)
	ctx.dbg.Geometry.PatternWidth = patternWidth
	if xIntersect > lastX+patternWidth*p.Geometry.MaxApexFactor {
		ctx.dbg.Logs.CheckGeometry = formatCheckGeometryDebug(geometrySnapshotFromCtx(ctx, pEnd, ceilingEnd, "apex too far", -1, -1, -1))
		ctx.reject(spec.ReasonApexTooFar)
		return
	}

	ctx.dbg.Logs.CheckGeometry = formatCheckGeometryDebug(geometrySnapshotFromCtx(ctx, pEnd, ceilingEnd, "passed", -1, -1, -1))
}

func stepCheckVolume(ctx *pipeCtx) {
	p := ctx.p
	patternStart := ctx.patternStart
	pEnd := ctx.dbg.Geometry.PEnd
	snap := CheckVolumeDebugSnapshot{
		PatternStart: patternStart,
		PEnd:         pEnd,
		Width:        pEnd - patternStart,
		MinWidth:     p.VolumeDecl.VolDeclMinWidth,
	}
	if pEnd-patternStart < p.VolumeDecl.VolDeclMinWidth {
		snap.Skipped = true
		snap.SkipNote = "pattern width below VolDeclMinWidth"
		ctx.dbg.Logs.CheckVolume = formatCheckVolumeDebug(snap)
		return
	}

	volPoints := make([]SwingPoint, 0, pEnd-patternStart+1)
	volSum := 0.0
	for i := patternStart; i <= pEnd; i++ {
		volPoints = append(volPoints, SwingPoint{Index: i, Value: ctx.candles[i].Volume})
		volSum += ctx.candles[i].Volume
	}
	avgVol := volSum / float64(len(volPoints))
	volSlope, _ := linearRegression(volPoints)
	norm := 0.0
	if avgVol > 0 {
		norm = volSlope / avgVol
	}
	snap.Points = volPoints
	snap.PointCount = len(volPoints)
	snap.AvgVol = avgVol
	snap.VolSlope = volSlope
	snap.NormalizedSlope = norm
	snap.SlopeMax = p.VolumeDecl.VolDeclSlopeMax
	ctx.dbg.Logs.CheckVolume = formatCheckVolumeDebug(snap)
	if avgVol > 0 && norm > p.VolumeDecl.VolDeclSlopeMax {
		ctx.reject(spec.ReasonVolumeNotDeclining)
	}
}

