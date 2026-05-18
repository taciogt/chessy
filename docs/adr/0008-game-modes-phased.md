# ADR-0008: Modos de Jogo — Humano vs Humano Primeiro

## Status
Accepted

## Context

O projeto eventualmente quer suportar Humano vs IA e IA vs IA, mas implementar o adapter de IA em paralelo com o core e a TUI aumenta o escopo inicial desnecessariamente.

## Decision

Implementar modos de jogo em fases:

| Fase | Modo |
|---|---|
| **1** | Humano vs Humano |
| **2** | Humano vs IA (após Minimax funcionar) |
| **3** | IA vs IA (modo debug/teste) |

## Alternatives Considered

| Alternativa | Motivo de rejeição |
|---|---|
| **Humano vs IA desde o início** | Exige implementar o adapter de Minimax em paralelo com o core e a TUI. Aumenta o escopo do MVP. |
| **Só Humano vs IA** | Impossibilita testar o jogo sem a IA funcionando. Humano vs Humano é o melhor modo para validar as regras. |

## Consequences

- O MVP entrega um jogo jogável (dois jogadores no mesmo terminal) antes de qualquer IA ser implementada.
- A interface `Player` é definida desde o início, mas a implementação de `HumanPlayer` é a única necessária no MVP.
- IA vs IA (Fase 3) não requer TUI interativa — útil para testar dois adapters de IA sem intervenção humana.
