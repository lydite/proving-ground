//! The four mutation cases, in Rust.
//!
//! Each function pairs with a specific mutation-report outcome, and the same four
//! names exist in `go/sdk` and `web/packages/ui`. The three languages must agree:
//! a divergence is a finding about the mutation engine, not about this crate.

/// Sums a slice. Its test calls it and asserts nothing, so every mutant of the
/// accumulator and its initial value survives.
pub fn tally(items: &[i64]) -> i64 {
    let mut total = 0;
    for item in items {
        total += item;
    }
    total
}

/// Has no test at all. Its mutants are skipped for want of covering tests, and
/// must not be reported as survivors — the two are different verdicts.
pub fn percent(part: i64, whole: i64) -> f64 {
    if whole == 0 {
        return 0.0;
    }
    (part as f64 / whole as f64) * 100.0
}

/// Returns its argument. No operator, no literal, no branch, so there is nothing
/// to mutate and the honest report is "no signal" rather than a score.
pub fn label(name: String) -> String {
    name
}

/// Clamps negatives to zero. Replacing `<` with `<=` is an equivalent mutant: at
/// `v == 0` both arms yield `0`, so no input distinguishes them. The tests below
/// are thorough, which is what makes this unkillable rather than merely untested.
pub fn floor_zero(v: i64) -> i64 {
    if v < 0 {
        0
    } else {
        v
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn tally_runs() {
        tally(&[1, 2, 3]);
    }

    #[test]
    fn label_returns_its_argument() {
        assert_eq!(label("cli".to_string()), "cli");
    }

    #[test]
    fn floor_zero_clamps_negatives() {
        assert_eq!(floor_zero(-7), 0);
        assert_eq!(floor_zero(-1), 0);
    }

    #[test]
    fn floor_zero_passes_through_zero_and_positives() {
        assert_eq!(floor_zero(0), 0);
        assert_eq!(floor_zero(1), 1);
        assert_eq!(floor_zero(42), 42);
    }
}
