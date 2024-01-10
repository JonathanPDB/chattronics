package evaluate

import (
	"encoding/json"
	"fmt"
	"github.com/chattronics/chattronics/internal/config"
	"github.com/chattronics/chattronics/internal/gpt"
	"github.com/chattronics/chattronics/internal/interaction"
	"github.com/chattronics/chattronics/internal/logging"
	"github.com/chattronics/chattronics/testbenches/cases"
	"os"
	"strconv"
	"strings"
)

const evaluationFileName = "evaluation.md"
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

func GiveScores(compilationList []string, requirements, apiKey string) error {
	evaluatorSystemPrompt := fmt.Sprintf(cases.BaseScoreEvaluator, requirements)
	evaluatorMsgs := gpt.ReplaceSystemPrompt(gpt.Messages{}, evaluatorSystemPrompt)

	requirementsList := strings.Split(requirements, "\n")

	var checklist map[string]string
	for i := range requirementsList {
		checklist[string(i)] = ""
	}

	gptModel := gpt.NewGPT(gpt.GPT4Turbo, apiKey, 0.6, "score")

	var totalScore []map[string]string
	var blockScores []map[string]string
	for i, explanation := range compilationList {
		msgs := gpt.AddUserMessage(evaluatorMsgs, explanation)

		response, err := gptModel.SendChat(msgs)
		if err != nil {
			return fmt.Errorf("failed to get score from gpt: %w", err)
		}

		blockScores[i] = make(map[string]string)
		err = json.Unmarshal([]byte(response), &blockScores[i])
		if err != nil {
			return fmt.Errorf("failed to unmarshal json from score: %w", err)
		}

		logging.Debug(response)
	}

	totalScore = blockScores

	for _, blockScore := range blockScores {
		for requirementNumber, overallScore := range checklist {
			if givenScore, ok := blockScore[requirementNumber]; ok {
				switch overallScore {
				case "":
					checklist[requirementNumber] = givenScore
				case "fulfilled":
					if givenScore == "failed" {
						checklist[requirementNumber] = "conflicting"
					}
				case "failed":
					if givenScore == "fulfilled" {
						checklist[requirementNumber] = "conflicting"
					}
				default:
					checklist[requirementNumber] = givenScore
				}
			} else {
				logging.Warn(fmt.Sprintf("did not get score for requirement: %s", requirementNumber))
			}
		}
	}

	totalScore = append(totalScore, checklist)

	totalScoreBytes, err := json.Marshal(totalScore)
	if err != nil {
		return fmt.Errorf("failed to marshal checklist: %w", err)
	}

	file := config.RunFolderPath + scoresFileName
	err = os.WriteFile(file, totalScoreBytes, 0666)
	if err != nil {
		return fmt.Errorf("error writing scores to file: %w", err)
	}

	return nil
}

func Evaluate(summary, requirements, apiKey string) (int, string, string, error) {
	evaluatorSystemPrompt := fmt.Sprintf(cases.BaseEvaluator, requirements)
	evaluatorMsgs := gpt.ReplaceSystemPrompt(gpt.Messages{}, evaluatorSystemPrompt)
	evaluatorMsgs = gpt.AddUserMessage(evaluatorMsgs, summary)

	gptModel := gpt.NewGPT(gpt.GPT4Turbo, apiKey, 0.2, "evaluator")

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

	evaluation := fmt.Sprintf(
		"Score: %d\n"+
			"Verdict: %s\n\n"+
			"Explanations: \n%s",
		score, verdict, explanations,
	)

	file := config.RunFolderPath + evaluationFileName
	err = os.WriteFile(file, []byte(evaluation), 0666)
	if err != nil {
		return 0, "", "", fmt.Errorf("error writing evaluation to file: %w", err)
	}

	return score, verdict, explanations, nil
}
