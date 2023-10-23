package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
	"time"
)

const baseRunsPath = "/Users/I564244/Personal/TCC/new-chattronics/runs/"

var RunFolderPath string
var LogFolderPath string

func CreateFolders() {
	if _, err := os.Stat(baseRunsPath); os.IsNotExist(err) {
		err = os.Mkdir(baseRunsPath, os.ModePerm)
		if err != nil {
			log.Fatalf("Error to create runs folder: %s", err.Error())
		}
	}

	date := time.Now().Format("Jan02_15-04-05")
	RunFolderPath = baseRunsPath + date

	if strings.HasSuffix(os.Args[0], ".test") {
		RunFolderPath += "_TEST"
	}

	RunFolderPath += "/"

	err := os.Mkdir(RunFolderPath, os.ModePerm)
	if err != nil {
		log.Fatalf("Error to create current run folder: %s", err.Error())
	}

	LogFolderPath = RunFolderPath + "logs/"

	err = os.Mkdir(LogFolderPath, os.ModePerm)
	if err != nil {
		log.Fatalf("Error to create lgos folder: %s", err.Error())
	}
}

func LoadApiKeyEnvVar() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apikey := os.Getenv("OPENAI_API_KEY")
	if apikey == "" {
		log.Fatal("Failed to load apikey from env var")
	}

	return apikey
}
