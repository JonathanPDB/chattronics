package cases

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed prompts/informative-sys
var BaseInvestigative string

//go:embed prompts/score-sys
var BaseScoreEvaluator string

//go:embed prompts/verdict-sys
var BaseVerdictEvaluator string

// Potentiometer
//
//go:embed potentiometer/available-info
var availableInfoPotentiometer string

//go:embed potentiometer/investigative-description
var investigativeDescPotentiometer string

//go:embed potentiometer/direct-description
var directDescPotentiometer string

//go:embed potentiometer/requirements
var requirementsPotentiometer string

// Thermometry
//
//go:embed thermometry/available-info
var availableInfoThermometry string

//go:embed thermometry/investigative-description
var investigativeDescThermometry string

//go:embed thermometry/direct-description
var directDescThermometry string

//go:embed thermometry/requirements
var requirementsThermometry string

// Accelerometry
//
//go:embed accelerometry/available-info
var availableInfoAccelerometry string

//go:embed accelerometry/investigative-description
var investigativeDescAccelerometry string

//go:embed accelerometry/direct-description
var directDescAccelerometry string

//go:embed accelerometry/requirements
var requirementsAccelerometry string

// Machine
//
//go:embed machine/available-info
var availableInfoMachine string

//go:embed machine/investigative-description
var investigativeDescMachine string

//go:embed machine/direct-description
var directDescMachine string

//go:embed machine/requirements
var requirementsMachine string

const (
	thermometry   = "thermometry"
	potentiometer = "potentiometer"
	accelerometry = "accelerometry"
	machine       = "machine"
)

func GetInvestigative(testcase string) (string, string, error) {
	switch strings.ToLower(testcase) {
	case thermometry:
		investigativeSysPrompt := BaseInvestigative + availableInfoThermometry
		return investigativeDescThermometry, investigativeSysPrompt, nil
	case potentiometer:
		investigativeSysPrompt := BaseInvestigative + availableInfoPotentiometer
		return investigativeDescPotentiometer, investigativeSysPrompt, nil
	case accelerometry:
		investigativeSysPrompt := BaseInvestigative + availableInfoAccelerometry
		return investigativeDescAccelerometry, investigativeSysPrompt, nil
	case machine:
		investigativeSysPrompt := BaseInvestigative + availableInfoMachine
		return investigativeDescMachine, investigativeSysPrompt, nil
	default:
		return "", "", fmt.Errorf("did not find test case for investigative prompts")
	}
}

func GetDirect(testcase string) (string, error) {
	switch strings.ToLower(testcase) {
	case thermometry:
		return directDescThermometry, nil
	case potentiometer:
		return directDescPotentiometer, nil
	case accelerometry:
		return directDescAccelerometry, nil
	case machine:
		return directDescMachine, nil
	default:
		return "", fmt.Errorf("did not find test case for direct prompts")
	}
}

func GetRequirements(testcase string) (string, error) {
	switch strings.ToLower(testcase) {
	case thermometry:
		return requirementsThermometry, nil
	case potentiometer:
		return requirementsPotentiometer, nil
	case accelerometry:
		return requirementsAccelerometry, nil
	case machine:
		return requirementsMachine, nil
	default:
		return "", fmt.Errorf("did not find test case for requirements")
	}

}
