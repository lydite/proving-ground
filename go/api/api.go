// Package api serves the counters endpoint and owns the OpenAPI description of
// it. The description is emitted to docs/openapi.json by cmd/genspec, and both
// web/packages/app and go/sdk are generated from that file — an edge across
// three languages that no build tool can derive, so the components declare it.
package api

import (
	"encoding/json"
	"net/http"
	"sort"
)

// Counter is one named tally.
type Counter struct {
	Name  string `json:"name"`
	Count int64  `json:"count"`
}

// Handler serves the routes described by Spec.
func Handler(counters map[string]int64) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /counters", func(w http.ResponseWriter, r *http.Request) {
		names := make([]string, 0, len(counters))
		for name := range counters {
			names = append(names, name)
		}
		sort.Strings(names)

		out := make([]Counter, 0, len(names))
		for _, name := range names {
			out = append(out, Counter{Name: name, Count: counters[name]})
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(out)
	})
	return mux
}

// Spec is the OpenAPI description of Handler's routes. Marshalling it sorts map
// keys, so cmd/genspec reproduces docs/openapi.json byte for byte.
func Spec() map[string]any {
	counter := map[string]any{
		"type":     "object",
		"required": []string{"name", "count"},
		"properties": map[string]any{
			"name":  map[string]any{"type": "string"},
			"count": map[string]any{"type": "integer", "format": "int64"},
		},
	}

	return map[string]any{
		"openapi": "3.1.0",
		"info": map[string]any{
			"title":   "proving-ground counters",
			"version": "0.1.0",
		},
		"paths": map[string]any{
			"/counters": map[string]any{
				"get": map[string]any{
					"operationId": "listCounters",
					"summary":     "List every counter",
					"responses": map[string]any{
						"200": map[string]any{
							"description": "the counters, ordered by name",
							"content": map[string]any{
								"application/json": map[string]any{
									"schema": map[string]any{
										"type":  "array",
										"items": map[string]any{"$ref": "#/components/schemas/Counter"},
									},
								},
							},
						},
					},
				},
			},
		},
		"components": map[string]any{
			"schemas": map[string]any{"Counter": counter},
		},
	}
}
