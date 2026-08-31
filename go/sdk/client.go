// Package sdk is the client for the counters API, derived from docs/openapi.json.
//
// It is a separate module from go/api because it is versioned and consumed on its
// own: the server and its client do not release together. That makes go/api and
// go/sdk two components, which is the case where splitting genuinely is correct.
package sdk

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Counter mirrors the Counter schema in docs/openapi.json.
type Counter struct {
	Name  string `json:"name"`
	Count int64  `json:"count"`
}

// Client calls the counters API.
type Client struct {
	BaseURL string
	HTTP    *http.Client
}

// ListCounters implements the listCounters operation.
func (c *Client) ListCounters(ctx context.Context) ([]Counter, error) {
	httpClient := c.HTTP
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.BaseURL+"/counters", nil)
	if err != nil {
		return nil, err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("listCounters: unexpected status %d", resp.StatusCode)
	}

	var out []Counter
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return out, nil
}
