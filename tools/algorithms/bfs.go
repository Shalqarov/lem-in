package algorithms

import (
	"fmt"
	"strings"
)

func (g *Graph) FindAvailablePaths(copiedGraph *Graph, from, to string, ants int) ([][]*Vertex, error) {
	if ants <= 2 {
		path, err := g.oneWaySearch(from, to)
		if err != nil {
			return nil, fmt.Errorf(err.Error())
		}
		return path, nil
	}
	crossings := []string{}
	for i := 0; ; i++ {
		crossedVertices, pathFound := copiedGraph.FindingCrossings(copiedGraph.getVertex(from), copiedGraph.getVertex(to))
		if !pathFound {
			if i == 0 {
				return nil, fmt.Errorf("no available paths")
			}
			break
		}
		crossings = append(crossings, crossedVertices...)
	}

	for _, v := range crossings {
		temp := strings.Split(v, " ")
		g.deleteEdge(g.getVertex(temp[0]), g.getVertex(temp[1]))
	}

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
			return foundPaths, nil
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
	return foundPaths, nil
}

func (g *Graph) oneWaySearch(from, to string) ([][]*Vertex, error) {
	path, _ := g.BFS(g.getVertex(from), g.getVertex(to))
	if path == nil {
		return nil, fmt.Errorf("no available paths")
	}
	res := [][]*Vertex{path}
	return res, nil
}

func haveVerticesCrossings(visitedVertices []map[*Vertex]bool, path map[*Vertex]bool) bool {
	for _, v := range visitedVertices {
		if crossed(v, path) {
			return true
		}
	}
	return false
}

func crossed(path, currentpath map[*Vertex]bool) bool {
	for vertex := range currentpath {
		if _, have := path[vertex]; have {
			return true
		}
	}
	return false
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
		g.deleteEdge(res[i], res[i].previous)
		g.AddOneDirectedEdge(res[i], res[i].previous)
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
