import { describe, expect, it } from "vitest";

import { countersPath, pathForOperation } from "./spec.js";

describe("pathForOperation", () => {
  it("reads the counters route out of the spec", () => {
    expect(countersPath).toBe("/counters");
  });

  it("rejects an operation the spec does not describe", () => {
    expect(() => pathForOperation("deleteEverything")).toThrow();
  });
});
