package chat

import (
	"bufio"
	"fmt"
	"github.com/TwiN/go-color"
	"log"
	"strconv"
	"strings"
)

func readConsole() string {
	reader := bufio.NewReader(stdin)

askUser:
	var lines []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if len(strings.TrimSpace(line)) == 0 {
			break
		}
		lines = append(lines, line)
	}
	output := strings.Join(lines, "")

	if output == "\n" || output == "" {
		fmt.Println("Please type valid text!")
		goto askUser
	}

	return strings.TrimSuffix(output, "\n")
}

func askQuestions(questions []string) string {
	var answers []string
	fmt.Println(color.Colorize(appColour, "GPT wants to ask some questions to make the problem clearer. Please answer one by one."))
	fmt.Println("")
	for i, question := range questions {
		fmt.Println(color.Colorize(gptColour, question))
		fmt.Printf(color.Colorize(auxColour, "Answer: "))
		questionNumber := strconv.Itoa(i+1) + ". "
		answers = append(answers, questionNumber+readConsole())
	}

	return strings.Join(answers, "\n")
}
