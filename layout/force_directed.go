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
	return (k * k) / dist
}

func attractiveForce(k, dist float64) float64 {
	return (dist * dist) / k
}

func ForceDirected(graph *models.Graph, iterations int, k float64, t_start float64) {
	nodes := graph.GetNodes()
	edges := graph.GetEdges()

	for _, n := range nodes {
		graph.Nodes[n.ID]["x"] = rand.Float64()
		graph.Nodes[n.ID]["y"] = rand.Float64()
	}

	disp := make(map[string]Point)

	t := t_start

	for iter := 0; iter < iterations; iter++ {
		for _, n := range nodes {
			disp[n.ID] = Point{0.0, 0.0}
		}

		for i, u := range nodes {
			for j, v := range nodes {
				if i != j {
					ux := graph.Nodes[u.ID]["x"].(float64)
					uy := graph.Nodes[u.ID]["y"].(float64)
					vx := graph.Nodes[v.ID]["x"].(float64)
					vy := graph.Nodes[v.ID]["y"].(float64)

					dx := ux - vx
					dy := uy - vy

					pu := Point{ux, uy}
					pv := Point{vx, vy}
					dist := Distance(pu, pv)

					if dist > 0 {
						repForce := repulsiveForce(k, dist)
						dv := disp[v.ID]
						disp[v.ID] = Point{
							X: dv.X + (dx/dist)*repForce,
							Y: dv.Y + (dy/dist)*repForce,
						}
					}
				}
			}
		}

		for _, edge := range edges {
			source := edge.Source.ID
			target := edge.Target.ID

			sx := graph.Nodes[source]["x"].(float64)
			sy := graph.Nodes[source]["y"].(float64)
			tx := graph.Nodes[target]["x"].(float64)
			ty := graph.Nodes[target]["y"].(float64)

			dx := sx - tx
			dy := sy - ty

			ps := Point{sx, sy}
			pt := Point{tx, ty}
			dist := Distance(ps, pt)

			if dist > 0 {
				attrForce := attractiveForce(k, dist)
				ds := disp[source]
				disp[source] = Point{
					X: ds.X - (dx/dist)*attrForce,
					Y: ds.Y - (dy/dist)*attrForce,
				}

				dt := disp[target]
				disp[target] = Point{
					X: dt.X + (dx/dist)*attrForce,
					Y: dt.Y + (dy/dist)*attrForce,
				}
			}
		}

		for _, n := range nodes {
			dm := Distance(Point{0.0, 0.0}, disp[n.ID])
			nx := graph.Nodes[n.ID]["x"].(float64)
			ny := graph.Nodes[n.ID]["y"].(float64)

			if dm > 0 {
				graph.Nodes[n.ID]["x"] = nx + (nx/dm)*math.Min(dm, t)
				graph.Nodes[n.ID]["y"] = ny + (ny/dm)*math.Min(dm, t)
			}

		}

		t *= 0.95

	}

}
