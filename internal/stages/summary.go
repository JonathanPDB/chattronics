package stages

import (
	"fmt"
	"github.com/chattronics/chattronics/internal/config"
	"github.com/chattronics/chattronics/internal/gpt"
	"github.com/chattronics/chattronics/internal/interaction"
	"github.com/chattronics/chattronics/internal/prompts"
	"os"
)

const summaryFileName = "summary.txt"

func GenerateSummary(m *gpt.GPT, msgs gpt.Messages) (string, error) {
	interaction.PrintAppMessage("Thank you for using the app! Generating Summary.\n")

	msgs = gpt.AddUserMessage(msgs, prompts.Summary)
	response, err := m.SendChat(msgs)
	if err != nil {
		return "", fmt.Errorf("failed to send summarize message")
	}
	summary := interaction.ExtractMarkdown(response, "summary").Block

	interaction.PrintAuxMessage(summary)

	file := config.RunFolderPath + summaryFileName
	err = os.WriteFile(file, []byte(summary), 0666)
	if err != nil {
		return "", fmt.Errorf("error writing summary to file: %w", err)
	}

	return summary, err
}
