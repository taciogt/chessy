# ADR-0001: Linguagem de Implementação — Go

## Status
Accepted

## Context

O projeto é uma aplicação de xadrez no terminal com planos de implementar IA própria (Minimax com Alpha-Beta Pruning). O autor conhece Python, Go e TypeScript. Os critérios de decisão foram: performance para IA, ecossistema TUI, e simplicidade de distribuição.

## Decision

Usar **Go**.

## Alternatives Considered

| Alternativa | Motivo de rejeição |
|---|---|
| **Python** | Performance insuficiente para busca Minimax em profundidades maiores que 5–6. Cerca de 15–20x mais lento que Go para o mesmo algoritmo. |
| **TypeScript** | Nenhuma vantagem clara para esse contexto. Toolchain mais pesada, sem ganho real em performance ou ecossistema TUI. |

## Consequences

- O ecossistema **Charm.sh** (Bubble Tea, Lip Gloss, Bubbles) cobre o TUI com qualidade suficiente para o projeto.
- Binário único facilita distribuição.
- Quando o Minimax precisar de otimização (profundidade maior), Go suporta sem refactor de linguagem.
- Se no futuro houver um frontend web, será necessário uma API separada (Go não compila para browser nativamente sem WASM).
