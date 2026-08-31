package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lydite/proving-ground/go/api"
)

func TestHandlerListsCountersOrderedByName(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/counters", nil)

	api.Handler(map[string]int64{"zulu": 3, "alpha": 1}).ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	var got []api.Counter
	if err := json.NewDecoder(rec.Body).Decode(&got); err != nil {
		t.Fatalf("decode: %v", err)
	}
	want := []api.Counter{{Name: "alpha", Count: 1}, {Name: "zulu", Count: 3}}
	if len(got) != len(want) {
		t.Fatalf("got %d counters, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("counter %d = %+v, want %+v", i, got[i], want[i])
		}
	}
}

func TestSpecDescribesTheCountersRoute(t *testing.T) {
	paths, ok := api.Spec()["paths"].(map[string]any)
	if !ok {
		t.Fatal("spec has no paths object")
	}
	if _, ok := paths["/counters"]; !ok {
		t.Error("spec does not describe /counters")
	}
}
