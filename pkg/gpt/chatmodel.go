package gpt

import (
	"fmt"
	"github.com/JonathanPDB/chattronics/pkg/logging"
	"github.com/sashabaranov/go-openai"
	"strings"
)

type ChatModel struct {
	engine    Engine
	prompt    []openai.ChatCompletionMessage
	totalCost float64
	logger    *logging.ChatLogger
}

type ChatResponse struct {
	Message            string
	Reason             string
	PromptTokenCount   int
	ResponseTokenCount int
}

func NewChatModel(engine Engine, alias, systemPrompt string) ChatModel {
	return ChatModel{
		engine: engine,
		prompt: []openai.ChatCompletionMessage{{
			Role:    SystemRole,
			Content: systemPrompt,
		}},
		logger: logging.NewChatLogger(alias),
	}
}

func (m *ChatModel) AddUserMessage(message string) {
	if m.prompt[len(m.prompt)-1].Role == UserRole {
		logging.Warn("adding back-to-back user prompts")
	}

	m.prompt = append(m.prompt, openai.ChatCompletionMessage{
		Role:    UserRole,
		Content: message,
	})
}

func (m *ChatModel) SendChat() (string, error) {
	resp, err := m.engine.SendChat(m.prompt)
	if err != nil {
		return "", fmt.Errorf("failed to send chat: %w", err)
	}

	if resp.Reason != string(openai.FinishReasonStop) {
		logging.Warn("Finish Reason was not stop", logging.AddField("finish_reason", resp.Reason))
	}

	m.totalCost += m.calculateCost(resp.PromptTokenCount, resp.ResponseTokenCount)

	err = m.log(m.prompt, resp)
	if err != nil {
		return "", fmt.Errorf("failed to log chat: %w", err)
	}

	m.prompt = append(m.prompt, openai.ChatCompletionMessage{
		Role:    AssistantRole,
		Content: resp.Message,
	})

	return resp.Message, nil
}

func (m *ChatModel) calculateCost(inputTokens, outputTokens int) float64 {
	if strings.HasPrefix(m.engine.GetModel(), GPT4Prefix) {
		inputCostUSD := (float64(inputTokens) * GPT4CostInput) / 1000
		outputCostUSD := (float64(outputTokens) * GPT4CostOutput) / 1000
		return getBRLfromUSD(inputCostUSD + outputCostUSD)
	}

	if strings.HasPrefix(m.engine.GetModel(), GPT3_5Prefix) {
		inputCostUSD := (float64(inputTokens) * GPT3_5CostInput) / 1000
		outputCostUSD := (float64(outputTokens) * GPT3_5CostOutput) / 1000
		return getBRLfromUSD(inputCostUSD + outputCostUSD)
	}
	logging.Warn("Failed to find model prefix to calculate cost")
	return 0.0
}

func (m *ChatModel) log(prompt []openai.ChatCompletionMessage, r ChatResponse) error {
	return m.logger.LogInteraction(logging.InteractionLog{
		Prompt:           prompt,
		ResponseMessage:  r.Message,
		InputTokenCount:  r.PromptTokenCount,
		OutputTokenCount: r.ResponseTokenCount,
		Cost:             m.calculateCost(r.PromptTokenCount, r.ResponseTokenCount),
		FinishReason:     r.Reason,
	})
}
