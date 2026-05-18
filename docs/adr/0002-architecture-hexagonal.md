# ADR-0002: Arquitetura Hexagonal (Ports & Adapters)

## Status
Accepted

## Context

O projeto tem como objetivo explícito separar a lógica de xadrez da UI, de forma que o terminal seja apenas o primeiro frontend — com possibilidade futura de UI gráfica, web, etc. A IA também deve ser intercambiável (Minimax básico → intermediário → Stockfish).

## Decision

Usar **Arquitetura Hexagonal (Ports & Adapters)** com a seguinte estrutura:

```
core/           ← domínio puro: regras, estado do jogo, validação de movimentos
ports/          ← contratos (interfaces Go)
  renderer.go   ← qualquer UI deve implementar
  player.go     ← qualquer jogador (humano ou IA) deve implementar
adapters/
  tui/          ← Bubble Tea (terminal)
  minimax/      ← IA com Minimax + Alpha-Beta Pruning
  stockfish/    ← integração via protocolo UCI
```

## Alternatives Considered

| Alternativa | Motivo de rejeição |
|---|---|
| **Layered simples** (domain/engine/ui) | Menos explícito sobre contratos. Risco de acoplamento crescer ao longo do tempo. |
| **Event-driven** | Flexível, mas complexidade desnecessária para o escopo atual. Overhead de aprendizado alto. |

## Consequences

- O `core` não tem dependências externas. Pode ser testado de forma totalmente isolada.
- Adicionar um novo adapter de IA (ex: Stockfish) não toca no `core`.
- Trocar a UI de terminal para web exige apenas um novo adapter `Renderer`.
- A interface `Player` unifica humano e IA — o `core` não sabe com quem está jogando.
- A evolução da IA (Minimax → Stockfish) é uma troca de adapter, não uma refatoração do core.
