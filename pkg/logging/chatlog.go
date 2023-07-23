package logging

import (
	"encoding/json"
	"log"
	"os"
	"strings"
	"time"
)

const baseLogPath = "/Users/I564244/Personal/TCC/chattronics/logs/"

var logDatetimePath string

type ChatLogger struct {
	logFilePath string
}

type InteractionLog struct {
	Prompt           interface{} `json:"prompt"`
	ResponseMessage  string      `json:"response"`
	InputTokenCount  int         `json:"inputTokens"`
	OutputTokenCount int         `json:"outputToken"`
	Cost             float64     `json:"cost"`
	FinishReason     string      `json:"finish_reason"`
}

func init() {
	if _, err := os.Stat(baseLogPath); os.IsNotExist(err) {
		err = os.Mkdir(baseLogPath, os.ModePerm)
		if err != nil {
			log.Fatalf("Error to create logs folder: %s", err.Error())
		}
	}

	date := time.Now().Format("Jan02_15-04-05")
	logDatetimePath = baseLogPath + date

	if strings.HasSuffix(os.Args[0], ".test") {
		logDatetimePath += "_TEST"
	}

	logDatetimePath += "/"

	err := os.Mkdir(logDatetimePath, os.ModePerm)
	if err != nil {
		log.Fatalf("Error to create specific log folder: %s", err.Error())
	}
}

func NewChatLogger(persona string) *ChatLogger {
	path := logDatetimePath + persona + ".json"
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
		return err
	}

	var logs []InteractionLog

	if len(content) != 0 {
		err = json.Unmarshal(content, &logs)
		if err != nil {
			return err
		}
	}

	logs = append(logs, incomingLog)
	logsBytes, err := json.MarshalIndent(logs, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(l.logFilePath, logsBytes, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
