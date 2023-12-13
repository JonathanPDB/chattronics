package testdetails

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/chattronics/chattronics/internal/gpt"
)

//go:embed thermometry-conversation.json
var conversationBytes []byte

func parseConversation(testCase string) (gpt.Messages, error) {
	var msgs gpt.Messages
	var fileBytes []byte

	switch testCase {
	case "thermometry":
		fileBytes = conversationBytes
	default:
		return nil, fmt.Errorf("test case not found when loading conversation file")
	}

	err := json.Unmarshal(fileBytes, &msgs)
	if err != nil {
		return nil, fmt.Errorf("failed to load messages from file: %w", err)
	}
	return msgs, nil
}
