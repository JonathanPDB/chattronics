package gpt

import (
	"github.com/chattronics/chattronics/internal/logging"
	"github.com/sashabaranov/go-openai"
)

func AddUserMessage(msgs Messages, new string) Messages {
	msgs = append(msgs, openai.ChatCompletionMessage{
		Role:    UserRole,
		Content: new,
	})

	return msgs
}

func AddAssistantMessage(msgs Messages, new string) Messages {
	msgs = append(msgs, openai.ChatCompletionMessage{
		Role:    AssistantRole,
		Content: new,
	})

	return msgs
}

func ReplaceSystemPrompt(msgs Messages, new string) Messages {
	if len(msgs) == 0 {
		msgs = append(msgs, openai.ChatCompletionMessage{
			Role:    SystemRole,
			Content: new,
		})
	}

	if msgs[0].Role != SystemRole {
		logging.Fatal("First message in messages is not a System prompt!")
	}

	msgs[0] = openai.ChatCompletionMessage{
		Role:    SystemRole,
		Content: new,
	}
	return msgs
}
