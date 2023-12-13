package interaction

import (
	"fmt"
	"github.com/chattronics/chattronics/internal/gpt"
	"github.com/chattronics/chattronics/internal/logging"
	"strconv"
	"strings"
)

const (
	gptUserModel       = gpt.GPT4Turbo
	gptUserTemperature = 0.2
)

type GPTUser struct {
	projectPrompt string
	systemPrompt  string
	counter       int
	gpt           *gpt.GPT
}

func CreateGPTUser(projectPrompt, systemPrompt, apiKey string) User {
	return &GPTUser{
		projectPrompt: projectPrompt,
		systemPrompt:  systemPrompt,
		counter:       0,
		gpt:           gpt.NewGPT(gptUserModel, apiKey, gptUserTemperature, "mock_user"),
	}
}

func (g *GPTUser) ReadConsole() string {
	if g.counter == 0 {
		g.counter++
		return g.projectPrompt
	}

	return "no"
}

func (g *GPTUser) AskQuestions(questions []string) string {
	var answers []string
	systemPrompt := gpt.ReplaceSystemPrompt(gpt.Messages{}, g.systemPrompt)

	for i, question := range questions {
		questionNumber := strconv.Itoa(i+1) + ". "

		msgs := gpt.AddUserMessage(systemPrompt, question)
		response, err := g.gpt.SendChat(msgs)
		if err != nil {
			logging.Fatal(fmt.Sprintf("failed to send gpt user message: %v", err))
		}

		if strings.EqualFold(response, "cannot answer question") {
			answers = append(answers, questionNumber+response)
		} else {
			answer := ExtractMarkdown(response, "answer").Block
			answers = append(answers, questionNumber+answer)
		}

	}
	return strings.Join(answers, "\n")
}

func (g *GPTUser) IsUserSatisfied() bool {
	return true
}
