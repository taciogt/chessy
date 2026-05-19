# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

All common tasks go through the `Makefile`:

| Command | Purpose |
|---|---|
| `make test` | Run all tests (`go test ./...`) |
| `make test-v` | Verbose test output |
| `make cover` | Tests with coverage |
| `make build` | Build all packages |
| `make run` | Run the application |
| `make fmt` / `make vet` / `make tidy` | Standard Go hygiene |

Run a single test or package:
```
go test ./core -run TestComputeStatus
go test ./core -run TestComputeStatus/checkmate -v
```

Go 1.25, module `github.com/taciogt/chessy`. No third-party dependencies yet.

## Architecture

Hexagonal (Ports & Adapters). See `CONTEXT.md` for the domain glossary and `docs/adr/` for decisions.

```
core/       pure chess domain — no external deps
ports/      interfaces: Player (produces moves), Renderer (displays state)
adapters/   tui/ (Bubble Tea), minimax/ (AI), stockfish/ (UCI) — currently empty stubs
```

`core` must never import from `ports` or `adapters`. Adapters depend on `core` and `ports`. The orchestrator (TBD) wires a `Player` + `Renderer` + `core.GameState` together.

### Core invariants

- **Board layout**: `board[rank][file]`, where `board[0][0]` = a1 and `board[7][7]` = h8. `File` 0–7 → a–h; `Rank` 0–7 → 1–8.
- **GameState is immutable by convention**: `ApplyMove` returns a new `GameState` and does not mutate the input. The new `History` slice has its own backing array so derived states stay independent.
- **`ApplyMove` does not legality-check** — callers source moves from `LegalMoves`.
- **`LegalMoves` = pseudo-legal filtered by "doesn't leave own King in check"**, implemented by playing the move and asking `IsSquareAttacked` about the resulting King square.

### Phased-implementation guard (important)

The engine is being built piece-by-piece. **Today only the King has movement logic**:
- `PseudoLegalMoves` returns `nil` for non-King kinds.
- `IsSquareAttacked` only considers enemy Kings.
- `ComputeStatus` (in `core/status.go`) uses `activeSideHasOnlyImplementedPieces` to avoid falsely declaring Checkmate/Stalemate when un-implemented pieces are on the board — it downgrades to `Check` or `Ongoing` instead.

**When adding a new piece type**, update *all three* in lockstep: `PseudoLegalMoves` switch, `IsSquareAttacked` enumeration, and the implemented-kinds set in `activeSideHasOnlyImplementedPieces`. Otherwise terminal-state detection silently breaks.

### Render hints

Per ADR-0005, the `Renderer` interface will eventually take a `RenderHints` second argument (preview moves, highlighted squares) so beginner/debug overlays live in the core, not the adapter. The current `Renderer` signature is `Render(state core.GameState)` — extend it carefully when hints land.

## Testing

Per ADR-0009: high coverage in `core` via table-driven tests; **no automated TUI tests** (validated manually). AI adapters get correctness tests via known positions (e.g. mate-in-1). The core tests double as executable documentation of the rules.

## Roadmap phases

Rules: MVP (basic moves + check/mate/stalemate) → castling + promotion → en passant → draw conditions → clock. AI: Minimax+AB → +transposition tables → Stockfish via UCI. Modes: H-vs-H → H-vs-AI → AI-vs-AI. See `CONTEXT.md` for the full table.

ADRs (`docs/adr/`) are authoritative for any decision they cover — consult before changing architecture, board representation, input format, or tooling.
