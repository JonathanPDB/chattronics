package graph

import (
	"bytes"
	"fmt"
	"github.com/chattronics/chattronics/internal/config"
	"github.com/chattronics/chattronics/internal/logging"
	"github.com/goccy/go-graphviz"
	"os/exec"
)

func RenderGraph(graphString, filename string) error {
	graphBytes := []byte(graphString)

	g := graphviz.New()
	graph, err := graphviz.ParseBytes(graphBytes)
	if err != nil {
		return fmt.Errorf("failed to parse graph bytes: %w", err)
	}

	var buf bytes.Buffer
	if err = g.Render(graph, graphviz.SVG, &buf); err != nil {
		return fmt.Errorf("failed to render graph: %w", err)
	}

	diagramSvgFile := config.RunFolderPath + filename + ".svg"
	diagramPngFile := config.RunFolderPath + filename + ".png"

	if err = g.RenderFilename(graph, graphviz.SVG, diagramSvgFile); err != nil {
		return fmt.Errorf("failed to render graph into SVG file: %w", err)
	}

	if err = g.RenderFilename(graph, graphviz.PNG, diagramPngFile); err != nil {
		return fmt.Errorf("failed to render graph into PNG file: %w", err)
	}
	openFile(diagramPngFile)

	return nil
}

func openFile(filepath string) {
	cmd := exec.Command("open", filepath)
	_, err := cmd.Output()
	if err != nil {
		logging.Warn(fmt.Sprintf("Failed to open %s", filepath))
	}
}
