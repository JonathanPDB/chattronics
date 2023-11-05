package gpt

import (
	"github.com/sashabaranov/go-openai"
	"new-chattronics/internal/logging"
)

func AddUserMessage(msgs Messages, new string) Messages {
	//if msgs[len(msgs)-1].Role == UserRole {
	//	logging.Warn("adding back-to-back user prompts")
	//}

	msgs = append(msgs, openai.ChatCompletionMessage{
		Role:    UserRole,
		Content: new,
	})

	return msgs
}

func AddAssistantMessage(msgs Messages, new string) Messages {
	//if msgs[len(msgs)-1].Role == AssistantRole {
	//	logging.Warn("adding back-to-back assistant prompts")
	//}

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
