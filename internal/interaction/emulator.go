package interaction

import (
	"fmt"
	"github.com/chattronics/chattronics/internal/gpt"
	"github.com/chattronics/chattronics/internal/logging"
	"strconv"
	"strings"
)

const (
	emulatorGptModel       = gpt.GPT4Turbo
	emulatorGptTemperature = 0.5
)

type Emulator struct {
	projectPrompt string
	systemPrompt  string
	counter       int
	gpt           *gpt.GPT
}

func CreateEmulatorUser(projectPrompt, systemPrompt, apiKey string) User {
	return &Emulator{
		projectPrompt: projectPrompt,
		systemPrompt:  systemPrompt,
		counter:       0,
		gpt:           gpt.NewGPT(emulatorGptModel, apiKey, emulatorGptTemperature, "emulator"),
	}
}

func (g *Emulator) ReadConsole() string {
	if g.counter == 0 {
		g.counter++
		return g.projectPrompt
	}

	return "no"
}

func (g *Emulator) AskQuestions(questions []string) string {
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

func (g *Emulator) IsUserSatisfied() bool {
	return true
}
