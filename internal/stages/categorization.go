package stages

import (
	"fmt"
	"github.com/chattronics/chattronics/internal/gpt"
	"github.com/chattronics/chattronics/internal/interaction"
	"github.com/chattronics/chattronics/internal/logging"
	"github.com/chattronics/chattronics/internal/prompts"
	"strings"
)

const categorySeparator = "="

var Categories = []string{
	"A", // Sensor
	"B", // Signal Conditioning
	"C", // Amplification
	"D", // Filtering
	"E", // Other Conditioning
	"F", // Direct Measurement
	"G", // Analog-to-Digital Conversion
	"H", // Digital Processing
	"I", // Other
}

func Categorize(m *gpt.GPT, msgs gpt.Messages) (map[string]string, error) {
	msgs = gpt.AddUserMessage(msgs, prompts.Categorization)
	response, err := m.SendChat(msgs)
	if err != nil {
		return nil, fmt.Errorf("failed to send categories message")
	}
	categoryLines := interaction.ExtractMarkdown(response, "categories").Lines

	categories, err := parseCategoriesIntoMap(categoryLines)
	if err != nil {
		return nil, fmt.Errorf("failed to parse categories")
	}

	return categories, nil
}

func parseCategoriesIntoMap(categoryLines []string) (map[string]string, error) {
	categoriesMap := make(map[string]string)
	for _, line := range categoryLines {
		parts := strings.Split(line, categorySeparator)
		if len(parts) > 2 {
			return nil, fmt.Errorf("was expecting to split line into 2 parts, but got %d", len(parts))
		}

		blockName := parts[0]
		blockCategory := parts[1]

		if !contains(Categories, blockCategory) {
			return nil, fmt.Errorf("received a invalid category: %s", blockCategory)
		}

		categoriesMap[blockName] = blockCategory
	}

	logging.Debug("parsed categories", logging.AddField("categories", categoriesMap))

	return categoriesMap, nil
}
