package prompts

import (
	_ "embed"
)

//go:embed inquisitor/system
var InquistorSystem string

//go:embed inquisitor/formulate-prefix
var FormulatePrefix string

//go:embed inquisitor/formulate-suffix
var FormulateSuffix string

//go:embed top-down/system
var TopDownSystem string

//go:embed test
var Test string
