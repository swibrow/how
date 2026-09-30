package llm

import (
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/swibrow/how/internal/config"
)

func NewLLM(cfg config.LLMConfig) (Provider, error) {
	apiKey := cfg.APIKey
	if apiKey == "" {
		apiKey = "anything" // gateways without auth still require the SDK to send some key
	}

	client := openai.NewClient(
		option.WithBaseURL(cfg.URL),
		option.WithAPIKey(apiKey),
	)

	return &openAICompatible{client: &client, model: cfg.Model, name: "llm"}, nil
}
