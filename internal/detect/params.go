package detect

type Params struct {
	Swing      SwingParams
	Horizontal HorizontalParams
	Timing     TimingParams
	Support    SupportFitParams
	Geometry   GeometryParams
	VolumeDecl VolumeDeclParams
	Breakout   BreakoutParams
}

type SwingParams struct {
	Radius int
}

type HorizontalParams struct {
	VolTolerance         float64
	BreakoutTolerance    float64
	MinResistanceSpacing int
}

type TimingParams struct {
	FirstTouchMaxRatio float64
	HighAboveVolMult   float64
	CrashVolMin        float64
}

type SupportFitParams struct {
	MaxFirstValleyCrash float64
	AllowedFlatVolMult  float64
	FloorTolerance      float64
	MinRSquared         float64
	MaxValleyDepthMin   float64
	ValleyDeviationMin  float64
}

type GeometryParams struct {
	CeilingTolMin     float64
	FloorTolMin       float64
	MaxNarrowingRatio float64
	MinPatternHeight  float64
	MinPatternWidth   int
	MaxApexFactor     float64
}

type VolumeDeclParams struct {
	VolDeclSlopeMax float64
	VolDeclMinWidth int
}

type BreakoutParams struct {
	BreakoutConfirm float64
	VolAvgWindow    int
}

func DefaultParams() Params {
	return Params{
		Swing: SwingParams{Radius: 3},
		Horizontal: HorizontalParams{
			VolTolerance:         0.002,
			BreakoutTolerance:    0.005,
			MinResistanceSpacing: 5,
		},
		Timing: TimingParams{
			FirstTouchMaxRatio: 2.0 / 5.0,
			HighAboveVolMult:   0.5,
			CrashVolMin:        0.05,
		},
		Support: SupportFitParams{
			MaxFirstValleyCrash: 0.015,
			AllowedFlatVolMult:  1.5,
			FloorTolerance:      0.003,
			MinRSquared:         0.85,
			MaxValleyDepthMin:   0.015,
			ValleyDeviationMin:  0.0015,
		},
		Geometry: GeometryParams{
			CeilingTolMin:     0.002,
			FloorTolMin:       0.0015,
			MaxNarrowingRatio: 0.7,
			MinPatternHeight:  0.005,
			MinPatternWidth:   15,
			MaxApexFactor:     2.0,
		},
		VolumeDecl: VolumeDeclParams{
			VolDeclSlopeMax: 0.01,
			VolDeclMinWidth: 10,
		},
		Breakout: BreakoutParams{
			BreakoutConfirm: 0.005,
			VolAvgWindow:    20,
		},
	}
}
