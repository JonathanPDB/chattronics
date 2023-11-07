package gpt

import (
	"context"
	"fmt"
	"github.com/chattronics/chattronics/internal/logging"
	"github.com/sashabaranov/go-openai"
	"strings"
)

type Messages []openai.ChatCompletionMessage

type GPT struct {
	model       string
	client      *openai.Client
	temperature float32
	prompt      Messages
	totalCost   float64
	logger      *logging.ChatLogger
}

func NewGPT(model, apikey string, temperature float32, persona string) *GPT {
	return &GPT{
		model:       model,
		client:      openai.NewClient(apikey),
		temperature: temperature,
		totalCost:   0,
		logger:      logging.NewChatLogger(persona),
	}
}

func (gpt *GPT) SendChat(messages Messages) (string, error) {
	req := openai.ChatCompletionRequest{
		Model:       gpt.model,
		Messages:    messages,
		Temperature: gpt.temperature,
	}
	resp, err := gpt.client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return "", fmt.Errorf("failed to get OpenAPI chat completion: %w", err)
	}
	logging.Debug("Successfully got chat completion")

	gpt.totalCost += gpt.calculateCost(resp.Usage.PromptTokens, resp.Usage.CompletionTokens)

	if resp.Choices[0].FinishReason != openai.FinishReasonStop {
		logging.Warn("Finish Reason was not stop", logging.AddField("finish_reason", string(resp.Choices[0].FinishReason)))
	}

	err = gpt.log(messages, resp)
	if err != nil {
		return "", fmt.Errorf("failed to log chat: %w", err)
	}

	return resp.Choices[0].Message.Content, nil
}

func (gpt *GPT) calculateCost(inputTokens, outputTokens int) float64 {
	var inputCost, outputCost float64

	if strings.HasPrefix(gpt.model, GPT4Prefix) {
		inputCost = (float64(inputTokens) * GPT4CostInput) / 1000
		outputCost = (float64(outputTokens) * GPT4CostOutput) / 1000
	} else if strings.HasPrefix(gpt.model, GPT3_5Prefix) {
		inputCost = (float64(inputTokens) * GPT3_5CostInput) / 1000
		outputCost = (float64(outputTokens) * GPT3_5CostOutput) / 1000
	} else {
		logging.Warn("Failed to find model prefix to calculate cost")
		return 0.0
	}
	return inputCost + outputCost
}

func (gpt *GPT) log(prompt Messages, r openai.ChatCompletionResponse) error {
	return gpt.logger.LogInteraction(logging.InteractionLog{
		Prompt:          prompt,
		ResponseMessage: r.Choices[0].Message.Content,
		Cost:            gpt.totalCost,
		FinishReason:    string(r.Choices[0].FinishReason),
	})
}
