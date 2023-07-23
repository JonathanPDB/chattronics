package gpt

import (
	"github.com/sashabaranov/go-openai"
)

type MockGPT struct {
	model string
}

func NewMockGPT(model string) Engine {
	return &MockGPT{
		model: model,
	}
}

func (m *MockGPT) SendChat(messages []openai.ChatCompletionMessage) (ChatResponse, error) {
	return ChatResponse{
		Message:            "mock response",
		Reason:             string(openai.FinishReasonStop),
		PromptTokenCount:   500,
		ResponseTokenCount: 150,
	}, nil
}

func (m *MockGPT) GetModel() string {
	return m.model
}
