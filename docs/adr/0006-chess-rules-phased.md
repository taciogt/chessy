# ADR-0006: Implementação Faseada das Regras do Xadrez

## Status
Accepted

## Context

O xadrez tem regras completas que incluem casos especiais complexos (en passant, roque, promoção, condições de empate). Implementar tudo de uma vez aumenta o risco de travar na implementação antes de ter um jogo jogável, e dificulta o aprendizado incremental.

## Decision

Implementar as regras em fases, da mais simples à mais complexa:

| Fase | Escopo |
|---|---|
| **MVP** | Movimento básico das 6 peças, detecção de xeque, xeque-mate, stalemate |
| **2** | Roque (kingside e queenside), promoção de peão |
| **3** | En passant |
| **4** | Condições de empate: regra dos 50 lances, repetição tripla, material insuficiente |
| **5** | Controle de tempo (relógio) |

## Alternatives Considered

| Alternativa | Motivo de rejeição |
|---|---|
| **Implementar regras completas desde o início** | Alto risco de travar nas bordas (en passant, roque) antes de ter um jogo funcional. Dificulta o aprendizado da arquitetura. |
| **MVP sem xeque-mate** | Tornaria o jogo injogável — sem condição de fim. |

## Consequences

- O MVP produz um jogo tecnicamente incompleto (sem roque, en passant), mas totalmente jogável para aprendizado.
- En passant e roque são casos especiais que requerem estado adicional no `GameState` (ex: disponibilidade de roque, casa de en passant disponível). O `GameState` precisa ser projetado para acomodar esses campos desde o início, mesmo que não sejam usados no MVP.
- As condições de empate da Fase 4 exigem histórico de movimentos — o `Game` deve manter esse histórico desde o início.
