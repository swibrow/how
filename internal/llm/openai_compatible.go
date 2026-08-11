package llm

import (
	"context"
	"fmt"

	"github.com/openai/openai-go"
)

// openAICompatible implements Provider for any backend that speaks the
// OpenAI chat completions API (OpenAI itself, Ollama, LiteLLM, etc).
type openAICompatible struct {
	client *openai.Client
	model  string
	name   string // used in error messages, e.g. "openai", "ollama", "litellm"
}

func (p *openAICompatible) Complete(ctx context.Context, systemPrompt, userQuery string) (string, error) {
	resp, err := p.client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Model: p.model,
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(systemPrompt),
			openai.UserMessage(userQuery),
		},
	})
	if err != nil {
		return "", fmt.Errorf("%s API error: %w", p.name, err)
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("%s returned no choices", p.name)
	}

	return resp.Choices[0].Message.Content, nil
}
