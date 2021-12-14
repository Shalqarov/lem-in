package algorithms

import (
	"strings"
)

func (g *Graph) FindAvailablePaths(copiedGraph *Graph, from, to string) [][]*Vertex {

	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	res := []string{}
	//DFS
	for path, pathFinding := copiedGraph.DFS(copiedGraph.getVertex(from), copiedGraph.getVertex(to)); pathFinding; path, pathFinding = copiedGraph.DFS(copiedGraph.getVertex(from), copiedGraph.getVertex(to)) {
		for _, v := range path {
			res = append(res, v)
		}
	}

	//deleting crossings
	for _, v := range res {
		temp := strings.Split(v, "-")
		g.deleteEdge(string(temp[0]), string(temp[1]))
	}

	//BFS
	foundPaths := [][]*Vertex{}
	visitChecking := []map[*Vertex]bool{}
	check := true
	for path, mapPath := g.BFS(g.getVertex(from), g.getVertex(to)); path != nil; path, mapPath = g.BFS(g.getVertex(from), g.getVertex(to)) {
		if len(path) == 2 {
			foundPaths = append(foundPaths, path)
			return foundPaths
		}
		if len(visitChecking) == 0 {
			visitChecking = append(visitChecking, mapPath)
			foundPaths = append(foundPaths, path)
			continue
		}
		for _, v := range visitChecking {
			if !crossingsChecking(v, mapPath, fromVertex, toVertex) {
				check = false
				continue
			}
		}
		if check {
			visitChecking = append(visitChecking, mapPath)
			foundPaths = append(foundPaths, path)
		}
	}
	return foundPaths
}

func crossingsChecking(path, currentpath map[*Vertex]bool, from, to *Vertex) bool {
	for vrtx := range currentpath {
		if vrtx == from || vrtx == to {
			continue
		}
		if _, have := path[vrtx]; have {
			return false
		}
	}
	return true
}

func (g *Graph) getpath(finish *Vertex) ([]*Vertex, map[*Vertex]bool) {
	reversed := []*Vertex{}
	for node := finish; node != nil; node = node.previous {
		node.reversed = true
		reversed = append(reversed, node)
	}
	res := make([]*Vertex, len(reversed))
	mapResult := make(map[*Vertex]bool)
	for i, j := len(reversed)-1, 0; i >= 0; i, j = i-1, j+1 {
		res[j] = reversed[i]
		mapResult[res[j]] = true
	}
	for i := 1; i < len(res); i++ {
		g.delEdges(res[i])
	}
	return res, mapResult
}

func (g *Graph) BFS(from, to *Vertex) ([]*Vertex, map[*Vertex]bool) {
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
				return g.getpath(v)
			}
			queue = append(queue, v)
		}
		queue = queue[1:]
	}
	return nil, nil
}
