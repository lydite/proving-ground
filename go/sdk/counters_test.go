package sdk_test

import (
	"testing"

	"github.com/lydite/proving-ground/go/sdk"
)

func TestTallyRuns(t *testing.T) {
	sdk.Tally([]sdk.Counter{{Name: "alpha", Count: 1}, {Name: "zulu", Count: 2}})
}

func TestLabelReturnsItsArgument(t *testing.T) {
	if got := sdk.Label("cli"); got != "cli" {
		t.Errorf("Label(cli) = %q, want %q", got, "cli")
	}
}

func TestFloorZeroClampsNegatives(t *testing.T) {
	for _, v := range []int64{-7, -1} {
		if got := sdk.FloorZero(v); got != 0 {
			t.Errorf("FloorZero(%d) = %d, want 0", v, got)
		}
	}
}

func TestFloorZeroPassesThroughZeroAndPositives(t *testing.T) {
	for _, v := range []int64{0, 1, 42} {
		if got := sdk.FloorZero(v); got != v {
			t.Errorf("FloorZero(%d) = %d, want %d", v, got, v)
		}
	}
}
