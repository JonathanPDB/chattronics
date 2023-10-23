package internal

import (
	"fmt"
	"github.com/sashabaranov/go-openai"
	"new-chattronics/internal/gpt"
	"new-chattronics/internal/interaction"
	"new-chattronics/internal/stages"
)

func RunApp(m *gpt.GPT) error {
	msgs := make([]openai.ChatCompletionMessage, 0)
	var err error

	interaction.GreetingsMessage()

	msgs, err = stages.DesignArchitecture(m, msgs)
	if err != nil {
		return fmt.Errorf("failed to run architecture stage: %w", err)
	}

	categories, err := stages.Categorize(m, msgs)
	if err != nil {
		return fmt.Errorf("failed to run categorization stage: %w", err)
	}

	msgs, err = stages.GetDetails(m, msgs, categories)
	if err != nil {
		return fmt.Errorf("failed to run details stage: %w", err)
	}

	err = stages.GenerateSummary(m, msgs)
	if err != nil {
		return fmt.Errorf("failed to run summary stage: %w", err)
	}

	return nil
}
