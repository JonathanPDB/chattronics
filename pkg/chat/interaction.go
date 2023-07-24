package chat

import (
	"bufio"
	"fmt"
	"github.com/JonathanPDB/chattronics/pkg/logging"
	"github.com/TwiN/go-color"
	"io"
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

func greetingsMessage() {
	fmt.Println(color.Colorize(appColour, "Hello!"))
	fmt.Println(color.Colorize(appColour, "Welcome to Chattronics, a helper tool to generate top-down solutions for electronics projects!"))
	fmt.Println(color.Colorize(appColour, "Provide a short and specific description of your project."))
	fmt.Println(color.Colorize(warningColour, "You can user newlines to write your text. To send your messages type return twice."))
	fmt.Printf(color.Colorize(auxColour, "Description: "))
}

func readConsole() string {
	reader := bufio.NewReader(stdin)

askUser:
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
		goto askUser
	}

	return strings.TrimSuffix(output, "\n")
}

func askQuestions(questions []string) string {
	var answers []string
	fmt.Printf(color.Colorize(appColour, "GPT wants to ask some questions to make the problem clearer. Please answer one by one.\n"))

	for i, question := range questions {
		fmt.Println(color.Colorize(gptColour, question))
		fmt.Printf(color.Colorize(auxColour, "Answer: "))
		questionNumber := strconv.Itoa(i+1) + ". "
		answers = append(answers, questionNumber+readConsole())
	}

	return strings.Join(answers, "\n")
}
