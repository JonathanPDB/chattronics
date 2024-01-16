package stages

import (
	"fmt"
	"github.com/chattronics/chattronics/internal/config"
	"github.com/chattronics/chattronics/internal/gpt"
	"github.com/chattronics/chattronics/internal/interaction"
	"github.com/chattronics/chattronics/internal/prompts"
	"os"
	"strings"
)

const summaryFileName = "summary.md"
const compilationFileName = "compilation.md"

func GenerateSummary(m *gpt.GPT, msgs gpt.Messages) (string, error) {
	interaction.PrintAppMessage("\nThank you for using the app! Generating Summary.\n\n\n")

	msgs = gpt.ReplaceSystemPrompt(msgs, prompts.ReviewSolution)
	response, err := m.SendChat(msgs)
	if err != nil {
		return "", fmt.Errorf("failed to send summarize message")
	}
	review := interaction.ExtractMarkdown(response, "review").Block
	interaction.PrintAppMessage("GPT has made some corrections to his implementation after review it as a whole:\n")
	interaction.PrintGPTMessage(review)

	summaryMsgs := gpt.AddUserMessage(msgs, prompts.Summary)
	response, err = m.SendChat(summaryMsgs)
	if err != nil {
		return "", fmt.Errorf("failed to send summarize message")
	}
	summary := interaction.ExtractMarkdown(response, "summary").Block

	file := config.RunFolderPath + summaryFileName
	err = os.WriteFile(file, []byte(summary), 0666)
	if err != nil {
		return "", fmt.Errorf("error writing summary to file: %w", err)
	}

	interaction.PrintGPTMessage("SUMMARY\n\n" + summary)

	//compilation := compileAssistantMessages(msgs)
	//file = config.RunFolderPath + compilationFileName
	//err = os.WriteFile(file, []byte(compilation), 0666)
	//if err != nil {
	//	return "", fmt.Errorf("error writing compilation to file: %w", err)
	//}

	return summary, err
}

func compileAssistantMessages(msgs gpt.Messages) string {
	var compiledSolution string
	for _, message := range msgs {
		if message.Role != gpt.AssistantRole {
			continue
		}
		if !strings.Contains(message.Content, "```") {
			continue
		}
		compiledSolution += message.Content + "\n"
	}
	return compiledSolution
}
