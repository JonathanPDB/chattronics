package interaction

import (
	"encoding/json"
	"fmt"
	"github.com/chattronics/chattronics/internal/logging"
	"reflect"
	"strings"
)

type MarkdownBlock struct {
	Language string   `json:"language"`
	Lines    []string `json:"lines"`
	Block    string   `json:"block"`
}

func extractAllMarkdownBlocks(message string) []MarkdownBlock {
	var codeBlocks []MarkdownBlock

	splitMessages := strings.Split(message, "```")
	for i, splitMessage := range splitMessages {
		if i%2 != 0 {
			lines := strings.Split(strings.TrimSuffix(splitMessage, "\n"), "\n")

			codeBlocks = append(codeBlocks, MarkdownBlock{
				Language: lines[0],
				Lines:    lines[1:],
				Block:    strings.Join(lines[1:], "\n"),
			})
		}
	}

	return codeBlocks
}

func ExtractMarkdown(message, language string) MarkdownBlock {
	blocks := extractAllMarkdownBlocks(message)
	var desiredBlock MarkdownBlock
	for _, block := range blocks {
		if strings.ToLower(block.Language) == strings.ToLower(language) {
			if reflect.DeepEqual(desiredBlock, MarkdownBlock{}) {
				desiredBlock = block
			} else {
				logging.Warn("message contains more than 1 of the desired block; returning the first one.")
				return desiredBlock
			}
		}
	}
	return desiredBlock
}

func ExtractJsonSlice(message string) (map[string][]string, error) {
	_, trimmedMessage, found := strings.Cut(message, "{")
	if !found {
		return nil, fmt.Errorf("did not find the character { inside the message")
	}

	trimmedMessage, _, found = strings.Cut(trimmedMessage, "}")
	if !found {
		return nil, fmt.Errorf("did not find the character } inside the message")
	}

	trimmedMessage = "{" + trimmedMessage + "}"

	var extractedJson map[string][]string
	err := json.Unmarshal([]byte(trimmedMessage), &extractedJson)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal trimmed message: %w", err)
	}

	return extractedJson, nil
}
