package gpt

import "github.com/sashabaranov/go-openai"

// Roles
const (
	UserRole      = openai.ChatMessageRoleUser
	SystemRole    = openai.ChatMessageRoleSystem
	AssistantRole = openai.ChatMessageRoleAssistant
)

// Models
const (
	GPT3_5Turbo = openai.GPT3Dot5Turbo
	GPT4        = openai.GPT4
	GPT4Turbo   = openai.GPT4TurboPreview

	GPT4Prefix      = "gpt-4"
	GPT4TurboPrefix = "gpt-4-1106"
	GPT3_5Prefix    = "gpt-3.5"
)

// Costs
const (
	GPT3_5CostInput     = 0.0015 // USD
	GPT3_5CostOutput    = 0.002  //  USD
	GPT4CostInput       = 0.03   // USD
	GPT4CostOutput      = 0.06   // USD
	GPT4TurboCostInput  = 0.01   // USD
	GPT4TurboCostOutput = 0.03   // USD
)
