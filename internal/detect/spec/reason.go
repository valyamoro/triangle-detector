package spec

type RejectReason string

const (
	ReasonFewSwingHighs            RejectReason = "01_few_swing_highs"
	ReasonResistanceLt3Touches     RejectReason = "02_resistance_<3_touches"
	ReasonHighBeforeFirstTouch     RejectReason = "03_high_before_first_touch"
	ReasonCrashBeforeFirstTouch    RejectReason = "04_crash_before_first_touch"
	ReasonFirstTouchTooLate        RejectReason = "05_first_touch_too_late"
	ReasonFewValleys               RejectReason = "06_few_valleys"
	ReasonValleyNotRising          RejectReason = "07_valley_not_rising"
	ReasonValleyTooDeep            RejectReason = "09_valley_too_deep"
	ReasonLowRSquared              RejectReason = "10_low_r_squared"
	ReasonValleyOffSupportLine     RejectReason = "11_valley_off_support_line"
	ReasonNoConvergence            RejectReason = "12_no_convergence"
	ReasonBreaksCeiling            RejectReason = "13_breaks_ceiling"
	ReasonBreaksSupportFloor       RejectReason = "14_breaks_support_floor"
	ReasonSupportAboveResistance   RejectReason = "15_support_above_resistance"
	ReasonNotNarrowing             RejectReason = "16_not_narrowing"
	ReasonTooFlat                  RejectReason = "17_too_flat"
	ReasonTooNarrow                RejectReason = "18_too_narrow"
	ReasonApexTooFar               RejectReason = "19_apex_too_far"
	ReasonVolumeNotDeclining       RejectReason = "23_volume_not_declining"
	ReasonSupportSlopeTooFlat      RejectReason = "24_support_slope_too_flat"
	ReasonResistanceLastTouchEarly RejectReason = "25_resistance_last_touch_early"
	ReasonResistanceGapTooLong     RejectReason = "26_resistance_gap_too_long"
)
