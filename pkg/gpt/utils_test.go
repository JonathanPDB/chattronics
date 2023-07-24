package gpt

import (
	"fmt"
	"github.com/JonathanPDB/chattronics/pkg/logging"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetBRLfromUSD(t *testing.T) {
	t.Run("Should convert USD to BRL", func(t *testing.T) {
		givenUSD := 0.03
		expectedBRL := 0.03 * 4.6

		actualBRL := getBRLfromUSD(givenUSD)
		logging.Debug(fmt.Sprintf("value in BRL: %f", actualBRL))

		assert.GreaterOrEqual(t, actualBRL, expectedBRL)
	})
}
