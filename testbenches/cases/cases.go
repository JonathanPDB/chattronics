package cases

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed prompts/informative-sys
var BaseInformative string

//go:embed prompts/evaluate-sys
var BaseEvaluator string

// Thermometry
//
//go:embed thermometry/available-info
var availableInfoThermometry string

//go:embed thermometry/informative-description
var informativeDescThermometry string

//go:embed thermometry/limited-description
var limitedDescThermometry string

//go:embed thermometry/requirements
var requirementsThermometry string

const (
	themometry = "thermometry"
)

func GetInformative(testcase string) (string, string, error) {
	switch strings.ToLower(testcase) {
	case themometry:
		informativeSysPrompt := BaseInformative + availableInfoThermometry
		return informativeDescThermometry, informativeSysPrompt, nil
	default:
		return "", "", fmt.Errorf("did not find test case for informative prompts")
	}
}

func GetLimited(testcase string) (string, error) {
	switch strings.ToLower(testcase) {
	case themometry:
		return limitedDescThermometry, nil
	default:
		return "", fmt.Errorf("did not find test case for limited prompts")
	}
}

func GetRequirements(testcase string) (string, error) {
	switch strings.ToLower(testcase) {
	case themometry:
		return requirementsThermometry, nil
	default:
		return "", fmt.Errorf("did not find test case for requirements")
	}

}
