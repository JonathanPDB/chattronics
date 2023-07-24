package chat

import (
	"fmt"
	"github.com/JonathanPDB/chattronics/pkg/gpt"
	"github.com/JonathanPDB/chattronics/pkg/graph"
	"github.com/JonathanPDB/chattronics/pkg/prompts"
)

var engine *gpt.Engine

func StartConversation(gptEngine *gpt.Engine) error {
	engine = gptEngine

	greetingsMessage()
	initialUserPrompt := readConsole()

	reviewedPrompt, err := reviewerChat(initialUserPrompt)
	if err != nil {
		return fmt.Errorf("ReviewerChat: %w", err)
	}

	architectureDiagram, err := architectChat(reviewedPrompt)
	if err != nil {
		return fmt.Errorf("ArchitectChat: %w", err)
	}

	err = graph.RenderGraph(architectureDiagram)
	if err != nil {
		return fmt.Errorf("RenderGraph: %w", err)
	}
	return nil
}

func reviewerChat(initialUserPrompt string) (string, error) {
	reviewer := gpt.NewChatModel(*engine, "reviewer", prompts.ReviewerSystem)

	reviewer.AddUserMessage(initialUserPrompt)

	response, err := reviewer.SendChat()
	if err != nil {
		return "", fmt.Errorf("failed to send first reviewer chat completion: %w", err)
	}

	questions := extractSingleBlock(response, "Questions").Lines
	answers := askQuestions(questions)

	newUserPrompt := prompts.FormulatePrefix + answers + prompts.FormulateSuffix
	reviewer.AddUserMessage(newUserPrompt)

	response, err = reviewer.SendChat()
	if err != nil {
		return "", fmt.Errorf("failed to send second reviewer chat completion: %w", err)
	}

	return extractSingleBlock(response, "Prompt").Block, nil
}

func architectChat(reformulatedPrompt string) (string, error) {
	architect := gpt.NewChatModel(*engine, "architect", prompts.ArchitectSystem)
	architect.AddUserMessage(reformulatedPrompt)

	response, err := architect.SendChat()
	if err != nil {
		return "", fmt.Errorf("failed to send architect chat completion: %w", err)
	}

	return extractSingleBlock(response, "dot").Block, nil
}
