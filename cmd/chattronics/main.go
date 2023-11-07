package main

import (
	"flag"
	"github.com/chattronics/chattronics/internal"
	"github.com/chattronics/chattronics/internal/config"
	"github.com/chattronics/chattronics/internal/gpt"
	"github.com/chattronics/chattronics/internal/interaction"
	"github.com/chattronics/chattronics/internal/logging"
)

var (
	model       string
	temperature float64
)

func init() {
	flag.StringVar(&model, "model", "gpt-3.5-turbo", "GPT model version")
	flag.Float64Var(&temperature, "temperature", 0.8, "GPT temperature. Higher values, higher variability")
}

func main() {
	flag.Parse()
	config.CreateFolders("real_", false)
	apiKey := config.LoadApiKeyEnvVar()
	logging.InitializeStandardLogger()

	gptModel := gpt.NewGPT(model, apiKey, float32(temperature), "engineer")
	user := interaction.CreateHumanUser()

	_, err := internal.RunApp(gptModel, user)
	if err != nil {
		logging.Fatal("Failed to run application", logging.AddField("error", err))
		return
	}
}
