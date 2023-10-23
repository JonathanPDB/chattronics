package main

import (
	"flag"
	"new-chattronics/internal"
	"new-chattronics/internal/config"
	"new-chattronics/internal/gpt"
	"new-chattronics/internal/logging"
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
	config.CreateFolders()
	apiKey := config.LoadApiKeyEnvVar()
	logging.InitializeStandardLogger()

	gptModel := gpt.NewGPT(model, apiKey, float32(temperature))

	err := internal.RunApp(gptModel)
	if err != nil {
		logging.Fatal("Failed to run application", logging.AddField("error", err))
		return
	}
}
