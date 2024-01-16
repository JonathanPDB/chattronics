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
		"2. Linearize the NTC analogically.\n" +
		"3. Use a resistor calculated via the 3 point method to linearize the NTC.\n" +
		"4. The order of the filtering is at least 3\n" +
		"5. There is a method to control the current that goes through the NTC." +
		"6. The value of the gain should be somewhat close to -2.15"
}

func TestGiveScores(t *testing.T) {
	config.CreateFolders("", false)
	apiKey := config.LoadApiKeyEnvVar()
	logging.InitializeStandardLogger()

	iterations := 1

	type testCase struct {
		givenSummary  string
		expectedScore int
	}

	testTable := []testCase{
		{
			givenSummary: "NTC Linearization to measure the temperature of a beaker with water.\n" +
				"\n" +
				"\n" +
				"This project consists in measuring the temperature of the water inside a glass beaker.\n" +
				"For this, a NTC Vishay BC123 will be used in conjunction with series of conditioning stages, which will " +
				"enable the measurement of a output voltage by using a simple commercial multimeter.\n" +
				"The first stage consists of the NTC sensor, which will be used for temperatures between 15 and " +
				"85 degrees celsius. Instead of measuring the resistance, a Wheatstone bridge will be used, in which " +
				"the other 3 elements of the bridge will guarantee a differential voltage of zero for a temperature " +
				"of 15 degrees. After the bridge, a inverting amplifier in which the resistor values are R_f=4k, R_in=2k, so that " +
				"the voltage can be easily measured by using a low quality multimeter. The final stage before the multimeter is" +
				"a sallen-key low pass filter with a cutoff frequency of 10 Hz, in such a way that 100 Hz the damping is at -40 dB",
			expectedScore: 3,
		},
	}

	for j, test := range testTable {
		fmt.Printf("\nStating Testcase %d\n", j)

		for i := 0; i < iterations; i++ {
			fmt.Printf("Iteration %d\n", i)

			score, err := GiveScores(test.givenSummary, getRequirements(), apiKey)

			assert.NoError(t, err)
			assert.Equal(t, test.expectedScore, score)
		}
	}
}
