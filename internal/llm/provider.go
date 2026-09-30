package llm

import (
	"context"
	"errors"
	"fmt"

	"github.com/swibrow/how/internal/config"
)

// Provider defines the interface for LLM backends.
type Provider interface {
	Complete(ctx context.Context, systemPrompt, userQuery string) (string, error)
}

// NewProvider creates a provider based on the config.
func NewProvider(cfg *config.Config) (Provider, error) {
	switch cfg.Provider {
	case "anthropic":
		return NewAnthropic(cfg.Anthropic)
	case "openai":
		return NewOpenAI(cfg.OpenAI)
	case "ollama":
		return NewOllama(cfg.Ollama)
	case "llm":
		return NewLLM(cfg.LLM)
	case "litellm":
		return nil, errors.New(`provider "litellm" was renamed to "llm": set provider: llm, rename the litellm: config section to llm:, and use LLM_API_KEY instead of LITELLM_API_KEY`)
	default:
		return nil, fmt.Errorf("unknown provider: %s", cfg.Provider)
	}
}
