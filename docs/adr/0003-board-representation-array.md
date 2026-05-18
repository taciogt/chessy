# ADR-0003: Representação do Tabuleiro — Array 8×8

## Status
Accepted

## Context

A representação interna do tabuleiro afeta a complexidade de implementação das regras, a legibilidade do código e a performance futura da IA. A decisão deve equilibrar simplicidade inicial com a capacidade de evoluir quando a performance do Minimax se tornar um gargalo.

## Decision

Usar **array 8×8** (`board[rank][file]`) onde cada célula contém uma `Piece` ou é vazia.

Refatorar para **representação híbrida** (8×8 + listas de peças) quando a performance do Minimax em profundidades maiores se tornar um problema observado.

## Alternatives Considered

| Alternativa | Motivo de rejeição |
|---|---|
| **Listas de peças** | "O que tem nessa casa?" exige busca linear. Menos intuitivo para implementar regras. |
| **Bitboards** | 12 inteiros de 64 bits, um por tipo/cor de peça. Extremamente rápido, mas ilegível para quem está aprendendo xadrez e game dev simultaneamente. Otimização prematura. |
| **Híbrido agora** | Redundância sem necessidade comprovada. Adiciona complexidade de manter dois estados sincronizados. |

## Consequences

- `board[0][0]` = a1, `board[7][7]` = h8. Layout visual direto enquanto se escreve o código.
- Verificar "o que tem nessa casa?" é O(1).
- Iterar sobre todas as peças de uma cor é O(64) — aceitável para MVP e Minimax básico.
- Quando a profundidade de busca do Minimax começar a impactar performance, o refactor natural é adicionar listas de peças (sem eliminar o array).
- Bitboards ficam disponíveis como otimização futura se necessário, mas não são o próximo passo óbvio.
