package store_test

import (
	"os"
	"testing"
	"time"

	"github.com/lydite/proving-ground/go/api/internal/store"
)

// Requires the db service from go/api/compose.yaml, so it is skipped unless
// PROVING_GROUND_PG is set and the suite passes on a machine with no container
// runtime.
func TestDatabaseIsReachable(t *testing.T) {
	if os.Getenv("PROVING_GROUND_PG") == "" {
		t.Skip("set PROVING_GROUND_PG to run against the db service")
	}
	if !store.Reachable(store.DefaultAddr, 2*time.Second) {
		t.Fatalf("%s is not accepting connections", store.DefaultAddr)
	}
}
