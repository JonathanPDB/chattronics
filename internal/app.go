package internal

import (
	"fmt"
	"github.com/chattronics/chattronics/internal/gpt"
	"github.com/chattronics/chattronics/internal/interaction"
	"github.com/chattronics/chattronics/internal/stages"
)

func RunApp(m *gpt.GPT, i interaction.User) (string, error) {
	msgs := make(gpt.Messages, 0)
	var err error

	interaction.GreetingsMessage()

	msgs, err = stages.DesignArchitecture(m, i, msgs)
	if err != nil {
		return "", fmt.Errorf("failed to run architecture stage: %w", err)
	}

	categories, err := stages.Categorize(m, msgs)
	if err != nil {
		return "", fmt.Errorf("failed to run categorization stage: %w", err)
	}

	msgs, err = stages.GetDetails(m, i, msgs, categories)
	if err != nil {
		return "", fmt.Errorf("failed to run details stage: %w", err)
	}

	compilation, err := stages.GenerateSummary(m, msgs)
	if err != nil {
		return "", fmt.Errorf("failed to run summary stage: %w", err)
	}

	return compilation, nil
}
