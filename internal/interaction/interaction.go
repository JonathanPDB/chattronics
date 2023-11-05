package interaction

import (
	"fmt"
	"github.com/TwiN/go-color"
)

var (
	appColour     = color.Cyan
	gptColour     = color.Green
	auxColour     = color.Gray
	errorColour   = color.Red
	warningColour = color.Yellow
)

type User interface {
	ReadConsole() string
	AskQuestions([]string) string
	IsUserSatisfied() bool
}

func GreetingsMessage() {
	fmt.Println(color.Colorize(appColour, "Hello!"))
	fmt.Println(color.Colorize(appColour, "Welcome to Chattronics, a helper tool to generate top-down solutions for electronics projects!"))
	fmt.Println(color.Colorize(appColour, "Provide a short and specific description of your project."))
	fmt.Println(color.Colorize(warningColour, "You can user newlines to write your text. To send your messages type return twice."))
}

func PrintAppMessage(message string) {
	fmt.Print(color.Colorize(appColour, message))
}

func PrintAuxMessage(message string) {
	fmt.Print(color.Colorize(auxColour, message))
}

func PrintGPTMessage(message string) {
	fmt.Print(color.Colorize(gptColour, message))
}
