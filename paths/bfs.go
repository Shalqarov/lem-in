package graphs

import (
	"fmt"
	"strings"
)

func (g *Graph) AllPathsBFS(copiedGraph *Graph, from, to string) [][]*Vertex {

	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	res := []string{}
	for path, pathFinding := copiedGraph.BFS(copiedGraph.getVertex(from), copiedGraph.getVertex(to)); pathFinding; path, pathFinding = copiedGraph.BFS(copiedGraph.getVertex(from), copiedGraph.getVertex(to)) {
		for _, v := range path {
			res = append(res, v)
		}
	}

	for _, v := range res {
		temp := strings.Split(v, "-")
		g.deleteEdge(string(temp[0]), string(temp[1]))
	}

	visited := map[*Vertex]bool{fromVertex: true}
	queue := []*Vertex{fromVertex}
	foundPaths := [][]*Vertex{}
	visitChecking := []map[*Vertex]bool{}

	for len(queue) > 0 {
		current := queue[0]
		for _, vertex := range current.adjacents {
			check := true
			if visited[vertex] {
				continue
			}
			visited[vertex] = true
			vertex.previous = current
			if vertex == toVertex {
				tempPath, tempMapPath := g.getpath(vertex)
				if len(tempPath) == 2 {
					foundPaths = append(foundPaths, tempPath)
					return foundPaths
				}
				if len(visitChecking) == 0 {
					visitChecking = append(visitChecking, tempMapPath)
					foundPaths = append(foundPaths, tempPath)
					visited[vertex] = false
					continue
				}
				for _, v := range visitChecking {
					if !crossingsChecking(v, tempMapPath, fromVertex, toVertex) {
						visited[vertex] = false
						check = false
						break
					}
				}
				if check {
					visitChecking = append(visitChecking, tempMapPath)
					foundPaths = append(foundPaths, tempPath)
					visited[vertex] = false
				}
			}
			queue = append(queue, vertex)
		}
		queue = queue[1:]
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
		reversed = append(reversed, node)
	}
	res := make([]*Vertex, len(reversed))
	mapResult := make(map[*Vertex]bool)
	for i, j := len(reversed)-1, 0; i >= 0; i, j = i-1, j+1 {
		res[j] = reversed[i]
		mapResult[res[j]] = true
	}
	return res, mapResult
}

func (g *Graph) BFS(from, to *Vertex) ([]string, bool) {
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
				temp, cross := g.reversepath(v)
				PrintPath(temp)
				return cross, true
			}
			queue = append(queue, v)
		}

		queue = queue[1:]
	}
	fmt.Println("All available paths has been found")
	return nil, false
}

func PrintPath(path []*Vertex) {
	if len(path) == 0 {
		return
	}
	fmt.Print(path[0].key)
	for i := 1; i < len(path); i++ {
		fmt.Printf(" --> %s", path[i].key)
	}
	fmt.Println()
}

func (g *Graph) reversepath(finish *Vertex) ([]*Vertex, []string) {
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
	return res, crossings
}
