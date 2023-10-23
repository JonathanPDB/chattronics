package logging

import (
	"encoding/json"
	"fmt"
	"log"
	"new-chattronics/internal/config"
	"os"
)

type ChatLogger struct {
	logFilePath string
}

type InteractionLog struct {
	Prompt          interface{} `json:"prompt"`
	ResponseMessage string      `json:"response"`
	Cost            float64     `json:"cost"`
	FinishReason    string      `json:"finish_reason"`
}

func NewChatLogger(persona string) *ChatLogger {
	if config.LogFolderPath == "" {
		log.Fatalf("LogFolderPath not set, can't initialize %s chat logger.", persona)
	}

	path := config.LogFolderPath + persona + ".json"
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %s", err.Error())
	}
	defer f.Close()

	return &ChatLogger{logFilePath: path}
}

func (l *ChatLogger) LogInteraction(incomingLog InteractionLog) error {
	content, err := os.ReadFile(l.logFilePath)
	if err != nil {
		return fmt.Errorf("failed to read log path: %w", err)
	}

	var logs []InteractionLog

	if len(content) != 0 {
		err = json.Unmarshal(content, &logs)
		if err != nil {
			return fmt.Errorf("failed to unmarshal interaction log content: %w", err)
		}
	}

	logs = append(logs, incomingLog)
	logsBytes, err := json.MarshalIndent(logs, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshall log bytes: %w", err)
	}

	err = os.WriteFile(l.logFilePath, logsBytes, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to write log file: %w", err)
	}

	return nil
}
