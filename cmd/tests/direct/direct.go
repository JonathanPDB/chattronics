package main

import (
	"flag"
	"github.com/chattronics/chattronics/internal"
	"github.com/chattronics/chattronics/internal/config"
	"github.com/chattronics/chattronics/internal/gpt"
	"github.com/chattronics/chattronics/internal/interaction"
	"github.com/chattronics/chattronics/internal/logging"
	"github.com/chattronics/chattronics/testbenches/cases"
	"github.com/chattronics/chattronics/testbenches/evaluate"
	"strconv"
)

const folderPrefix = "direct_"

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

	directScenarioPrompt, err := cases.GetDirect(testCase)
	if err != nil {
		logging.Fatal("Failed to get direct scenario information", logging.AddField("error", err))
		return
	}

	requirements, err := cases.GetRequirements(testCase)
	if err != nil {
		logging.Fatal("Failed to get requirements", logging.AddField("error", err))
		return
	}

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
		user := interaction.CreateMockUser(directScenarioPrompt)

		config.CreateFolders(folderName+"/"+strconv.Itoa(i), true)

		gptModel := gpt.NewGPT(model, apiKey, float32(temperature), "engineer")

		compilation, err := internal.RunApp(gptModel, user)
		if err != nil {
			logging.Error("Failed to run application", err)
			continue
		}

		verdict, err := evaluate.GiveVerdict(compilation, requirements, apiKey)
		if err != nil {
			logging.Error("Failed to get verdict", err)
			continue
		}

		score, err := evaluate.GiveScores(compilation, requirements, apiKey)
		if err != nil {
			logging.Error("Failed to get scores", err)
			continue
		}

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
	}

	logging.Info("Final results.",
		logging.AddField("score_average", scoresSum/iterations),
		logging.AddField("verdicts", verdicts),
	)
}
