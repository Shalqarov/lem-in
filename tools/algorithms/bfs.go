package algorithms

import (
	"strings"
)

func (g *Graph) FindAvailablePaths(copiedGraph *Graph, from, to string, ants int) [][]*Vertex {
	crossings := []string{}

	if ants <= 2 {
		temp, _ := g.BFS(g.getVertex(from), g.getVertex(to))
		res := [][]*Vertex{temp}
		return res
	}

	for {
		crossedVertices, pathFound := copiedGraph.DFS(copiedGraph.getVertex(from), copiedGraph.getVertex(to))
		if !pathFound {
			break
		}
		crossings = append(crossings, crossedVertices...)
	}

	//deleting crossings
	for _, v := range crossings {
		temp := strings.Split(v, " ")
		g.deleteEdge(g.getVertex(temp[0]), g.getVertex(temp[1]), false)
	}

	//BFS
	foundPaths := [][]*Vertex{}
	visitedVertices := []map[*Vertex]bool{}
	for {
		haveCrossings := false
		path, mapPath := g.BFS(g.getVertex(from), g.getVertex(to))
		if path == nil {
			break
		}
		if len(path) == 2 {
			// returning path without start vertex
			foundPaths = append(foundPaths, path[1:])
			return foundPaths
		}
		if len(visitedVertices) == 0 {
			visitedVertices = append(visitedVertices, mapPath)
			foundPaths = append(foundPaths, path[1:])
			continue
		}
		if haveVerticesCrossings(visitedVertices, mapPath) {
			haveCrossings = true
			continue
		}
		if !haveCrossings {
			visitedVertices = append(visitedVertices, mapPath)
			foundPaths = append(foundPaths, path[1:])
		}
	}
	return foundPaths
}

func haveVerticesCrossings(visitedVertices []map[*Vertex]bool, path map[*Vertex]bool) bool {
	for _, v := range visitedVertices {
		if !crossingsChecking(v, path) {
			return true
		}
	}
	return false
}

func crossingsChecking(path, currentpath map[*Vertex]bool) bool {
	for vertex := range currentpath {
		if _, have := path[vertex]; have {
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
		if j == 0 || i+1 >= 0 {
			continue
		}
		mapResult[res[j]] = true
	}
	for i := 1; i < len(res); i++ {
		g.deleteEdge(res[i], res[i].previous, true)
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
