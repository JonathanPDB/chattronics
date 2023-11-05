package stages

import (
	"fmt"
	"new-chattronics/internal/gpt"
	"new-chattronics/internal/interaction"
	"new-chattronics/internal/prompts"
	"strings"
)

func GetDetails(m *gpt.GPT, i interaction.User, baseMsgs gpt.Messages, categories map[string]string) (gpt.Messages, error) {
	concatenatedMessages := baseMsgs

	for name, category := range categories {
		detailsMsgs := gpt.ReplaceSystemPrompt(baseMsgs, prompts.DetailsQuestionsSystem)

		detailsMsgs = gpt.AddUserMessage(detailsMsgs, buildDetailsPrompt(name, category))
		response, err := m.SendChat(detailsMsgs)
		if err != nil {
			return nil, fmt.Errorf("failed to send request for detail questions: %w", err)
		}
		detailsMsgs = gpt.AddAssistantMessage(detailsMsgs, response)

		jsonResponse, err := interaction.ExtractJsonSlice(response)
		questions := jsonResponse["questions"]
		appMessage := fmt.Sprintf("GPT wants to ask some questions to get the details and requiremntes of the %s block. "+
			"Please answer one by one.\n\n", name)
		interaction.PrintAppMessage(appMessage)
		answers := i.AskQuestions(questions)

		interaction.PrintAppMessage("Are there any additional comments you would like to add? If not, answer \"no\".\n")
		additionComment := i.ReadConsole()
		if strings.ToLower(additionComment) != "no" {
			answers = strings.Join([]string{answers, "Additional Comment: " + additionComment}, "\n")
		}

		detailsMsgs = gpt.ReplaceSystemPrompt(detailsMsgs, prompts.GetDetailsSystemPrompt(category))
		detailsMsgs, err = detailsSatisfactionLoop(m, i, detailsMsgs, answers)
		concatenatedMessages = append(concatenatedMessages, detailsMsgs[3:]...)
	}

	return concatenatedMessages, nil
}

func detailsSatisfactionLoop(m *gpt.GPT, i interaction.User, originalMsgs gpt.Messages, userInput string) (gpt.Messages, error) {
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

		details := interaction.ExtractMarkdown(response, "details").Block
		interaction.PrintGPTMessage(details)

		isSatisfied := i.IsUserSatisfied()
		if isSatisfied {
			break
		}

		interaction.PrintAppMessage("Please provide some feedback for the model, so it can improve on the block details.\n")
		interaction.PrintAuxMessage("Feedback: ")
		userInput = prompts.DetailsFeedback + i.ReadConsole()
	}

	originalMsgs = gpt.AddUserMessage(originalMsgs, strings.Join(compiledUserInputs, "\n"))
	originalMsgs = gpt.AddAssistantMessage(originalMsgs, response)
	originalMsgs = gpt.AddUserMessage(originalMsgs, "Thank you. That is satisfactory.")

	return originalMsgs, nil
}

func buildDetailsPrompt(name, category string) string {
	prompt := "Ask questions regarding the block you called: " + name
	//prompt += prompts.GetDetailsQuestionPrompt(category)
	return prompt
}
