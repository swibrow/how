package llm

import (
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/swibrow/how/internal/config"
)

func NewLiteLLM(cfg config.LiteLLMConfig) (Provider, error) {
	apiKey := cfg.APIKey
	if apiKey == "" {
		apiKey = "anything" // litellm proxy accepts any key unless a master_key is configured
	}

	client := openai.NewClient(
		option.WithBaseURL(cfg.URL),
		option.WithAPIKey(apiKey),
	)

	return &openAICompatible{client: &client, model: cfg.Model, name: "litellm"}, nil
}
