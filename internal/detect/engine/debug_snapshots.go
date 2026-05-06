package engine

type CalcATRBarTrace struct {
	Index                 int
	FirstBar              bool
	O, H, L, C, PrevClose float64

	HighLow  float64
	D1, D2   float64
	D1TookTR bool
	D2TookTR bool
	TR       float64
	SumTR    float64
}

type CalcATRDebugSnapshot struct {
	BarCount int
	Bars     []CalcATRBarTrace
	SumTR    float64
	ATR      float64
}

type SwingHighScanRow struct {
	Index       int
	High        float64
	IsSwingHigh bool
	BlockIndex  int
	BlockHigh   float64
}

type FindSwingHighsDebugSnapshot struct {
	Radius     int
	N          int
	Rows       []SwingHighScanRow
	SwingHighs []SwingPoint
}

type HorizontalResistanceGroupDebug struct {
	Points      []SwingPoint
	Sum         float64
	AvgAll      float64
	SpacedValid []SwingPoint
}

type FindHorizontalResistanceDebugSnapshot struct {
	Vol             float64
	VolTolParam     float64
	Tolerance       float64
	Breakout        float64
	MinSpacing      int
	HighsIn         []SwingPoint
	Groups          []HorizontalResistanceGroupDebug
	BestGroupIdx    int
	BestLevel       float64
	BestTouchPoints []SwingPoint
	FailReason      string
	FailPairIdx     int
	FailBar         int
	FailClose       float64
	FailLimit       float64
	Level           float64
	Touches         int
	TouchPoints     []SwingPoint
}

type FindValleysSegmentDebug struct {
	FromTouch int
	ToTouch   int
	Start     int
	End       int
	Skipped   bool
	SkipNote  string
	Valley    SwingPoint
	HasValley bool
}

type FindValleysDebugSnapshot struct {
	Touches  []SwingPoint
	Segments []FindValleysSegmentDebug
	Valleys  []SwingPoint
}

type CheckTimingDebugSnapshot struct {
	FirstTouchIdx       int
	N                   int
	HighAboveThreshold  float64
	CrashThreshold      float64
	FirstTouchMaxRatio  float64
	MaxFirstTouchIdx    float64
	PrecedingBars       int
	PreSlope            float64
	PrecedingChecked    bool
	BarChecks           []TimingBarCheckRow
}

type TimingBarCheckRow struct {
	Index       int
	High        float64
	Low         float64
	HighOK      bool
	LowOK       bool
	FailHigh    bool
	FailLow     bool
}

type ValidateValleysDebugSnapshot struct {
	AvgPrice           float64
	ResistanceLevel    float64
	FirstVIdx          int
	MaxCrashRange      float64
	CrashLimit         float64
	AllowedFlat        float64
	FloorTolerance     float64
	MaxValleyDepth     float64
	Valleys            []SwingPoint
	PairChecks         []ValleyPairCheckRow
	FloorChecks        []ValleyFloorCheckRow
	DepthChecks        []ValleyDepthCheckRow
}

type ValleyPairCheckRow struct {
	I            int
	PrevVal      float64
	CurrVal      float64
	MinAllowed   float64
	OK           bool
}

type ValleyFloorCheckRow struct {
	I           int
	CurrVal     float64
	FloorMin    float64
	OK          bool
}

type ValleyDepthCheckRow struct {
	Index       int
	Value       float64
	MinAllowed  float64
	OK          bool
}

type FitSupportDebugSnapshot struct {
	Valleys           []SwingPoint
	Slope             float64
	Intercept         float64
	MinRSquared       float64
	RSquared          float64
	RSquaredChecked   bool
	ValleyDeviation   float64
	DeviationRows     []SupportDeviationRow
}

type SupportDeviationRow struct {
	Index    int
	Value    float64
	Expected float64
	RelErr   float64
	MaxOK    float64
	OK       bool
}

type CheckGeometryDebugSnapshot struct {
	PatternStart      int
	PatternEnd        int
	ResistanceLevel   float64
	SupportSlope      float64
	SupportIntercept  float64
	XIntersect        float64
	LastX             float64
	CeilingTol        float64
	Ceiling           float64
	CeilingEnd        int
	FloorTol          float64
	HeightAtStart     float64
	HeightAtEnd       float64
	MaxNarrowingRatio float64
	MinPatternHeight  float64
	MinPatternWidth   int
	MaxApexFactor     float64
	LastResistanceIdx int
	LastValleyIdx     int
	PEnd              int
	PatternWidth      float64
	CeilingBreakBar   int
	FloorBreakBar     int
	SupportCrossBar   int
	StageNote         string
}

type CheckVolumeDebugSnapshot struct {
	PatternStart   int
	PEnd           int
	Width          int
	MinWidth       int
	Skipped        bool
	SkipNote       string
	PointCount     int
	AvgVol         float64
	VolSlope       float64
	NormalizedSlope float64
	SlopeMax       float64
	Points         []SwingPoint
}

