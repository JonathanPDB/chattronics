package gpt

import "github.com/sashabaranov/go-openai"

type Engine interface {
	SendChat(messages []openai.ChatCompletionMessage) (ChatResponse, error)
	GetModel() string
}
