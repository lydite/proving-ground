# proving-ground

The polyglot repository [lydite](https://github.com/lydite/lydite) is validated
against: a Rust CLI, a Go API and a React app, arranged to be scanned, tested,
gated and mutated.

It ships nothing. Nothing here is released, imported or deployed, and no code
here is worth reading for its own sake — every line exists to exercise something
in lydite's component platform ([ADR 0016](https://github.com/lydite/lydite/blob/main/docs/adr/0016-components-and-lydite-run-tests.md)).

## This repository is deliberately awkward

A tidy repository validates the happy path. **Several things here are wrong on
purpose, and fixing them destroys what they are for.** Each is listed below with
the behaviour it exists to prove. Before "correcting" anything, check this table.

| what looks wrong | why it is here |
|---|---|
| `rust/` is one component containing three crates, though only `tally-cli` is a deployable | a component is a **build unit**, not a deployable |
| `web/` is one component containing two packages | the same, on the node side — one workspace, one `node_modules`, one install |
| `go/api` and `go/sdk` are separate modules rather than one | the case where splitting genuinely **is** correct: separate version lines, consumed separately |
| `web` and `sdk` declare `depends_on: [api]` when nothing imports it | both are derived from `docs/openapi.json`, which `go/api` emits. No tool sees that edge, so it is declared |
| `go/api/compose.yaml` and `rust/compose.yaml` both publish host port `5432` | two components genuinely collide, so the local scheduler must serialise them |
| those two compose services are named `db` and `postgres` | the port lock is keyed on **ports, not service names**. Naming both `db` would let a name-keyed lock pass this repository |
| the root `Makefile` and `VERSION` sit under no component | files outside every component that invalidate exactly one. Only `tally` watches them |
| `tally-cli` embeds `VERSION` with `include_str!` | it makes that `watch` entry real. A `VERSION` nothing reads would let the entry be deleted with every test still passing |
| **`scripts/seed.ts` is under no component and is not excluded** | the orphan gate must fire on it. See below |
| four functions have poor or missing tests | mutation is only observable when mutants survive. See below |

### `scripts/seed.ts` is meant to be an orphan

The orphan gate fails on any source file under no component's directory and under
no explicit exclude. `scripts/seed.ts` is that file, and `.lydite/components.yml`
deliberately carries no exclude for it.

The tempting fix is wrong twice over. Widening the `web` component's `dir` does
not work — `web/` is an npm workspace and `scripts/` is outside it. Declaring a
component for it does not work either — nothing tests it and no runner builds it.
The honest resolution is an explicit exclude, and leaving that unwritten is what
keeps the gate observable.

### The four mutation cases

The same four function names exist in `rust/crates/tally-core`, `go/sdk` and
`web/packages/ui`. The identical source shape must produce the identical verdict
in all three languages; a divergence is a finding about lydite's mutation engine,
not about this repository.

| function | shape | expected verdict |
|---|---|---|
| `tally` | its test calls it and asserts nothing | every mutant **survives** |
| `percent` | no test calls it | mutants **skipped** — not reported as survivors |
| `label` | returns its argument; no operator, literal or branch | **no signal**, honestly reported |
| `floorZero` | `v < 0 ? 0 : v`, thoroughly tested | `<` → `<=` is **equivalent** and unkillable |

`floorZero` is the one to be careful with. Its tests are deliberately thorough,
because that is what makes the surviving mutant an equivalent one rather than a
symptom of a weak test — and telling those two apart is the whole point of being
able to acknowledge a mutant.

## Running it

Everything passes with no container runtime present. The tests that need
Postgres are skipped: `#[ignore]` in Rust, and `PROVING_GROUND_PG` in Go.

lydite runs those tests rather than skipping them, because it brings the
services up first — the `tally` component passes `--run-ignored=all` and the
`api` component sets `PROVING_GROUND_PG`. That is what keeps the two compose
declarations load-bearing: delete either one and a suite starts failing.

```sh
make test                     # every component's suite
make spec                     # regenerate docs/openapi.json from go/api
make version                  # rewrite VERSION, which tally-cli embeds
make up                       # bring up Postgres for go/api
make down                     # tear both compose files down
```

Per component:

```sh
cd rust   && cargo nextest run --workspace
cd go/api && go test ./...
cd go/sdk && go test ./...
cd web    && npm ci && npm test
```

`make up` starts only one of the two compose files, because they both bind
`5432` and the second would fail to bind. That is the collision, reproduced by
hand.

## Components

Declared in [`.lydite/components.yml`](.lydite/components.yml).

| component | dir | runner | services | watches |
|---|---|---|---|---|
| `tally` | `rust` | `cargo-nextest` | `postgres` on 5432 | `Makefile`, `VERSION` |
| `api` | `go/api` | `go-test` | `db` on 5432 | — |
| `sdk` | `go/sdk` | `go-test` | — | `docs/openapi.json` |
| `web` | `web` | `vitest` | — | `docs/openapi.json` |

This repository has no CI of its own. The check that keeps it from rotting lives
in lydite's `ci-end2end.yml`: a second copy of the expectations here would drift,
and the correct verdict for this repository is not all-green.
