package stages

import (
	"fmt"
	"new-chattronics/internal/gpt"
	"new-chattronics/internal/graph"
	"new-chattronics/internal/interaction"
	"new-chattronics/internal/prompts"
	"strings"
)

func DesignArchitecture(m *gpt.GPT, msgs gpt.Messages) (gpt.Messages, error) {
	interaction.PrintAuxMessage("Project Description: ")
	description := interaction.ReadConsole()

	msgs = gpt.ReplaceSystemPrompt(msgs, prompts.ArchitectureQuestionsSystem)
	msgs = gpt.AddUserMessage(msgs, description)
	response, err := m.SendChat(msgs)
	if err != nil {
		return nil, fmt.Errorf("failed to send project description: %w", err)
	}
	msgs = gpt.AddAssistantMessage(msgs, response)

	questions := interaction.ExtractSingleBlock(response, "Questions").Lines
	interaction.PrintAppMessage("GPT wants to ask some questions to make the problem clearer. Please answer one by one.\n\n")
	answers := interaction.AskQuestions(questions)

	interaction.PrintAppMessage("Are there any additional comments you would like to add? If not, answer \"no\".\n")
	additionComment := interaction.ReadConsole()
	if strings.ToLower(additionComment) != "no" {
		answers = strings.Join([]string{answers, "Additional Comment: " + additionComment}, "\n")
	}

	msgs = gpt.ReplaceSystemPrompt(msgs, prompts.ArchitectureDesignSystem)
	return architectureSatisfactionLoop(m, msgs, answers)
}

func architectureSatisfactionLoop(m *gpt.GPT, originalMsgs gpt.Messages, userInput string) (gpt.Messages, error) {
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

		graphString := interaction.ExtractSingleBlock(response, "Graph").Block
		explanations := interaction.ExtractSingleBlock(response, "Explanations").Block

		filename := fmt.Sprintf("diagram_%d", c)
		err = graph.RenderGraph(graphString, filename)
		if err != nil {
			return nil, fmt.Errorf("failed to render graph in iteration %d: %w", c, err)
		}
		interaction.PrintGPTMessage(explanations + "\n\n")

		isSatisfied := interaction.IsUserSatisfied()
		if isSatisfied {
			break
		}

		interaction.PrintAppMessage("Please provide some feedback for the model, so it can improve the design.\n")
		interaction.PrintAuxMessage("Feedback: ")
		userInput = prompts.ArchitectureFeedback + interaction.ReadConsole()
	}

	originalMsgs = gpt.AddUserMessage(originalMsgs, strings.Join(compiledUserInputs, "\n"))
	originalMsgs = gpt.AddAssistantMessage(originalMsgs, response)
	originalMsgs = gpt.AddUserMessage(originalMsgs, "Thank you. That is satisfactory.")

	return originalMsgs, nil
}
