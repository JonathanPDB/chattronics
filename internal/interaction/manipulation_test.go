package interaction

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
		actualOutput := extractMarkdownBlocks(givenInput)

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
		actualOutput := extractMarkdownBlocks(givenInput)

		assert.EqualValues(t, expectedOutput, actualOutput)
	})

	t.Run("Successfully extract 2 markdown block", func(t *testing.T) {
		givenInput := "This is a Mock input.\nHere goes the code block:\n```Golang\npackage main\n\nfunc main() {\n\n}\n```\n```Python\nprint(\"Hello World\")\n```\nFinished"
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

		actualOutput := extractMarkdownBlocks(givenInput)

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
