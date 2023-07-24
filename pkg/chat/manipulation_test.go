package chat

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadConsole(t *testing.T) {
	t.Run("Successfully Read input", func(t *testing.T) {
		providedInput := "This is a test input.\nIt has multiple lines.\nTo end press enter twice.\n\n"
		expectedInput := "This is a test input.\nIt has multiple lines.\nTo end press enter twice."

		var stdinTest bytes.Buffer

		stdinTest.Write([]byte(providedInput))
		stdin = &stdinTest

		actualInput := readConsole()

		assert.Equal(t, expectedInput, actualInput)

	})
}

func TestMarkdownExtractCodeBlocks(t *testing.T) {
	t.Run("Successfully extract 1 markdown block", func(t *testing.T) {
		givenInput := "This is a Mock input.\nHere goes the code block:\n```Golang\npackage main\n\nfunc main() {\n\n}\n```\nDone. Code printed."
		expectedOutput := []MarkdownBlock{{
			Language: "Golang",
			Lines: []string{
				"package main",
				"",
				"func main() {",
				"",
				"}",
			},
			Block: "package main\n\nfunc main() {\n\n}",
		}}
		actualOutput := ExtractMarkdownBlocks(givenInput)

		assert.EqualValues(t, expectedOutput, actualOutput)
	})

	t.Run("Successfully extract 1 isolated markdown block", func(t *testing.T) {
		givenInput := "```Golang\npackage main\n\nfunc main() {\n\n}\n```"
		expectedOutput := []MarkdownBlock{{
			Language: "Golang",
			Lines: []string{
				"package main",
				"",
				"func main() {",
				"",
				"}",
			},
			Block: "package main\n\nfunc main() {\n\n}",
		}}
		actualOutput := ExtractMarkdownBlocks(givenInput)

		assert.EqualValues(t, expectedOutput, actualOutput)
	})

	t.Run("Successfully extract 2 markdown block", func(t *testing.T) {
		givenInput := "This is a Mock input.\nHere goes the code block:\n```Golang\npackage main\n\nfunc main() {\n\n}\n```\nNext block:\n```Python\nprint(\"Hello World\")\n```\nFinished"
		expectedOutput := []MarkdownBlock{{
			Language: "Golang",
			Lines: []string{
				"package main",
				"",
				"func main() {",
				"",
				"}",
			},
			Block: "package main\n\nfunc main() {\n\n}",
		}, {
			Language: "Python",
			Lines: []string{
				"print(\"Hello World\")",
			},
			Block: "print(\"Hello World\")",
		},
		}

		actualOutput := ExtractMarkdownBlocks(givenInput)

		assert.EqualValues(t, expectedOutput, actualOutput)
	})
}

func TestExtractSingleBlock(t *testing.T) {
	t.Run("Successfully extract Golang Block from response with 1 block", func(t *testing.T) {
		givenInput := "This is a Mock input.\nHere goes the code block:\n```Golang\npackage main\n\nfunc main() {\n\n}\n```\nDone. Code printed."
		givenLanguage := "golang"
		expectedResponse := MarkdownBlock{
			Language: "Golang",
			Lines: []string{
				"package main",
				"",
				"func main() {",
				"",
				"}",
			},
			Block: "package main\n\nfunc main() {\n\n}",
		}

		actualResponse := ExtractSingleBlock(givenInput, givenLanguage)

		assert.EqualValues(t, expectedResponse, actualResponse)
	})

	t.Run("Successfully extract Python Block from response with 2 block", func(t *testing.T) {
		givenInput := "This is a Mock input.\nHere goes the code block:\n```Golang\npackage main\n\nfunc main() {\n\n}\n```\nNext block:\n```Python\nprint(\"Hello World\")\n```\nFinished"
		givenLanguage := "python"
		expectedResponse := MarkdownBlock{
			Language: "Python",
			Lines: []string{
				"print(\"Hello World\")",
			},
			Block: "print(\"Hello World\")",
		}

		actualResponse := ExtractSingleBlock(givenInput, givenLanguage)

		assert.EqualValues(t, expectedResponse, actualResponse)
	})
}

func TestAskQuestions(t *testing.T) {
	t.Run("Successfully ask 1 questions and get response", func(t *testing.T) {
		providedQuestions := []string{"1. Does your project have to be in strictly digital or analog format?"}

		providedInput := "Yes.\nStrictly Analog.\n\n"
		expectedInput := "1. Yes.\nStrictly Analog."

		var stdinTest bytes.Buffer

		stdinTest.Write([]byte(providedInput))
		stdin = &stdinTest

		actualInput := askQuestions(providedQuestions)

		assert.Contains(t, actualInput, expectedInput)
	})
}
