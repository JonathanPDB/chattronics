package interaction

import (
	"github.com/chattronics/chattronics/internal/config"
	"github.com/chattronics/chattronics/internal/logging"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMockUser_ReadConsole(t *testing.T) {
	config.CreateFolders("", false)
	logging.InitializeStandardLogger()

	t.Run("Successfully return project description only in the first iteration", func(t *testing.T) {
		user := CreateMockUser("Mock system prompt")

		resp := user.ReadConsole()
		assert.Equal(t, "Mock system prompt", resp)

		for i := 0; i < 20; i++ {
			resp = user.ReadConsole()
			assert.Equal(t, "no", resp)
		}
	})
}
