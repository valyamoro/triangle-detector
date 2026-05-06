package detect

import (
	"github.com/gopherchan2006/go-triangle-detector/internal/detect/engine"
	"github.com/gopherchan2006/go-triangle-detector/internal/detect/spec"
)

type Result = engine.Result
type SwingPoint = engine.SwingPoint
type DebugInfo = engine.DebugInfo
type StepDebugLogs = engine.StepDebugLogs
type ATRDebug = engine.ATRDebug
type SwingDebug = engine.SwingDebug
type ResistanceDebug = engine.ResistanceDebug
type SupportDebug = engine.SupportDebug
type GeometryDebug = engine.GeometryDebug
type Params = spec.Params
type SwingParams = spec.SwingParams
type HorizontalParams = spec.HorizontalParams
type TimingParams = spec.TimingParams
type SupportFitParams = spec.SupportFitParams
type GeometryParams = spec.GeometryParams
type VolumeDeclParams = spec.VolumeDeclParams
type BreakoutParams = spec.BreakoutParams
type RejectReason = spec.RejectReason
type RejectCounter = spec.RejectCounter

const (
	ReasonFewSwingHighs            = spec.ReasonFewSwingHighs
	ReasonResistanceLt3Touches     = spec.ReasonResistanceLt3Touches
	ReasonHighBeforeFirstTouch     = spec.ReasonHighBeforeFirstTouch
	ReasonCrashBeforeFirstTouch    = spec.ReasonCrashBeforeFirstTouch
	ReasonFirstTouchTooLate        = spec.ReasonFirstTouchTooLate
	ReasonFewValleys               = spec.ReasonFewValleys
	ReasonValleyNotRising          = spec.ReasonValleyNotRising
	ReasonNegativeSlope            = spec.ReasonNegativeSlope
	ReasonValleyTooDeep            = spec.ReasonValleyTooDeep
	ReasonLowRSquared              = spec.ReasonLowRSquared
	ReasonValleyOffSupportLine     = spec.ReasonValleyOffSupportLine
	ReasonNoConvergence            = spec.ReasonNoConvergence
	ReasonBreaksCeiling            = spec.ReasonBreaksCeiling
	ReasonBreaksSupportFloor       = spec.ReasonBreaksSupportFloor
	ReasonSupportAboveResistance   = spec.ReasonSupportAboveResistance
	ReasonNotNarrowing             = spec.ReasonNotNarrowing
	ReasonTooFlat                  = spec.ReasonTooFlat
	ReasonTooNarrow                = spec.ReasonTooNarrow
	ReasonApexTooFar               = spec.ReasonApexTooFar
	ReasonFirstValleyCrash         = spec.ReasonFirstValleyCrash
	ReasonFirstValleyNotFloor      = spec.ReasonFirstValleyNotFloor
	ReasonPrecedingTrendNotUp      = spec.ReasonPrecedingTrendNotUp
	ReasonVolumeNotDeclining       = spec.ReasonVolumeNotDeclining
	ReasonSupportSlopeTooFlat      = spec.ReasonSupportSlopeTooFlat
	ReasonResistanceLastTouchEarly = spec.ReasonResistanceLastTouchEarly
	ReasonResistanceGapTooLong     = spec.ReasonResistanceGapTooLong
)

func DefaultParams() Params {
	return spec.DefaultParams()
}
