package valueobjects

// TargetingRuleVO represents a rule for targeting users with a setting flag.
type TargetingRuleVO struct {
	Percentage int
	Value      bool // TODO value should be boolean true := "on" | false := "off", int, double or string
}
