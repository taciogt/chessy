# ADR-0007: Progressão da IA — Minimax até Stockfish

## Status
Accepted

## Context

O projeto quer suportar um oponente de IA, mas o autor quer implementar a IA do zero para aprender os conceitos — não apenas integrar uma engine pronta. Ao mesmo tempo, quer ter acesso a uma IA de nível mundial eventualmente.

## Decision

Implementar a IA em três níveis progressivos, cada um como um adapter separado:

| Nível | Técnica | Adapter |
|---|---|---|
| **1** | Minimax com Alpha-Beta Pruning, profundidade fixa | `adapters/minimax` |
| **2** | + Tabelas de transposição + avaliação posicional | `adapters/minimax` (evoluído) |
| **3** | Stockfish via protocolo UCI | `adapters/stockfish` |

## Alternatives Considered

| Alternativa | Motivo de rejeição |
|---|---|
| **Integrar Stockfish direto** | Não há aprendizado sobre IA de xadrez. O objetivo é entender Minimax. |
| **Só Minimax, nunca Stockfish** | Perde a oportunidade de ter uma IA forte para desafio real. |
| **Usar uma biblioteca de Minimax** | Elimina o aprendizado do algoritmo, que é central para o objetivo do projeto. |

## Consequences

- O Minimax básico (Nível 1) vai ser visivelmente lento em profundidades maiores que 5–6 com array 8×8. Isso é esperado e aceitável inicialmente.
- Alpha-Beta Pruning é implementado junto com o Minimax (não separado) — não faz sentido implementar Minimax puro.
- A transição de Nível 1 para 2 é uma evolução do mesmo adapter. A transição para Stockfish é um adapter novo — o `core` e o port `Player` não mudam.
- O protocolo UCI (Universal Chess Interface) permite que qualquer engine compatível (não apenas Stockfish) seja conectado pelo adapter da Fase 3.
