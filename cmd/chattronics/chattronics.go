package main

import (
	"flag"
	"github.com/JonathanPDB/chattronics/pkg/chat"
	"github.com/JonathanPDB/chattronics/pkg/config"
	"github.com/JonathanPDB/chattronics/pkg/gpt"
	"github.com/JonathanPDB/chattronics/pkg/logging"
)

var (
	model       string
	temperature float64
	mock        bool
)

func init() {
	flag.StringVar(&model, "model", gpt.GPT3_5Turbo, "GPT model version")
	flag.Float64Var(&temperature, "temperature", 0.1, "GPT temperature. Higher values, higher variability")
	flag.BoolVar(&mock, "mock", true, "Use mock gpt instead of real (priced) model")
}

func main() {
	config.CreateFolders()
	apiKey := config.LoadApiKeyEnvVar()
	logging.InitializeStandardLogger()

	engine := gpt.NewGPT(model, apiKey, float32(temperature))

	err := chat.StartConversation(&engine)
	if err != nil {
		logging.Fatal("Failed to start Conversation", logging.AddField("error", err))
		return
	}
}
