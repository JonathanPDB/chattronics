package testdetails

import (
	"fmt"
	"github.com/chattronics/chattronics/internal/config"
	"github.com/chattronics/chattronics/internal/gpt"
	"github.com/chattronics/chattronics/internal/interaction"
	"github.com/chattronics/chattronics/internal/logging"
	"github.com/chattronics/chattronics/internal/stages"
	"github.com/chattronics/chattronics/testbenches/cases"
	"github.com/chattronics/chattronics/testbenches/evaluate"
	"github.com/stretchr/testify/assert"
	"testing"
)

const testCase = "thermometry"

func TestGetDetails(t *testing.T) {
	_ = config.CreateFolders("unit-test-details_", false)
	logging.InitializeStandardLogger()
	apiKey := config.LoadApiKeyEnvVar()

	projectDesc, informativeSysPrompt, err := cases.GetInformative(testCase)
	if err != nil {
		logging.Fatal("Failed to get limited scenario information", logging.AddField("error", err))
		return
	}
	user := interaction.CreateGPTUser(projectDesc, informativeSysPrompt, apiKey)
	_ = user.ReadConsole()

	gptModel := gpt.NewGPT(gpt.GPT4Turbo, apiKey, float32(0.8), "test-engineer")

	msgs, err := parseConversation(testCase)
	assert.NoError(t, err)

	t.Run("Amplification", func(t *testing.T) {
		categories := map[string]string{
			"C": "Amplification",
		}

		detailsMsgs, err := stages.GetDetails(gptModel, user, msgs, categories)
		assert.NoError(t, err)

		requirements := "1. Explicitly informed gain value."
		expectedScore := 1

		summary, err := stages.GenerateSummary(gptModel, detailsMsgs)

		score, verdict, explanation, err := evaluate.Evaluate(summary, requirements, apiKey)

		fmt.Println(verdict)
		fmt.Println(explanation)

		assert.NoError(t, err)
		assert.Equal(t, expectedScore, score)
	})

}
