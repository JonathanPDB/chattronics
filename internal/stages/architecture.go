package stages

import (
	"fmt"
	"github.com/chattronics/chattronics/internal/gpt"
	"github.com/chattronics/chattronics/internal/graph"
	"github.com/chattronics/chattronics/internal/interaction"
	"github.com/chattronics/chattronics/internal/prompts"
	"strings"
)

func DesignArchitecture(m *gpt.GPT, i interaction.User, msgs gpt.Messages) (gpt.Messages, error) {
	interaction.PrintAuxMessage("Project Description: ")
	description := i.ReadConsole()

	msgs = gpt.ReplaceSystemPrompt(msgs, prompts.ArchitectureQuestionsSystem)
	msgs = gpt.AddUserMessage(msgs, description)
	response, err := m.SendChat(msgs)
	if err != nil {
		return nil, fmt.Errorf("failed to send project description: %w", err)
	}
	msgs = gpt.AddAssistantMessage(msgs, response)

	jsonResponse, err := interaction.ExtractJsonSlice(response)
	questions := jsonResponse["questions"]
	interaction.PrintAppMessage("GPT wants to ask some questions to make the problem clearer. Please answer one by one.\n\n")
	answers := i.AskQuestions(questions)

	interaction.PrintAppMessage("Are there any additional comments you would like to add? If not, answer \"no\".\n")
	additionComment := i.ReadConsole()
	if strings.ToLower(additionComment) != "no" {
		answers = strings.Join([]string{answers, "Additional Comment: " + additionComment}, "\n")
	}

	msgs = gpt.ReplaceSystemPrompt(msgs, prompts.ArchitectureDesignSystem)
	return architectureSatisfactionLoop(m, i, msgs, answers)
}

func architectureSatisfactionLoop(m *gpt.GPT, i interaction.User, originalMsgs gpt.Messages, userInput string) (gpt.Messages, error) {
	var compiledUserInputs []string
	var response string
	var err error
	msgs := originalMsgs

	c := 0
	for {
		compiledUserInputs = append(compiledUserInputs, userInput)
		msgs = gpt.AddUserMessage(msgs, userInput)
		response, err = m.SendChat(msgs)
		if err != nil {
			return nil, fmt.Errorf("failed to send chat in diagram loop iteration %d: %w", c, err)
		}
		msgs = gpt.AddAssistantMessage(msgs, response)

		graphString := interaction.ExtractMarkdown(response, "Graph").Block
		explanations := interaction.ExtractMarkdown(response, "Explanations").Block

		filename := fmt.Sprintf("diagram_%d", c)
		err = graph.RenderGraph(graphString, filename)
		if err != nil {
			return nil, fmt.Errorf("failed to render graph in iteration %d: %w", c, err)
		}
		interaction.PrintGPTMessage(explanations + "\n\n")

		isSatisfied := i.IsUserSatisfied()
		if isSatisfied {
			break
		}

		interaction.PrintAppMessage("Please provide some feedback for the model, so it can improve the design.\n")
		interaction.PrintAuxMessage("Feedback: ")
		userInput = prompts.ArchitectureFeedback + i.ReadConsole()
	}

	originalMsgs = gpt.AddUserMessage(originalMsgs, strings.Join(compiledUserInputs, "\n"))
	originalMsgs = gpt.AddAssistantMessage(originalMsgs, response)
	originalMsgs = gpt.AddUserMessage(originalMsgs, "Thank you. That is satisfactory.")

	return originalMsgs, nil
}
