package prompts

import (
	_ "embed"
)

//go:embed reviewer/system
var ReviewerSystem string

//go:embed reviewer/formulate-prefix
var FormulatePrefix string

//go:embed reviewer/formulate-suffix
var FormulateSuffix string

//go:embed architect/system
var ArchitectSystem string

//go:embed mock-responses/questions
var MockQuestionsResponse string

//go:embed mock-responses/ideal-prompt
var MockIdealPromptResponse string

//go:embed mock-responses/dot-diagram
var MockDiagramResponse string
