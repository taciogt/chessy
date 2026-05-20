# Chessy

A terminal chess application in Go, built with extensibility in mind.

## Features

- Full chess rules engine (check, checkmate, stalemate)
- Hexagonal architecture — swap UIs and AI engines without touching the core
- Human vs Human mode (terminal UI)
- AI opponent via Minimax with Alpha-Beta pruning (in progress)
- Stockfish integration via UCI protocol (planned)

## Quick start

```sh
make run
```

## Development

See [CONTRIBUTING.md](CONTRIBUTING.md) for setup instructions.

## Architecture

```
core/       Pure chess domain — no external dependencies
ports/      Interfaces: Player (produces moves), Renderer (displays state)
adapters/   tui/ (Bubble Tea), minimax/ (AI), stockfish/ (UCI)
```

`core` never imports from `ports` or `adapters`. Adapters depend on `core` and `ports`.

## Roadmap

| Phase | Scope |
|---|---|
| MVP | Basic piece movement, check, checkmate, stalemate |
| 2 | Castling, pawn promotion |
| 3 | En passant |
| 4 | Draw conditions (50-move rule, threefold repetition) |
| 5 | Time controls |
