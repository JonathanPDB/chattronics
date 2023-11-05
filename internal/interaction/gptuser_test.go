package interaction

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"new-chattronics/internal/config"
	"new-chattronics/internal/logging"
	"new-chattronics/tests/cases"
	"slices"
	"strings"
	"testing"
)

func TestGPTUser_ReadConsole(t *testing.T) {
	config.CreateFolders("", false)
	logging.InitializeStandardLogger()

	t.Run("Successfully return project description only in the first iteration", func(t *testing.T) {
		givenProjectDesc := "Mock project description"

		user := CreateGPTUser(givenProjectDesc, "Mock system prompt", "123")
		resp := user.ReadConsole()
		assert.Equal(t, givenProjectDesc, resp)

		for i := 0; i < 5; i++ {
			resp = user.ReadConsole()
			assert.Equal(t, "no", resp)
		}
	})
}

func TestGPTUser_AskQuestions(t *testing.T) {
	config.CreateFolders("", false)
	apiKey := config.LoadApiKeyEnvVar()
	logging.InitializeStandardLogger()

	iterations := 3

	baseSysPrompt := cases.BaseInformative

	type testCase struct {
		sysPrompt                   string
		questions                   []string
		expectedUnansweredQuestions []int
	}

	testTable := []testCase{
		{
			sysPrompt: baseSysPrompt + "" +
				"1. The maximum error should be 1%\n" +
				"2. The water is between 15 and 85 degrees celsius\n" +
				"3. The voltage range should be between 0 and 20 Volts\n" +
				"4. Self heating should be below 5%",
			questions: []string{
				"1. What is the preferred method of linearization?",
				"2. What should the expected temperature range?",
				"3. Should the self heating effect be accounted for?",
				"4. Should there be a digital filtering stage?",
			},
			expectedUnansweredQuestions: []int{0, 3},
		},
	}

	for k, test := range testTable {
		fmt.Printf("\nStating Testcase %d\n", k)
		user := CreateGPTUser("", test.sysPrompt, apiKey)

		for i := 0; i < iterations; i++ {
			fmt.Printf("Iteration %d\n", i)

			answersBlock := user.AskQuestions(test.questions)
			answers := strings.Split(answersBlock, "\n")

			assert.Len(t, answers, len(test.questions))

			for j := 0; j < len(answers); j++ {
				fmt.Printf("Question %d\n", j)
				if slices.Contains(test.expectedUnansweredQuestions, j) {
					assert.Contains(t, strings.ToLower(answers[j]), "cannot answer question")
				} else {
					assert.NotContains(t, strings.ToLower(answers[j]), "cannot answer question")
				}
			}
		}
	}
}
