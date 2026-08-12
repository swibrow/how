package llm

import (
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/swibrow/how/internal/config"
)

func NewOllama(cfg config.OllamaConfig) (Provider, error) {
	client := openai.NewClient(
		option.WithBaseURL(cfg.URL),
		option.WithAPIKey("ollama"), // Ollama doesn't need a real key
	)

	return &openAICompatible{client: &client, model: cfg.Model, name: "ollama"}, nil
}
