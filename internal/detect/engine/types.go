package engine

import "github.com/gopherchan2006/go-triangle-detector/internal/detect/spec"

type SwingPoint struct {
	Index int
	Value float64
}

type StepDebugLogs struct {
	CalcATR                  string
	FindSwingHighs           string
	FindHorizontalResistance string
	CheckTimingAndHighs      string
	FindValleys              string
	ValidateValleys          string
	FitSupportLine           string
	CheckGeometry            string
	CheckVolume              string
}

type ATRDebug struct {
	AvgPrice float64
	ATRValue float64
	Vol      float64
}

type SwingDebug struct {
	SwingHighsCount int
}

type ResistanceDebug struct {
	ResistanceLevel    float64
	ResistanceTouches  int
	FirstTouchIdx      int
	HighAboveThreshold float64
	CrashThreshold     float64
}

type SupportDebug struct {
	ValleysCount     int
	FirstVIdx        int
	MaxCrashRange    float64
	AllowedFlat      float64
	SupportSlope     float64
	SupportIntercept float64
	MaxValleyDepth   float64
	ValleyDeviation  float64
}

type GeometryDebug struct {
	PatternStart      int
	PatternEnd        int
	XIntersect        float64
	LastX             float64
	CeilingTol        float64
	Ceiling           float64
	FloorTol          float64
	HeightAtStart     float64
	HeightAtEnd       float64
	LastResistanceIdx int
	LastValleyIdx     int
	PEnd              int
	PatternWidth      float64
}

type DebugInfo struct {
	Logs       StepDebugLogs
	ATR        ATRDebug
	Swing      SwingDebug
	Resistance ResistanceDebug
	Support    SupportDebug
	Geometry   GeometryDebug
}

type Result struct {
	Found                 bool
	RejectReason          spec.RejectReason
	ResistanceLevel       float64
	ResistanceTouches     int
	ResistanceTouchPoints []SwingPoint
	SupportSlope          float64
	SupportIntercept      float64
	SupportTouchPoints    []SwingPoint
	Debug                 DebugInfo
	TargetPrice           float64
	BreakoutDetected      bool
	BreakoutVolumeRatio   float64
}

