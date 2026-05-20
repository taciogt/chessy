# Contributing to Chessy

## Prerequisites

- **Go 1.25+** — [install](https://go.dev/dl/)
- **golangci-lint** — static analysis

### Install golangci-lint

**macOS (Homebrew)**:
```sh
brew install golangci-lint
```

**Linux / other**:
```sh
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(go env GOPATH)/bin
```

See the [official install guide](https://golangci-lint.run/welcome/install/) for other options.

## Common tasks

| Command | Purpose |
|---|---|
| `make test` | Run all tests |
| `make test-v` | Verbose test output |
| `make cover` | Tests with coverage |
| `make lint` | Run golangci-lint |
| `make fmt` | Format code |
| `make vet` | Run go vet |
| `make build` | Build all packages |
| `make run` | Run the application |
| `make tidy` | Tidy go modules |

Run a single test:
```sh
go test ./core -run TestComputeStatus
go test ./core -run TestComputeStatus/checkmate -v
```

## Guidelines

- All new chess rules live in `core/` — no external dependencies allowed there.
- `core` must never import from `ports` or `adapters`.
- Tests are table-driven; the core tests double as executable documentation of the rules.
- When adding a new piece type, update `PseudoLegalMoves`, `IsSquareAttacked`, and `activeSideHasOnlyImplementedPieces` in lockstep (see `CLAUDE.md`).
- ADRs in `docs/adr/` are authoritative for architecture decisions — consult before changing board representation, input format, or tooling.
