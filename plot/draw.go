package plot

import (
	"github.com/fogleman/gg"
	"github.com/jamesrashford/graphkit/layout"
	"github.com/jamesrashford/graphkit/models"
)

// Turn this into a struct with methods
type GraphPlotter struct {
	Width, Height int
}

func NewGraphPlotter(width int, height int) *GraphPlotter {
	gp := GraphPlotter{
		Width:  width,
		Height: height,
	}
	return &gp
}

func (gp *GraphPlotter) Draw(graph *models.Graph, pos map[string]layout.Point, labels bool, filename string) {
	nodes := graph.GetNodes()
	edges := graph.GetEdges()

	const scale = 10
	const rad = 10

	dc := gg.NewContext(gp.Width, gp.Height)

	// Translate to centre
	dc.Translate((float64(gp.Width)/2.0)-scale, (float64(gp.Height)/2.0)-scale)

	dc.SetRGB(1, 1, 1)
	dc.Clear()

	// Draw edges
	dc.SetRGB(0, 0, 0)
	for _, edge := range edges {
		p1 := pos[edge.Source.ID]
		p2 := pos[edge.Target.ID]

		dc.DrawLine(p1.X*scale, p1.Y*scale, p2.X*scale, p2.Y*scale)
		dc.Stroke()
	}

	// Draw nodes
	for _, node := range nodes {
		p := pos[node.ID]
		dc.DrawCircle(p.X*scale, p.Y*scale, rad)
		dc.SetRGB(0, 0, 1)
		dc.Fill()

		// Draw labels
		if labels {
			dc.SetRGB(0, 0, 0)
			dc.DrawStringAnchored(node.Label, p.X*scale, p.Y*scale, 0.5, 0.5)
		}
	}

	// Output
	dc.SavePNG(filename)
}
