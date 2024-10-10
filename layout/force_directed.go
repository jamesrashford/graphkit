package layout

import (
	"math"
	"math/rand"

	"github.com/jamesrashford/graphkit/models"
)

func repulsiveForce(k, dist float64) float64 {
	if dist == 0 {
		return 0
	}
	return k * k / dist
}

func attractiveForce(k, dist float64) float64 {
	return dist * dist / k
}

func ForceDirected(graph *models.Graph, posPrev map[string]Point, iterations int, k float64, t_start float64) map[string]Point {
	nodes := graph.GetNodes()
	edges := graph.GetEdges()

	pos := make(map[string]Point)

	// Load posPrev, if exists, into pos
	if posPrev == nil {
		for _, n := range nodes {
			pos[n.ID] = Point{rand.Float64(), rand.Float64()}
		}
	}

	for _, n := range nodes {
		if _, ok := posPrev[n.ID]; !ok {
			pos[n.ID] = Point{rand.Float64(), rand.Float64()}
		} else {
			pos[n.ID] = posPrev[n.ID]
		}
	}

	disp := make(map[string]Point)

	t := t_start

	for iter := 0; iter < iterations; iter++ {
		for _, n := range nodes {
			disp[n.ID] = Point{0.0, 0.0}
		}

		for _, u := range nodes {
			for _, v := range nodes {
				dx := pos[u.ID].X - pos[v.ID].X
				dy := pos[u.ID].Y - pos[v.ID].Y
				dist := Distance(pos[u.ID], pos[v.ID])

				if dist > 0 {
					repForce := repulsiveForce(k, dist)
					disp_u := disp[u.ID]
					disp_u.X += (dx / dist) * repForce
					disp_u.Y += (dy / dist) * repForce

					disp_v := disp[v.ID]
					disp_v.X -= (dx / dist) * repForce
					disp_v.Y -= (dy / dist) * repForce
				}
			}
		}

		for _, edge := range edges {
			source := edge.Source.ID
			target := edge.Target.ID

			dx := pos[source].X - pos[target].X
			dy := pos[source].Y - pos[target].Y
			dist := Distance(pos[source], pos[target])

			if dist > 0 {
				attrForce := attractiveForce(k, dist)
				disp_s := disp[source]
				disp_s.X = disp_s.X - (dx/dist)*attrForce
				disp_s.Y -= (dy / dist) * attrForce

				disp_t := disp[target]
				disp_t.X += (dx / dist) * attrForce
				disp_t.Y += (dy / dist) * attrForce
			}
		}

		for _, n := range nodes {
			dm := Distance(Point{0.0, 0.0}, disp[n.ID])
			pos_n := pos[n.ID]
			if dm > 0 {
				pos_n.X += (disp[n.ID].X / dm) * math.Min(dm, t)
				pos_n.Y += (disp[n.ID].Y / dm) * math.Min(dm, t)
			}

		}

		t *= 0.95
	}

	return pos
}
