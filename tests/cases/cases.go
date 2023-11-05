package cases

import (
	_ "embed"
	"fmt"
	"os"
)

//go:embed prompts/informative-sys
var BaseInformative string

//go:embed prompts/evaluate-sys
var BaseEvaluator string

func GetInformative(testcase string) (string, string, error) {
	informativeProjectDescription, err := loadFile("tests/cases/" + testcase + "/informative-description")
	if err != nil {
		return "", "", fmt.Errorf("failed to get informative project description for testcase: %w", err)
	}
	availableInfo, err := loadFile("tests/cases/" + testcase + "/available-info")
	if err != nil {
		return "", "", fmt.Errorf("failed to get available info for testcase: %w", err)
	}

	informativeSysPrompt := BaseInformative + availableInfo
	return informativeProjectDescription, informativeSysPrompt, nil

}

func GetLimited(testcase string) (string, error) {
	limitedProjectDescription, err := loadFile("tests/cases/" + testcase + "/limited-description")
	if err != nil {
		return "", fmt.Errorf("failed to get limited project description for testcase: %w", err)
	}
	return limitedProjectDescription, nil
}

func GetRequirements(testcase string) (string, error) {
	requirements, err := loadFile("tests/cases/" + testcase + "/requirements")
	if err != nil {
		return "", fmt.Errorf("failed to get requirements for testcase: %w", err)
	}
	return requirements, nil

}

func loadFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("failed to open %s file: %w", filename, err)
	}
	return string(content), nil
}
