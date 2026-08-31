# Convenience targets, plus the one target that participates in a build.
#
# This file sits under no component. Only the `tally` component watches it, and
# `make version` is why: it writes VERSION, which rust/crates/tally-cli embeds.
# Changing VERSION must run `tally` and nothing else.

VERSION_VALUE ?= 0.1.0

.PHONY: version spec test up down

version:
	printf '%s\n' '$(VERSION_VALUE)' > VERSION

spec:
	cd go/api && go run ./cmd/genspec

test:
	cd rust && cargo nextest run
	cd go/api && go test ./...
	cd go/sdk && go test ./...
	cd web && npm test

up:
	podman compose -f go/api/compose.yaml up -d

down:
	podman compose -f go/api/compose.yaml down -v
	podman compose -f rust/compose.yaml down -v
