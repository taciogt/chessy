# ADR-0010: Tooling — golangci-lint + Makefile + GitHub Actions

## Status
Accepted

## Context

Configurar tooling desde o início reduz débito técnico e establece boas práticas antes que o projeto cresça. O conjunto escolhido deve ser simples o suficiente para não ser um overhead no início e robusto o suficiente para durar.

## Decision

| Ferramenta | Uso |
|---|---|
| **golangci-lint** | Linter com múltiplas regras. Roda localmente e no CI. |
| **Makefile** | Atalhos para `make test`, `make run`, `make lint`. Interface única independente do ambiente. |
| **GitHub Actions** | CI que roda `lint` e `test` em todo push e PR. |

## Alternatives Considered

| Alternativa | Motivo de rejeição |
|---|---|
| **pre-commit hooks** | Pode atrapalhar o fluxo quando se está experimentando. Adicionado depois se necessário. |
| **goreleaser** | Prematuro — distribuição de binários não é prioridade agora. |
| **Taskfile (task)** | Alternativa ao Makefile, mas Makefile é mais universal e sem dependência extra. |

## Consequences

- `make test` roda todos os testes do core. `make lint` roda o golangci-lint. `make run` sobe a TUI.
- O CI garante que nenhum commit quebra testes ou introduz problemas de lint detectáveis.
- golangci-lint inclui `staticcheck`, `errcheck`, `govet` e outros — mais abrangente que `go vet` sozinho.
- pre-commit hooks ficam como opcional futuro se o autor quiser feedback local mais rápido.
