// Command genspec writes the OpenAPI description of the API to docs/openapi.json.
// The file is committed: web/packages/app and go/sdk read it, and requiring a Go
// toolchain to run a vitest suite would hide a broken TypeScript runner behind a
// Go failure.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/lydite/proving-ground/go/api"
)

func main() {
	out := flag.String("out", "../../docs/openapi.json", "where to write the spec")
	flag.Parse()

	body, err := json.MarshalIndent(api.Spec(), "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "genspec: %v\n", err)
		os.Exit(1)
	}
	body = append(body, '\n')

	if err := os.WriteFile(*out, body, 0o644); err != nil {
		fmt.Fprintf(os.Stderr, "genspec: %v\n", err)
		os.Exit(1)
	}
}
