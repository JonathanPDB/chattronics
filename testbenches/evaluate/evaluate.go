package evaluate

import (
	"fmt"
	"github.com/chattronics/chattronics/internal/config"
	"github.com/chattronics/chattronics/internal/gpt"
	"github.com/chattronics/chattronics/internal/interaction"
	"github.com/chattronics/chattronics/testbenches/cases"
	"os"
	"strconv"
)

const scoresFileName = "scores.md"
const verdictFileName = "verdict.md"

func GiveVerdict(compilation, requirements, apiKey string) (string, error) {
	evaluatorSystemPrompt := fmt.Sprintf(cases.BaseVerdictEvaluator, requirements)
	evaluatorMsgs := gpt.ReplaceSystemPrompt(gpt.Messages{}, evaluatorSystemPrompt)
	evaluatorMsgs = gpt.AddUserMessage(evaluatorMsgs, compilation)

	gptModel := gpt.NewGPT(gpt.GPT4Turbo, apiKey, 0.6, "verdict")

	response, err := gptModel.SendChat(evaluatorMsgs)
	if err != nil {
		return "", fmt.Errorf("failed to get verdict from gpt: %w", err)
	}

	verdict := interaction.ExtractMarkdown(response, "verdict").Block
	if err != nil {
		return "", fmt.Errorf("failed to extract verdict: %v", err)
	}

	explanations := interaction.ExtractMarkdown(response, "explanations").Block
	if err != nil {
		return "", fmt.Errorf("failed to extract explanations: %v", err)
	}

	evaluation := fmt.Sprintf(
		"Verdict: %s\n\n"+
			"Explanations: \n%s",
		verdict, explanations,
	)

	file := config.RunFolderPath + verdictFileName
	err = os.WriteFile(file, []byte(evaluation), 0666)
	if err != nil {
		return "", fmt.Errorf("error writing verdict to file: %w", err)
	}

	return verdict, nil
}

func GiveScores(compilation, requirements, apiKey string) (int, error) {
	evaluatorSystemPrompt := fmt.Sprintf(cases.BaseScoreEvaluator, requirements)
	evaluatorMsgs := gpt.ReplaceSystemPrompt(gpt.Messages{}, evaluatorSystemPrompt)
	evaluatorMsgs = gpt.AddUserMessage(evaluatorMsgs, compilation)

	gptModel := gpt.NewGPT(gpt.GPT4Turbo, apiKey, 0.6, "evaluator")

	response, err := gptModel.SendChat(evaluatorMsgs)
	if err != nil {
		return 0, fmt.Errorf("failed to get score from gpt: %w", err)
	}

	scoreString := interaction.ExtractMarkdown(response, "score").Block
	score, err := strconv.Atoi(scoreString)
	if err != nil {
		return 0, fmt.Errorf("failed to convert gpt final score %s from string to int: %w", scoreString, err)
	}

	explanations := interaction.ExtractMarkdown(response, "explanations").Block
	if err != nil {
		return 0, fmt.Errorf("failed to convert gpt final score %s from string to int: %w", scoreString, err)
	}

	evaluation := fmt.Sprintf(
		"Score: %d\n"+
			"Explanations: \n%s",
		score, explanations,
	)

	file := config.RunFolderPath + scoresFileName
	err = os.WriteFile(file, []byte(evaluation), 0666)
	if err != nil {
		return 0, fmt.Errorf("error writing scores to file: %w", err)
	}

	return score, nil
}
