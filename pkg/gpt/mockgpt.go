package gpt

import (
	"github.com/JonathanPDB/chattronics/pkg/prompts"
	"github.com/sashabaranov/go-openai"
)

type MockGPT struct {
	model     string
	responses []ChatResponse
	counter   int
}

func NewMockGPT(model string) Engine {
	return &MockGPT{
		model:     model,
		responses: initializeMockResponses(),
		counter:   0,
	}
}

func (m *MockGPT) SendChat(messages []openai.ChatCompletionMessage) (ChatResponse, error) {
	i := m.counter
	m.counter++
	return m.responses[i], nil
}

func (m *MockGPT) GetModel() string {
	return m.model
}

func initializeMockResponses() []ChatResponse {
	return []ChatResponse{{
		Message:            prompts.MockQuestionsResponse,
		Reason:             string(openai.FinishReasonStop),
		ResponseTokenCount: 100,
		PromptTokenCount:   200,
	}, {
		Message:            prompts.MockIdealPromptResponse,
		Reason:             string(openai.FinishReasonStop),
		ResponseTokenCount: 80,
		PromptTokenCount:   120,
	}, {
		Message:            prompts.MockDiagramResponse,
		Reason:             string(openai.FinishReasonStop),
		ResponseTokenCount: 50,
		PromptTokenCount:   80,
	}}
}
