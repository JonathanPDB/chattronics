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

// Potentiometer
//
//go:embed potentiometer/available-info
var availableInfoPotentiometer string

//go:embed potentiometer/informative-description
var informativeDescPotentiometer string

//go:embed potentiometer/limited-description
var limitedDescPotentiometer string

//go:embed potentiometer/requirements
var requirementsPotentiometer string

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
	thermometry   = "thermometry"
	potentiometer = "potentiometer"
)

func GetInformative(testcase string) (string, string, error) {
	switch strings.ToLower(testcase) {
	case thermometry:
		informativeSysPrompt := BaseInformative + availableInfoThermometry
		return informativeDescThermometry, informativeSysPrompt, nil
	case potentiometer:
		informativeSysPrompt := BaseInformative + availableInfoPotentiometer
		return informativeDescPotentiometer, informativeSysPrompt, nil
	default:
		return "", "", fmt.Errorf("did not find test case for informative prompts")
	}
}

func GetLimited(testcase string) (string, error) {
	switch strings.ToLower(testcase) {
	case thermometry:
		return limitedDescThermometry, nil
	case potentiometer:
		return limitedDescPotentiometer, nil
	default:
		return "", fmt.Errorf("did not find test case for limited prompts")
	}
}

func GetRequirements(testcase string) (string, error) {
	switch strings.ToLower(testcase) {
	case thermometry:
		return requirementsThermometry, nil
	case potentiometer:
		return requirementsPotentiometer, nil
	default:
		return "", fmt.Errorf("did not find test case for requirements")
	}

}
