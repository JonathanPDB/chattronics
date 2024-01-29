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

const folderPrefix = "investigative_"

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
	flag.IntVar(&iterations, "iterations", 1, "How many iterations of the test case to run")
}

func main() {
	flag.Parse()
	folderName := config.CreateFolders(folderPrefix, false)
	logging.InitializeStandardLogger()
	apiKey := config.LoadApiKeyEnvVar()

	projectDesc, informativeSysPrompt, err := cases.GetInvestigative(testCase)
	if err != nil {
		logging.Fatal("Failed to get investigative scenario information", logging.AddField("error", err))
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
		user := interaction.CreateEmulatorUser(projectDesc, informativeSysPrompt, apiKey)

		config.CreateFolders(folderName+"/"+strconv.Itoa(i), true) //TODO: Fix how the log folder is working

		gptModel := gpt.NewGPT(model, apiKey, float32(temperature), "engineer")

		compilation, err := internal.RunApp(gptModel, user)
		if err != nil {
			logging.Fatal("Failed to run application", logging.AddField("error", err))
			return
		}

		verdict, err := evaluate.GiveVerdict(compilation, requirements, apiKey)
		if err != nil {
			logging.Fatal("Failed to get verdict", logging.AddField("error", err))
			return
		}

		score, err := evaluate.GiveScores(compilation, requirements, apiKey)
		if err != nil {
			logging.Fatal("Failed to get scores", logging.AddField("error", err))
			return
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
