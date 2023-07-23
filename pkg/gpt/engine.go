package gpt

import (
	"chattronics/pkg/logging"
	"context"
	"github.com/sashabaranov/go-openai"
)

type Engine interface {
	SendChat(messages []openai.ChatCompletionMessage) (ChatResponse, error)
	GetModel() string
}

type GPT struct {
	model       string
	client      *openai.Client
	temperature float32
}

type ChatResponse struct {
	Message            string
	Reason             string
	PromptTokenCount   int
	ResponseTokenCount int
}

func NewGPTEngine(model, apikey string, temperature float32) Engine {
	return &GPT{
		model:       model,
		client:      openai.NewClient(apikey),
		temperature: temperature,
	}
}

func (gpt *GPT) SendChat(messages []openai.ChatCompletionMessage) (ChatResponse, error) {
	req := openai.ChatCompletionRequest{
		Model:       gpt.model,
		Messages:    messages,
		Temperature: gpt.temperature,
	}
	resp, err := gpt.client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return ChatResponse{}, err
	}

	logging.Debug("Successfully got chat completion")

	return ChatResponse{
		Message:            resp.Choices[0].Message.Content,
		Reason:             string(resp.Choices[0].FinishReason),
		PromptTokenCount:   resp.Usage.PromptTokens,
		ResponseTokenCount: resp.Usage.CompletionTokens,
	}, nil
}

func (gpt *GPT) GetModel() string {
	return gpt.model
}
