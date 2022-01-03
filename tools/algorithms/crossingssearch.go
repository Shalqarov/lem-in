package algorithms

func (g *Graph) BhandariCrossings(from, to *Vertex) ([]string, bool) {
	visited := map[*Vertex]bool{from: true}
	queue := []*Vertex{from}

	for len(queue) > 0 {
		current := queue[0]
		for _, v := range current.adjacents {
			if visited[v] {
				continue
			}
			visited[v] = true
			v.previous = current
			if v == to {
				cross := g.reversePathWithCrossings(v)
				return cross, true
			}
			queue = append(queue, v)
		}
		queue = queue[1:]
	}
	return nil, false
}

func (g *Graph) reversePathWithCrossings(finish *Vertex) []string {
	crossings := []string{}
	for vertex := finish; vertex != nil; vertex = vertex.previous {
		if vertex.previous != nil {
			if vertex.reversed && vertex.previous.reversed {
				crossings = append(crossings, vertex.previous.key+" "+vertex.key)
			}
		}
		if vertex != nil && vertex.previous != nil {
			g.deleteEdge(vertex, vertex.previous)
			g.AddOneDirectedEdge(vertex, vertex.previous)
		}
		vertex.reversed = true
	}
	return crossings
}
