package stages

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseCategories(t *testing.T) {
	t.Run("Should parse categories correctly", func(t *testing.T) {
		givenCategories := []string{
			"NTC Thermistor=A",
			"Linearization Circuit=E",
			"Temperature-Voltage Conversion Circuit=B",
			"Scaling and Shift Circuit=C",
			"Low-Pass Filter=D",
			"Output Voltage=F",
		}

		expectedCategoriesMap := map[string]string{
			"NTC Thermistor":                         "A",
			"Linearization Circuit":                  "E",
			"Temperature-Voltage Conversion Circuit": "B",
			"Scaling and Shift Circuit":              "C",
			"Low-Pass Filter":                        "D",
			"Output Voltage":                         "F",
		}

		actualCategoriesMap, err := parseCategoriesIntoMap(givenCategories)

		assert.NoError(t, err)
		assert.Equal(t, expectedCategoriesMap, actualCategoriesMap)
	})
}
