package evaluate

import (
	"fmt"
	"github.com/chattronics/chattronics/internal/gpt"
	"github.com/chattronics/chattronics/internal/interaction"
	"github.com/chattronics/chattronics/tests/cases"
	"strconv"
)

func Evaluate(summary, requirements, apiKey string) (int, string, string, error) {
	evaluatorSystemPrompt := fmt.Sprintf(cases.BaseEvaluator, requirements)
	evaluatorMsgs := gpt.ReplaceSystemPrompt(gpt.Messages{}, evaluatorSystemPrompt)
	evaluatorMsgs = gpt.AddUserMessage(evaluatorMsgs, summary)

	gptModel := gpt.NewGPT(gpt.GPT4, apiKey, 0.2, "evaluator")

	response, err := gptModel.SendChat(evaluatorMsgs)
	if err != nil {
		return 0, "", "", fmt.Errorf("failed to get score from gpt: %w", err)
	}

	scoreString := interaction.ExtractMarkdown(response, "score").Block
	score, err := strconv.Atoi(scoreString)
	if err != nil {
		return 0, "", "", fmt.Errorf("failed to convert gpt final score %s from string to int: %w", scoreString, err)
	}

	verdict := interaction.ExtractMarkdown(response, "verdict").Block
	if err != nil {
		return 0, "", "", fmt.Errorf("failed to convert gpt final score %s from string to int: %w", scoreString, err)
	}

	explanations := interaction.ExtractMarkdown(response, "explanations").Block
	if err != nil {
		return 0, "", "", fmt.Errorf("failed to convert gpt final score %s from string to int: %w", scoreString, err)
	}

	return score, verdict, explanations, nil
}
