package chat

import (
	"chattronics/pkg/gpt"
	"chattronics/pkg/logging"
	"chattronics/pkg/prompts"
	"fmt"
	"strings"
)

func StartConversation(gptEngine *gpt.Engine) error {
	Inquisitor := gpt.NewChatModel(*gptEngine, "inquisitor", prompts.InquistorSystem)

	fmt.Println("Send message")
	Inquisitor.AddUserMessage(readConsole())

	response, err := Inquisitor.SendChat()
	if err != nil {
		return err
	}

	questionBlocks := ExtractMarkdownBlocks(response)
	if len(questionBlocks) != 1 {
		logging.Warn("Received two code blocks, when only 1 was expected")
	}

	questionBlock := questionBlocks[0]
	if questionBlock.Language != "Questions" {
		logging.Warn("GPT did not follow 'Questions' language name in markdown block")
	}

	answers := askQuestions(questionBlock.Lines)

	newUserPrompt := prompts.FormulatePrefix + answers + prompts.FormulateSuffix
	Inquisitor.AddUserMessage(newUserPrompt)

	idealPromptResponse, err := Inquisitor.SendChat()
	if err != nil {
		return err
	}

	idealPrompt := ExtractMarkdownBlocks(idealPromptResponse)[0].Block

	EngineerTopDown := gpt.NewChatModel(*gptEngine, "engineerTopDown", prompts.TopDownSystem)
	EngineerTopDown.AddUserMessage(idealPrompt)

	architectureReponse, err := EngineerTopDown.SendChat()
	if err != nil {
		return err
	}

	yuml := ExtractMarkdownBlocks(architectureReponse)[0].Block

	fmt.Println(yuml)
	return nil
}

func askQuestions(questions []string) string {
	var answers []string
	fmt.Println("GPT wants to ask some questions to make the problem clearer. Please answer one by one.")
	fmt.Println("")
	for i, question := range questions {
		fmt.Println(question)
		fmt.Printf("Answer: ")
		questionNumber := string(rune(i+1)) + ". "
		answers = append(answers, questionNumber+readConsole())
	}

	return strings.Join(answers, "\n")
}
