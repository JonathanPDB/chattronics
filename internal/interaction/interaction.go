package interaction

import (
	"bufio"
	"fmt"
	"github.com/TwiN/go-color"
	"io"
	"new-chattronics/internal/logging"
	"os"
	"strconv"
	"strings"
)

var (
	appColour     = color.Cyan
	gptColour     = color.Green
	auxColour     = color.Gray
	errorColour   = color.Red
	warningColour = color.Yellow
)

var stdin io.Reader

func init() {
	stdin = os.Stdin
}

func GreetingsMessage() {
	fmt.Println(color.Colorize(appColour, "Hello!"))
	fmt.Println(color.Colorize(appColour, "Welcome to Chattronics, a helper tool to generate top-down solutions for electronics projects!"))
	fmt.Println(color.Colorize(appColour, "Provide a short and specific description of your project."))
	fmt.Println(color.Colorize(warningColour, "You can user newlines to write your text. To send your messages type return twice."))
}

func ReadConsole() string {
	reader := bufio.NewReader(stdin)

readUserInput:
	var lines []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			logging.Fatal(fmt.Sprintf("failed to read user input: %s", err.Error()))
		}
		if len(strings.TrimSpace(line)) == 0 {
			break
		}
		lines = append(lines, line)
	}
	output := strings.Join(lines, "")

	if output == "\n" || output == "" {
		fmt.Println(color.Colorize(errorColour, "Please type valid text!"))
		goto readUserInput
	}

	return strings.TrimSuffix(output, "\n")
}

func AskQuestions(questions []string) string {
	var answers []string
	for i, question := range questions {
		fmt.Println(color.Colorize(gptColour, question))
		fmt.Printf(color.Colorize(auxColour, "Answer: "))
		questionNumber := strconv.Itoa(i+1) + ". "
		answers = append(answers, questionNumber+ReadConsole())
	}
	return strings.Join(answers, "\n")
}

//func GetAdditionalComment() string {
//	comment := ReadConsole()
//	if strings.ToLower(comment) == "no" || strings.ToLower(comment) == "n" {
//		return ""
//	}
//
//	return fmt.Sprintf("Additional comment: %s", comment)
//}

func PrintExplanations(explanations []string) {
	for _, explanation := range explanations {
		fmt.Println(color.Colorize(gptColour, explanation) + "\n")
	}
}

func IsUserSatisfied() bool {
	reader := bufio.NewReader(stdin)

	fmt.Println(color.Colorize(appColour, "Are you satisfied with this answer? [y/n]"))
	defer fmt.Println()

readUserInput:
	isSatisfied, err := reader.ReadString('\n')
	isSatisfied = strings.TrimSuffix(isSatisfied, "\n")
	if err != nil {
		logging.Warn("Failed to get user satisfaction answer.")
		return false
	}
	switch strings.ToLower(isSatisfied) {
	case "y":
		return true
	case "yes":
		return true
	case "n":
		return false
	case "no":
		return false
	default:
		fmt.Println(color.Colorize(warningColour, "Please insert a valid answer!\n"))
		goto readUserInput
	}
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

func getChangeRequest() string {
	fmt.Println(color.Colorize(appColour, "Tell GPT which parts it should change and why (if possible)."))
	return ReadConsole()
}

func requiremetsPrompt() {
	fmt.Println(color.Colorize(appColour, "Now GPT will provide projects specifications to each block and for that, they might ask some questions regarding requirements,"))
}
