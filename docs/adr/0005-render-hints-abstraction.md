# ADR-0005: Abstração RenderHints para Overlays Visuais

## Status
Accepted

## Context

O projeto precisa suportar: (1) preview de movimentos antes da confirmação, (2) modos de visualização para iniciantes (casas sob ataque, peças que podem se mover, etc.), (3) modo debug. Essas funcionalidades são diferentes em aparência mas idênticas em estrutura: o core calcula informação extra e a UI decide como exibir.

## Decision

O port `Renderer` recebe um `RenderHints` junto com o `GameState` a cada renderização:

```go
type RenderHints struct {
    PreviewMove      *Move
    HighlightSquares []Square
}

type Renderer interface {
    Render(state GameState, hints RenderHints)
}
```

O `core` é responsável por calcular os hints (quais casas destacar, qual movimento pré-visualizar). O adapter `Renderer` é responsável apenas por como exibir essa informação.

## Alternatives Considered

| Alternativa | Motivo de rejeição |
|---|---|
| **Lógica de highlight no adapter** | O adapter precisaria entender regras de xadrez para calcular casas sob ataque. Viola a separação de responsabilidades. |
| **Eventos separados** (`OnPreview`, `OnHighlight`) | Mais granular, mas adiciona complexidade de protocolo desnecessária para o escopo atual. |
| **Renderer especializado por modo** | Criaria múltiplas interfaces em vez de uma única extensível. |

## Consequences

- O modo "iniciante" (mostra casas sob ataque, peças movíveis) e o preview de movimentos são a mesma abstração — apenas com hints diferentes.
- O adapter TUI de terminal usa cores/caracteres para destacar. Um futuro adapter web usaria classes CSS ou ícones. O contrato é o mesmo.
- Adicionar novos tipos de hint (ex: `LastMove`, `Threats`) é uma expansão do struct, não uma mudança de interface.
- O `core` precisa implementar cálculo de casas atacadas mesmo para o modo de visualização — o que de qualquer forma é necessário para validação de movimentos.
