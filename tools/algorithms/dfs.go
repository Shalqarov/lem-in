package algorithms

func (g *Graph) DFS(from, to *Vertex) ([]string, bool) {
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
				cross := g.reversepath(v)
				// PrintPath(temp)
				return cross, true
			}
			queue = append(queue, v)
		}

		queue = queue[1:]
	}
	return nil, false
}

func (g *Graph) reversepath(finish *Vertex) []string {
	reversed := []*Vertex{}
	crossings := []string{}
	for node := finish; node != nil; node = node.previous {
		if node != nil && node.previous != nil {
			if node.reversed && node.previous.reversed {
				temp := node.previous.key + "-" + node.key
				crossings = append(crossings, temp)
			}
		}
		node.reversed = true
		reversed = append(reversed, node)
	}
	res := make([]*Vertex, len(reversed))
	for i, j := len(reversed)-1, 0; i >= 0; i, j = i-1, j+1 {
		res[j] = reversed[i]
	}
	for i := 1; i < len(res); i++ {
		g.delEdges(res[i])
	}
	return crossings
}