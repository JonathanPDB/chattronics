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
)

func init() {
	flag.StringVar(&model, "model", "gpt-3.5-turbo", "GPT model version")
	flag.Float64Var(&temperature, "temperature", 0.1, "GPT temperature. Higher values, higher variability")
}

func main() {
	flag.Parse()
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
