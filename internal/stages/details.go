package stages

import (
	"fmt"
	"new-chattronics/internal/gpt"
	"new-chattronics/internal/interaction"
	"new-chattronics/internal/prompts"
	"strings"
)

func GetDetails(m *gpt.GPT, msgs gpt.Messages, categories map[string]string) (gpt.Messages, error) {

	for name, category := range categories {
		msgs = gpt.ReplaceSystemPrompt(msgs, prompts.DetailsQuestionsSystem)

		msgs = gpt.AddUserMessage(msgs, prompts.GetDetailsQuestionPrompt(category))
		response, err := m.SendChat(msgs)
		if err != nil {
			return nil, fmt.Errorf("failed to send request for detail questions: %w", err)
		}
		msgs = gpt.AddAssistantMessage(msgs, response)

		questions := interaction.ExtractSingleBlock(response, "Questions").Lines
		appMessage := fmt.Sprintf("GPT wants to ask some questions to get the details and requiremntes of the %s block. "+
			"Please answer one by one.\n\n", name)
		interaction.PrintAppMessage(appMessage)
		answers := interaction.AskQuestions(questions)

		interaction.PrintAppMessage("Are there any additional comments you would like to add? If not, answer \"no\".\n")
		additionComment := interaction.ReadConsole()
		if strings.ToLower(additionComment) != "no" {
			answers = strings.Join([]string{answers, "Additional Comment: " + additionComment}, "\n")
		}

		msgs = gpt.ReplaceSystemPrompt(msgs, prompts.GetDetailsSystemPrompt(category))
		msgs, err = detailsSatisfactionLoop(m, msgs, answers)
	}

	return msgs, nil
}

func detailsSatisfactionLoop(m *gpt.GPT, originalMsgs gpt.Messages, userInput string) (gpt.Messages, error) {
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
			return nil, fmt.Errorf("failed to send chat in details loop iteration %d: %w", c, err)
		}
		msgs = gpt.AddAssistantMessage(msgs, response)

		details := interaction.ExtractSingleBlock(response, "details").Block
		interaction.PrintGPTMessage(details)

		isSatisfied := interaction.IsUserSatisfied()
		if isSatisfied {
			break
		}

		interaction.PrintAppMessage("Please provide some feedback for the model, so it can improve on the block details.\n")
		interaction.PrintAuxMessage("Feedback: ")
		userInput = prompts.DetailsFeedback + interaction.ReadConsole()
	}

	originalMsgs = gpt.AddUserMessage(originalMsgs, strings.Join(compiledUserInputs, "\n"))
	originalMsgs = gpt.AddAssistantMessage(originalMsgs, response)
	originalMsgs = gpt.AddUserMessage(originalMsgs, "Thank you. That is satisfactory.")

	return originalMsgs, nil
}
