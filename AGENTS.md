# how

Smart terminal cheatsheet: a Go CLI that turns a natural-language question into a
shell command via an LLM (Anthropic, OpenAI, or local Ollama).

## Layout

- `cmd/how/main.go` — cobra CLI entrypoint (`how [question]`, `how config show|init`)
- `internal/config` — loads/saves `~/.config/how/config.yaml`, env vars override file
- `internal/llm` — `Provider` interface with `Anthropic`, `OpenAI`, `Ollama` implementations
- `internal/prompt` — system prompt, OS-specific hints appended at runtime
- `internal/ui` — response parsing, terminal display (lipgloss), command execution/confirmation

## Build & test

Uses `just` (see `justfile`), not `make`.

```sh
just          # lint + test + build
just test     # go test -race ./...
just lint     # golangci-lint run
just coverage # coverage report
just build    # compile ./how
```

CI (`.github/workflows/ci.yml`) runs `go test`/`go vet`/`golangci-lint-action` directly,
not through `just` — keep both in sync if you change build steps.

## Conventions

- Provider implementations satisfy `llm.Provider` (`Complete(ctx, systemPrompt, userQuery) (string, error)`).
- Config: env vars (`ANTHROPIC_API_KEY`, `OPENAI_API_KEY`) take precedence over the YAML file.
- Versioning is automated via release-please + goreleaser (Conventional Commits drive version bumps).
