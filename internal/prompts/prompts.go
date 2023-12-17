package prompts

import _ "embed"

//go:embed architecture/questions-sys
var ArchitectureQuestionsSystem string

//go:embed architecture/design-sys
var ArchitectureDesignSystem string

//go:embed architecture/feedback
var ArchitectureFeedback string

//go:embed categorization/categorization
var Categorization string

//go:embed details/questions-sys
var DetailsQuestionsSystem string

//go:embed details/feedback
var DetailsFeedback string

//go:embed summary/summary
var Summary string

//go:embed details/A-Sensor/questions
var detailsQuestions_A string

//go:embed details/A-Sensor/details-sys
var detailsSystem_A string

//go:embed details/B-Signal_conditioning/questions
var detailsQuestions_B string

//go:embed details/B-Signal_conditioning/details-sys
var detailsSystem_B string

//go:embed details/C-Amplification/questions
var detailsQuestions_C string

//go:embed details/C-Amplification/details-sys
var detailsSystem_C string

//go:embed details/D-Filtering/questions
var detailsQuestions_D string

//go:embed details/D-Filtering/details-sys
var detailsSystem_D string

//go:embed details/E-Other_conditioning/questions
var detailsQuestions_E string

//go:embed details/E-Other_conditioning/details-sys
var detailsSystem_E string

//go:embed details/F-Direct_measurement/questions
var detailsQuestions_F string

//go:embed details/F-Direct_measurement/details-sys
var detailsSystem_F string

//go:embed details/G-ADC/questions
var detailsQuestions_G string

//go:embed details/G-ADC/details-sys
var detailsSystem_G string

//go:embed details/H-Digital_processing/questions
var detailsQuestions_H string

//go:embed details/H-Digital_processing/details-sys
var detailsSystem_H string

//go:embed details/I-Other/questions
var detailsQuestions_I string

//go:embed details/I-Other/details-sys
var detailsSystem_I string

var categoriesDetailsMap = map[string]string{
	"A": detailsSystem_A,
	"B": detailsSystem_B,
	"C": detailsSystem_C,
	"D": detailsSystem_D,
	"E": detailsSystem_E,
	"F": detailsSystem_F,
	"G": detailsSystem_G,
	"H": detailsSystem_H,
	"I": detailsSystem_I,
}

var categoriesDetailsQuestionsMap = map[string]string{
	"A": detailsQuestions_A,
	"B": detailsQuestions_B,
	"C": detailsQuestions_C,
	"D": detailsQuestions_D,
	"E": detailsQuestions_E,
	"F": detailsQuestions_F,
	"G": detailsQuestions_G,
	"H": detailsQuestions_H,
	"I": detailsQuestions_I,
}

func GetDetailsSystemPrompt(category string) string {
	return categoriesDetailsMap[category]
}

func GetDetailsQuestionPrompt(category string) string {
	return categoriesDetailsQuestionsMap[category]
}
