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
	for node := finish; node != nil; node = node.previous {
		if node.previous != nil {
			if node.reversed && node.previous.reversed {
				temp := node.previous.key + " " + node.key
				crossings = append(crossings, temp)
			}
		}
		if node != nil && node.previous != nil {
			g.deleteEdge(node, node.previous)
			g.AddOneDirectedEdge(node, node.previous)
		}
		node.reversed = true
	}
	return crossings
}
