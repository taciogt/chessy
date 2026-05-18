# ADR-0009: Abordagem de Testes — Cobertura Alta no Core, Sem Testes de TUI

## Status
Accepted

## Context

O core de xadrez (validação de movimentos, detecção de xeque) é composto de funções puras: dado um estado de tabuleiro, retorna movimentos válidos ou status do jogo. Isso é ideal para testes unitários. A TUI, por outro lado, é difícil de testar de forma automatizada e tem valor de teste baixo.

## Decision

- **Alta cobertura no `core`**: todos os movimentos de peças, detecção de xeque, xeque-mate e stalemate devem ter testes unitários com casos de tabela (table-driven tests).
- **Sem testes automáticos para a TUI**: o adapter `tui` é validado manualmente jogando.
- **Testes para adapters de IA**: o Minimax pode ser testado com posições conhecidas (ex: mate em 1 deve ser detectado).
- **Sem TDD estrito**: implementar o core primeiro, escrever testes em paralelo enquanto a API estabiliza.

## Consequences

- Os testes do core servem como documentação executável das regras do xadrez — útil para o autor aprender as regras enquanto implementa.
- Table-driven tests em Go permitem cobrir dezenas de posições de xadrez com pouco boilerplate.
- A ausência de testes de TUI é uma decisão consciente — o retorno sobre investimento é baixo para o estágio atual do projeto.
- Quando o Minimax estiver implementado, testes de posições conhecidas (puzzles de "mate em N") servem tanto para corretude quanto para regressão de performance.
