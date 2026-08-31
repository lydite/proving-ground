import { describe, expect, it } from "vitest";

import { floorZero, label, tally } from "./counters.js";

describe("tally", () => {
  it("runs", () => {
    tally([
      { name: "alpha", count: 1 },
      { name: "zulu", count: 2 },
    ]);
  });
});

describe("label", () => {
  it("returns its argument", () => {
    expect(label("cli")).toBe("cli");
  });
});

describe("floorZero", () => {
  it("clamps negatives", () => {
    expect(floorZero(-7)).toBe(0);
    expect(floorZero(-1)).toBe(0);
  });

  it("passes through zero and positives", () => {
    expect(floorZero(0)).toBe(0);
    expect(floorZero(1)).toBe(1);
    expect(floorZero(42)).toBe(42);
  });
});
