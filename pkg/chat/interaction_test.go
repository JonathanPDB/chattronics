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
