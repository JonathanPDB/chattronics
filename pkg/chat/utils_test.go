package chat

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadConsole(t *testing.T) {
	t.Run("Successfully Read input", func(t *testing.T) {
		providedInput := "This is a test input.\nIt has multiple lines.\nTo end press enter twice.\n\n"
		expectedInput := "This is a test input.\nIt has multiple lines.\nTo end press enter twice.\n"

		var stdinTest bytes.Buffer

		stdinTest.Write([]byte(providedInput))
		stdin = &stdinTest

		actualInput := readConsole()

		assert.Equal(t, expectedInput, actualInput)

	})
}

func TestMarkdownExtractCodeBlocks(t *testing.T) {
	t.Run("Successfully extract 1 markdown block", func(t *testing.T) {
		givenInput := "This is a Mock input.\nHere goes the code block:\n'''Golang\npackage main\n\nfunc main() {\n\n}\n'''\nDone. Code printed."
		expectedOutput := []MarkdownBlock{{
			Language: "Golang",
			Lines: []string{
				"package main",
				"",
				"func main() {",
				"",
				"}",
				"",
			},
		}}
		actualOutput := ExtractMarkdownBlocks(givenInput)

		assert.EqualValues(t, expectedOutput, actualOutput)
	})

	t.Run("Successfully extract 1 isolated markdown block", func(t *testing.T) {
		givenInput := "'''Golang\npackage main\n\nfunc main() {\n\n}\n'''"
		expectedOutput := []MarkdownBlock{{
			Language: "Golang",
			Lines: []string{
				"package main",
				"",
				"func main() {",
				"",
				"}",
				"",
			},
		}}
		actualOutput := ExtractMarkdownBlocks(givenInput)

		assert.EqualValues(t, expectedOutput, actualOutput)
	})

	t.Run("Successfully extract 2 markdown block", func(t *testing.T) {
		givenInput := "This is a Mock input.\nHere goes the code block:\n'''Golang\npackage main\n\nfunc main() {\n\n}\n'''\nNext block:\n'''Python\nprint(\"Hello World\")\n'''\nFinished"
		expectedOutput := []MarkdownBlock{{
			Language: "Golang",
			Lines: []string{
				"package main",
				"",
				"func main() {",
				"",
				"}",
				"",
			},
		}, {
			Language: "Python",
			Lines: []string{
				"print(\"Hello World\")",
				"",
			},
		},
		}

		actualOutput := ExtractMarkdownBlocks(givenInput)

		assert.EqualValues(t, expectedOutput, actualOutput)
	})
}
