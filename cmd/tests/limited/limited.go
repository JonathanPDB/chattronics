package main

import (
	"flag"
	"new-chattronics/internal"
	"new-chattronics/internal/config"
	"new-chattronics/internal/gpt"
	"new-chattronics/internal/interaction"
	"new-chattronics/internal/logging"
	"new-chattronics/tests/cases"
	"new-chattronics/tests/evaluate"
	"strconv"
)

const folderPrefix = "limited_"

var (
	model       string
	temperature float64
	testCase    string
	iterations  int
)

func init() {
	flag.StringVar(&model, "model", "gpt-3.5-turbo", "GPT model version")
	flag.Float64Var(&temperature, "temperature", 0.8, "GPT temperature. Higher values, higher variability")
	flag.StringVar(&testCase, "testCase", "", "test case to be executed")
	flag.IntVar(&iterations, "iterations", 1, "How many iterations of the testcase to run")
}

func main() {
	flag.Parse()
	folderName := config.CreateFolders(folderPrefix, false)
	logging.InitializeStandardLogger()
	apiKey := config.LoadApiKeyEnvVar()

	limitedScenarioPrompt, err := cases.GetLimited(testCase)
	if err != nil {
		logging.Fatal("Failed to get limited scenario information", logging.AddField("error", err))
		return
	}

	requirements, err := cases.GetRequirements(testCase)
	if err != nil {
		logging.Fatal("Failed to get requirements", logging.AddField("error", err))
		return
	}

	user := interaction.CreateMockUser(limitedScenarioPrompt)

	verdicts := map[string]int{
		"generic":    0,
		"acceptable": 0,
		"optimal":    0,
		"unfeasible": 0,
		"incorrect":  0,
		"INVALID":    0,
	}

	scoresSum := 0

	for i := 0; i < iterations; i++ {
		config.CreateFolders(folderName+"/"+strconv.Itoa(i), true)

		gptModel := gpt.NewGPT(model, apiKey, float32(temperature), "engineer")

		summary, err := internal.RunApp(gptModel, user)
		if err != nil {
			logging.Fatal("Failed to run application", logging.AddField("error", err))
			return
		}

		score, verdict, explanations, err := evaluate.Evaluate(summary, requirements, apiKey)

		if _, ok := verdicts[verdict]; ok {
			verdicts[verdict]++
		} else {
			verdicts["INVALID"]++
		}

		scoresSum += score

		logging.Info("Finished iteration.",
			logging.AddField("score", score),
			logging.AddField("verdict", verdict),
		)
		logging.Info("Explanations.",
			logging.AddField("explanations", explanations),
		)
	}

	logging.Info("Final results.",
		logging.AddField("score_average", scoresSum/iterations),
		logging.AddField("verdicts", verdicts),
	)
}
