package graph

import (
	"github.com/stretchr/testify/assert"
	"new-chattronics/internal/config"
	"testing"
)

func getMockGraphvizDiagram() string {
	return `digraph example_graph {
    // Define node styles
    node [shape=ellipse, fontname="Arial", color="blue"];

    // Define nodes
    A [label="Node A"];
    B [label="Node B"];
    C [label="Node C"];
    D [label="Node D"];

    // Define edges
    A -> B [label="Edge 1", color="red"];
    B -> C [label="Edge 2", color="green"];
    C -> D [label="Edge 3", color="blue"];
    D -> A [label="Edge 4", color="purple"];
}`
}

func TestRenderGraph(t *testing.T) {
	t.Run("Successfully render graphviz diagram", func(t *testing.T) {
		config.CreateFolders("", false)
		err := RenderGraph(getMockGraphvizDiagram(), "diagram")
		assert.NoError(t, err)
		assert.FileExists(t, config.RunFolderPath+"diagram.svg")
		assert.FileExists(t, config.RunFolderPath+"diagram.png")
	})
}
