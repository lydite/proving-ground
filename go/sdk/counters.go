package sdk

// The four mutation cases, in Go. Each function pairs with a specific
// mutation-report outcome, and the same four names exist in
// rust/crates/tally-core and web/packages/ui. The three languages must agree: a
// divergence is a finding about the mutation engine, not about this package.

// Tally sums counts. Its test calls it and asserts nothing, so every mutant of
// the accumulator and its initial value survives.
func Tally(counters []Counter) int64 {
	var total int64 = 0
	for _, counter := range counters {
		total += counter.Count
	}
	return total
}

// Percent has no test at all. Its mutants are skipped for want of covering
// tests, and must not be reported as survivors — the two are different verdicts.
func Percent(part, whole int64) float64 {
	if whole == 0 {
		return 0.0
	}
	return (float64(part) / float64(whole)) * 100.0
}

// Label returns its argument. No operator, no literal, no branch, so there is
// nothing to mutate and the honest report is "no signal" rather than a score.
func Label(name string) string {
	return name
}

// FloorZero clamps negatives to zero. Replacing < with <= is an equivalent
// mutant: at v == 0 both arms yield 0, so no input distinguishes them. The tests
// are thorough, which is what makes this unkillable rather than merely untested.
func FloorZero(v int64) int64 {
	if v < 0 {
		return 0
	}
	return v
}
