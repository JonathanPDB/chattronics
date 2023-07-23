package chat

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var stdin io.Reader

func init() {
	stdin = os.Stdin
}

type MarkdownBlock struct {
	Language string   `json:"language"`
	Lines    []string `json:"lines"`
	Block    string   `json:"block"`
}

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

	return output
}

func ExtractMarkdownBlocks(message string) []MarkdownBlock {
	var codeBlocks []MarkdownBlock

	splitMessages := strings.Split(message, "'''")
	for i, splitMessage := range splitMessages {
		if i%2 != 0 {
			lines := strings.Split(splitMessage, "\n")

			codeBlocks = append(codeBlocks, MarkdownBlock{
				Language: lines[0],
				Lines:    lines[1:],
				Block:    strings.Join(lines[1:], "\n"),
			})
		}
	}

	return codeBlocks
}
