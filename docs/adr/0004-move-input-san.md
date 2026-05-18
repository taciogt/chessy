# ADR-0004: Formato de Input de Movimentos — SAN com Preview

## Status
Accepted

## Context

A forma como o jogador insere movimentos afeta tanto a UX quanto a representação interna de `Move`. O autor é iniciante em xadrez e quer usar o projeto para se familiarizar com os conceitos reais do jogo, incluindo notação.

## Decision

**Fase 1:** Standard Algebraic Notation (SAN) com preview visual do movimento antes da confirmação.
**Fase 2:** Cursor interativo (teclas de seta para selecionar peça e destino).

O preview funciona assim: o jogador digita o movimento em SAN, o tabuleiro exibe visualmente a jogada (destacando a peça e o destino), e o Enter confirma a execução.

## Alternatives Considered

| Alternativa | Motivo de rejeição |
|---|---|
| **Long algebraic notation** (`e2e4`) | Sem ambiguidade e fácil de parsear, mas não é o que jogadores reais usam. Perde o objetivo de aprender notação. |
| **Cursor interativo direto** | UX superior, mas implementação mais complexa no Bubble Tea. Adia o início do jogo jogável. |
| **Coordenadas com espaço** (`e2 e4`) | Arbitrário, sem vantagem sobre long algebraic. |

## Consequences

- O parser de SAN precisa conhecer o estado atual do tabuleiro para resolver ambiguidades (ex: qual cavalo pode ir para f3?). Isso acopla o parsing ao `core`, mas é inevitável na SAN.
- O preview de movimentos requer que o `Renderer` suporte `RenderHints` com um `PreviewMove` (ver ADR-0005).
- O cursor interativo, quando implementado, não muda a representação interna de `Move` — apenas o adapter de input muda.
- Digitar `Nf3` enquanto aprende xadrez reforça o reconhecimento da notação padrão usada em livros, puzzles e análise online.
