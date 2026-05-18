# Chessy — Context

A terminal chess application in Go, designed for extensibility across UIs and AI engines.

## Glossary

### Domain

**Square** — A single cell on the board, identified by file (a–h) and rank (1–8). e.g. `e4`.

**Piece** — A chess piece with a type (Pawn, Rook, Knight, Bishop, Queen, King) and a color (White, Black).

**Board** — The 8×8 grid of squares. Internally represented as a 2D array `board[rank][file]`.

**Move** — A transition of a piece from one square to another. Represented as `{From: Square, To: Square}`. Encoding in user input follows SAN.

**SAN (Standard Algebraic Notation)** — The standard chess notation for moves (e.g. `e4`, `Nf3`, `O-O`). The primary input format.

**Game** — The full state of a chess match: board position, active color, move history, and game status.

**GameStatus** — The current outcome state: `Ongoing`, `Check`, `Checkmate`, `Stalemate`.

**Check** — The active player's King is under attack.

**Checkmate** — The active player is in Check with no legal moves. The game is over.

**Stalemate** — The active player has no legal moves but is not in Check. A draw.

**Legal Move** — A move that is valid for the piece type, does not leave the active King in Check, and respects the current board state.

**Player** — Any agent that can produce a Move given a GameState. Can be a human (via terminal input) or an AI engine.

**Renderer** — Any agent that can display a GameState. Can be a terminal UI, a web frontend, or a graphical engine.

**RenderHints** — Optional overlay information passed to a Renderer alongside GameState. Includes preview moves and highlighted squares (e.g. attacked squares, valid move targets). Enables beginner and debug modes without coupling the core to any specific UI.

### Architecture

**Core** — The chess domain: rules, move validation, game state. Has no external dependencies.

**Port** — An interface (contract) that separates the Core from the outside world. Two primary ports: `Player` and `Renderer`.

**Adapter** — A concrete implementation of a Port. Examples: `tui` (terminal UI), `minimax` (AI engine), `stockfish` (UCI engine integration).

**UCI (Universal Chess Interface)** — The protocol used to communicate with external chess engines like Stockfish.

### AI

**Minimax** — A search algorithm that evaluates all possible moves to a given depth, minimizing the opponent's best outcome. The first AI adapter.

**Alpha-Beta Pruning** — An optimization for Minimax that skips branches that cannot affect the result.

**Transposition Table** — A cache of previously evaluated positions, used in the intermediate AI level to avoid redundant computation.

**Depth** — How many half-moves (plies) ahead the Minimax searches.

## Phased Implementation Plan

| Phase | Scope |
|---|---|
| MVP | Basic piece movement, check, checkmate, stalemate |
| 2 | Castling, pawn promotion |
| 3 | En passant |
| 4 | Draw conditions (50-move rule, threefold repetition, insufficient material) |
| 5 | Time controls (clock) |

### AI Phases

| Level | Technique |
|---|---|
| 1 | Minimax + Alpha-Beta Pruning |
| 2 | + Transposition tables + positional evaluation |
| 3 | Stockfish via UCI protocol |

### Game Mode Phases

| Phase | Mode |
|---|---|
| 1 | Human vs Human |
| 2 | Human vs AI |
| 3 | AI vs AI (debug/testing) |

### Input Phases

| Phase | Format |
|---|---|
| 1 | SAN with move preview before confirmation |
| 2 | Interactive cursor (arrow keys) |
