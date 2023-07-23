package main

import (
	"chattronics/pkg/chat"
	"chattronics/pkg/gpt"
	"chattronics/pkg/logging"
	"flag"
	"github.com/joho/godotenv"
	"os"
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
	err := godotenv.Load()
	if err != nil {
		logging.Fatal("Error loading .env file")
	}

	var engine gpt.Engine
	if mock {
		engine = gpt.NewMockGPT(model)
	} else {
		apikey := os.Getenv("OPENAI_API_KEY")
		if apikey == "" {
			logging.Fatal("Failed to load apikey from env var")
		}
		engine = gpt.NewGPTEngine(model, apikey, float32(temperature))
	}

	err = chat.StartConversation(&engine)
	if err != nil {
		logging.Fatal("Failed to start Conversation", logging.AddField("error", err))
		return
	}

}
