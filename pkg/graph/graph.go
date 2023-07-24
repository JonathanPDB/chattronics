package graph

import (
	"bytes"
	"fmt"
	"github.com/JonathanPDB/chattronics/pkg/config"
	"github.com/goccy/go-graphviz"
)

func RenderGraph(graphString string) error {
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

	//image, err := g.RenderImage(graph)
	//if err != nil {
	//	log.Fatal(err)
	//}

	diagramFile := config.RunFolderPath + "diagram.svg"

	if err = g.RenderFilename(graph, graphviz.SVG, diagramFile); err != nil {
		return fmt.Errorf("failed to render graph into file: %w", err)
	}

	return nil
}
