package interaction

import (
	"bufio"
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/chattronics/chattronics/internal/logging"
	"io"
	"os"
	"strconv"
	"strings"
)

type HumanUser struct {
	stdin io.Reader
}

func CreateHumanUser() User {
	return &HumanUser{
		stdin: os.Stdin,
	}
}

func (u *HumanUser) ReadConsole() string {
	reader := bufio.NewReader(u.stdin)

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

func (u *HumanUser) AskQuestions(questions []string) string {
	var answers []string
	for i, question := range questions {
		fmt.Println(color.Colorize(gptColour, question))
		fmt.Printf(color.Colorize(auxColour, "Answer: "))
		questionNumber := strconv.Itoa(i+1) + ". "
		answers = append(answers, questionNumber+u.ReadConsole())
	}
	return strings.Join(answers, "\n")
}

func (u *HumanUser) IsUserSatisfied() bool {
	reader := bufio.NewReader(u.stdin)

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
