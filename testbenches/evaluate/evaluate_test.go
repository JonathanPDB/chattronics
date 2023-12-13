package evaluate

import (
	"fmt"
	"github.com/chattronics/chattronics/internal/config"
	"github.com/chattronics/chattronics/internal/logging"
	"github.com/stretchr/testify/assert"
	"testing"
)

func getRequirements() string {
	return "" +
		"1. Use an NTC\n" +
		"2. Linearize the NTC.\n" +
		"3. Use a resistor calculated via the 3 point method to linearize the NTC.\n" +
		"4. Use a Wheatstone bridge instead of measuring the resistance.\n" +
		"5. Amplify the signal before measuring.\n" +
		"6. Take into account the fact that the maximum current cannot cause a significant auto heating effect."
}

func TestEvaluate(t *testing.T) {
	config.CreateFolders("", false)
	apiKey := config.LoadApiKeyEnvVar()
	logging.InitializeStandardLogger()

	iterations := 2

	type testCase struct {
		givenSummary    string
		expectedScore   int
		expectedVerdict string
	}

	testTable := []testCase{
		{
			givenSummary: "NTC Linearization to measure the temperature of a beaker with water.\n" +
				"\n" +
				"\n" +
				"This project consists in measuring the temperature of the water inside a glass beaker.\n" +
				"For this, a NTC will be used in conjunction with series of conditioning stages, which will " +
				"enable the measurement of a output voltage by using a simple commercial multimeter.\n" +
				"The first stage consists of the NTC sensor, which will be used for temperatures between 15 and " +
				"85 degrees celsius. Instead of measuring the resistance, a Wheatstone bridge will be used, in which " +
				"the other 3 elements of the bridge will guarantee a differential voltage of zero for a temperature " +
				"of 15 degrees. After the bridge, a instrumentation amplifier will be used with a gain of 200, so that " +
				"the voltage can be easily measured by using a low quality multimeter. Given that there is no digital " +
				"processing, there is no need of using a low pass filter for anti-aliasing.",
			expectedScore:   3,
			expectedVerdict: "acceptable",
		},
	}

	for j, test := range testTable {
		fmt.Printf("\nStating Testcase %d\n", j)

		for i := 0; i < iterations; i++ {
			fmt.Printf("Iteration %d\n", i)
			score, verdict, explanations, err := Evaluate(test.givenSummary, getRequirements(), apiKey)

			assert.NoError(t, err)
			assert.NotEmpty(t, explanations)
			assert.NotEqual(t, test.expectedScore, score)
			assert.Equal(t, test.expectedVerdict, verdict)
		}
	}
}
