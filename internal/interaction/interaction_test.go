package interaction

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"new-chattronics/internal/config"
	"testing"
)

func TestReadConsole(t *testing.T) {
	t.Run("Successfully Read input", func(t *testing.T) {
		providedInput := "This is a test input.\nIt has multiple lines.\nTo end press enter twice.\n\n"
		expectedInput := "This is a test input.\nIt has multiple lines.\nTo end press enter twice."

		var stdinTest bytes.Buffer

		stdinTest.Write([]byte(providedInput))
		stdin = &stdinTest

		actualInput := ReadConsole()

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

		actualInput := AskQuestions(providedQuestions)

		assert.Contains(t, actualInput, expectedInput)
	})
}

func TestIsUserSatisfied(t *testing.T) {
	config.CreateFolders(false)
	var stdinTest bytes.Buffer
	stdin = &stdinTest

	t.Run("Successfully read yes", func(t *testing.T) {
		stdinTest.Reset()

		providedInput := "yes\n"

		stdinTest.Write([]byte(providedInput))

		output := IsUserSatisfied()

		assert.True(t, output)
	})
	t.Run("Successfully read yes", func(t *testing.T) {
		stdinTest.Reset()

		providedInput := "yes\n"

		stdinTest.Write([]byte(providedInput))

		output := IsUserSatisfied()

		assert.True(t, output)
	})
	t.Run("Successfully read no", func(t *testing.T) {
		stdinTest.Reset()

		providedInput := "no\n"

		stdinTest.Write([]byte(providedInput))

		output := IsUserSatisfied()

		assert.False(t, output)
	})
	t.Run("Successfully read n", func(t *testing.T) {
		stdinTest.Reset()

		providedInput := "n\n"

		stdinTest.Write([]byte(providedInput))

		output := IsUserSatisfied()

		assert.False(t, output)
	})
	t.Run("Successfully read yes after invalid tries", func(t *testing.T) {
		stdinTest.Reset()

		providedInput := "invalid\nagain\nyes\n"

		stdinTest.Write([]byte(providedInput))

		output := IsUserSatisfied()

		assert.True(t, output)
	})

}
