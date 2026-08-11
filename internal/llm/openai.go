package llm

import (
	"fmt"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/swibrow/how/internal/config"
)

func NewOpenAI(cfg config.OpenAIConfig) (Provider, error) {
	if cfg.APIKey == "" {
		return nil, fmt.Errorf("openai API key not set (set OPENAI_API_KEY or configure in ~/.config/how/config.yaml)")
	}

	client := openai.NewClient(option.WithAPIKey(cfg.APIKey))

	return &openAICompatible{client: &client, model: cfg.Model, name: "openai"}, nil
}
