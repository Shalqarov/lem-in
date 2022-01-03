package algorithms

import (
	"fmt"
	"strings"
)

func (g *Graph) FindAvailablePaths(cloneGraph *Graph, ants int) ([][]*Vertex, error) {
	if ants <= 2 {
		path, err := g.oneWaySearch()
		if err != nil {
			return nil, fmt.Errorf(err.Error())
		}
		return path, nil
	}
	crossings, err := cloneGraph.findingCrossings()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	g.deleteCrossings(crossings)

	foundPaths := g.pathsSearch()
	return foundPaths, nil
}

func (g *Graph) oneWaySearch() ([][]*Vertex, error) {
	path, _ := g.BFS(g.Start, g.End)
	if path == nil {
		return nil, fmt.Errorf("no available paths")
	}
	//without start
	res := [][]*Vertex{path[1:]}
	return res, nil
}

func (g *Graph) pathsSearch() [][]*Vertex {
	foundPaths := [][]*Vertex{}
	visitedVertices := []map[*Vertex]bool{}
	for {
		path, mapPath := g.BFS(g.Start, g.End)
		if path == nil {
			break
		}
		// if path has only start - end
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
			continue
		}
		visitedVertices = append(visitedVertices, mapPath)
		foundPaths = append(foundPaths, path[1:])
	}
	return foundPaths
}

func (g *Graph) findingCrossings() ([]string, error) {
	_, pathFound := g.BhandariCrossings(g.Start, g.End)
	if !pathFound {
		return nil, fmt.Errorf("invalid data format")
	}
	crossings := []string{}
	for {
		crossedVertices, pathFound := g.BhandariCrossings(g.Start, g.End)
		if !pathFound {
			break
		}
		crossings = append(crossings, crossedVertices...)
	}
	return crossings, nil
}

func (g *Graph) deleteCrossings(crossings []string) {
	for _, v := range crossings {
		temp := strings.Split(v, " ")
		g.deleteEdge(g.GetVertex(temp[0]), g.GetVertex(temp[1]))
	}
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
				return g.reversePath(v)
			}
			queue = append(queue, v)
		}
		queue = queue[1:]
	}
	return nil, nil
}

func (g *Graph) reversePath(finish *Vertex) ([]*Vertex, map[*Vertex]bool) {
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
	delete(mapResult, res[0])
	delete(mapResult, res[len(res)-1])
	for i := 1; i < len(res); i++ {
		g.deleteEdge(res[i], res[i].previous)
		g.AddOneDirectedEdge(res[i], res[i].previous)
	}
	return res, mapResult
}
