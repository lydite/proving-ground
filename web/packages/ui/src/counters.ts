// The four mutation cases, in TypeScript. Each function pairs with a specific
// mutation-report outcome, and the same four names exist in
// rust/crates/tally-core and go/sdk. The three languages must agree: a divergence
// is a finding about the mutation engine, not about this package.

export interface Counter {
  name: string;
  count: number;
}

/** Sums counts. Its test calls it and asserts nothing, so every mutant of the
 * accumulator and its initial value survives. */
export function tally(counters: Counter[]): number {
  let total = 0;
  for (const counter of counters) {
    total += counter.count;
  }
  return total;
}

/** Has no test at all. Its mutants are skipped for want of covering tests, and
 * must not be reported as survivors — the two are different verdicts. */
export function percent(part: number, whole: number): number {
  if (whole === 0) {
    return 0;
  }
  return (part / whole) * 100;
}

/** Returns its argument. No operator, no literal, no branch, so there is nothing
 * to mutate and the honest report is "no signal" rather than a score. */
export function label(name: string): string {
  return name;
}

/** Clamps negatives to zero. Replacing `<` with `<=` is an equivalent mutant: at
 * `v === 0` both arms yield `0`, so no input distinguishes them. The tests are
 * thorough, which is what makes this unkillable rather than merely untested. */
export function floorZero(v: number): number {
  if (v < 0) {
    return 0;
  }
  return v;
}
