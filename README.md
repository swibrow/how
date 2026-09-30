# how

A smart terminal cheatsheet — ask a natural language question, get back a shell command.

```
$ how find all go files modified in the last 24 hours

  $ find . -name '*.go' -mtime -1
  Finds all .go files modified within the last 24 hours
```

## Features

- Natural language to shell command translation
- Multiple LLM backends: **Anthropic**, **OpenAI**, **Ollama** (local), and any **OpenAI-compatible gateway** (`llm`)
- Clean, colorized terminal output
- Quiet mode for piping (`-q`)
- Optional auto-execution (`-y`)

## Installation

### Homebrew

```sh
brew install --cask swibrow/tap/how
```

### From source

```sh
go install github.com/swibrow/how/cmd/how@latest
```

### From releases

Download a prebuilt binary from [Releases](https://github.com/swibrow/how/releases).

## Usage

```sh
# Ask a question
how reverse a string in bash

# Run the suggested command immediately
how -y list listening ports

# Output only the command (useful for piping)
how -q convert png to jpg with imagemagick | sh
```

## Configuration

Initialize a config file:

```sh
how config init
```

This creates `~/.config/how/config.yaml`:

```yaml
provider: anthropic
anthropic:
  api_key: ""
  model: claude-sonnet-4-6
openai:
  api_key: ""
  model: gpt-4o
ollama:
  model: llama3
  url: http://localhost:11434/v1
llm:
  api_key: ""
  model: gpt-3.5-turbo
  url: http://localhost:4000/v1
```

### API keys

Set via environment variables (recommended) or in the config file:

```sh
export ANTHROPIC_API_KEY=sk-...
# or
export OPENAI_API_KEY=sk-...
# or
export LLM_API_KEY=sk-...
```

For **Ollama**, no API key is needed — just have Ollama running locally.

For any other **OpenAI-compatible gateway** (agentgateway, LiteLLM, vLLM, ...), set `provider: llm` and point
`url` at the gateway's OpenAI base URL, including the version prefix (e.g. `https://llm.example.com/v1`).
`how` sends requests to `<url>/chat/completions`. An API key is only required if the gateway enforces auth.

> The `litellm` provider was renamed to `llm` in v3. Rename `provider: litellm` and the `litellm:` section to
> `llm`, add `/v1` to `url` if your gateway needs it, and use `LLM_API_KEY` instead of `LITELLM_API_KEY`.

### View current config

```sh
how config show
```

## Development

### Prerequisites

- Go 1.25+
- [golangci-lint](https://golangci-lint.run/) (for linting)

### Build & test

```sh
just          # lint + test + build
just test     # run tests with race detector
just lint     # run linters
just coverage # generate coverage report
just build    # compile binary
```

## License

MIT
