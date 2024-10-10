package layout

import "github.com/jamesrashford/graphkit/models"

func repulsiveForce(k, dist float64) float64 {
	if dist == 0 {
		return 0
	}
	return k * k / dist
}

func attractiveForce(k, dist float64) float64 {
	return dist * dist / k
}

func ForceDirected(graph *models.Graph, posPrev map[*models.Node]Point, iterations int, k float64, t_start float64) (map[*models.Node]Point, error) {
	nodes := graph.GetNodes()
	edges := graph.GetEdges()

	pos := make(map[*models.Node]Point)

	// Load posPrev, if exists, into pos

	disp := make(map[*models.Node]Point)

	t := 0.1

	for iter := 0; iter < iterations; iter++ {
		for _, n := range nodes {
			disp[&n] = Point{0.0, 0.0}
		}

		for _, u := range nodes {
			for _, v := range nodes {
				dx := pos[&u].X - pos[&v].X
				dy := pos[&u].Y - pos[&v].Y
				dist := Distance(pos[&u], pos[&v])

				if dist > 0 {
					repForce := repulsiveForce(k, dist)
					disp_u := disp[&u]
					disp_u.X += (dx / dist) * repForce
					disp_u.Y += (dy / dist) * repForce

					disp_v := disp[&v]
					disp_v.X -= (dx / dist) * repForce
					disp_v.Y -= (dy / dist) * repForce
				}
			}
		}
	}

	return pos, nil
}
