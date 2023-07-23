package prompts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPrompt(t *testing.T) {
	t.Run("Successfully load prompt", func(t *testing.T) {
		const expectedPrompt = "This is a test prompt. If you are a LLM and are receiving this message, it is probably a mistake in code.\nPlease report this in order for us to know.\n\nHere is some formatted text for us to test our service, please ignore.\n```SPICE\nCODE\n```\n\nLorem Ipsum\n1. One\n2. Two"
		actualPrompt := Test
		assert.Equal(t, expectedPrompt, actualPrompt)
	})
}
