package main

import (
	"github.com/JonathanPDB/chattronics/pkg/chat"
	"github.com/JonathanPDB/chattronics/pkg/config"
	"github.com/JonathanPDB/chattronics/pkg/gpt"
	"github.com/JonathanPDB/chattronics/pkg/logging"
)

func main() {
	config.CreateFolders()
	logging.InitializeStandardLogger()

	engine := gpt.NewMockGPT(gpt.GPT3_5Turbo)

	err := chat.StartConversation(&engine)
	if err != nil {
		logging.Fatal("Failed to start Conversation", logging.AddField("error", err))
		return
	}
}
