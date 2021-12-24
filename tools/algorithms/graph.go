package algorithms

import "fmt"

type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	reversed  bool
	key       string
	adjacents []*Vertex
	previous  *Vertex
}

func (v *Vertex) GetAdjacents() []*Vertex {
	return v.adjacents
}

func (v *Vertex) GetKey() string {
	return v.key
}

func GraphInit() *Graph {
	return &Graph{}
}

func (g *Graph) AppendVertex(key string) {
	if g.containsVertex(key) {
		fmt.Printf("Oh snap: '%s' is already exists", key)
	}
	g.vertices = append(g.vertices, &Vertex{key: key})
}

func (g *Graph) AddEdge(from, to string, isOneDirection bool) {
	fromVertex := g.getVertex(from)
	if fromVertex == nil {
		fmt.Printf("Oh snap: '%s' don't exists", from)
		return
	}
	toVertex := g.getVertex(to)
	if toVertex == nil {
		fmt.Printf("Oh snap: '%s' don't exists", to)
		return
	}
	fromVertex.adjacents = append(fromVertex.adjacents, toVertex)
	if !isOneDirection {
		toVertex.adjacents = append(toVertex.adjacents, fromVertex)
	}
}

func (g *Graph) deleteEdge(from, to *Vertex, isChangeDirection bool) {
	for index, vertex := range from.adjacents {
		if vertex == to {
			from.adjacents = append(from.adjacents[:index], from.adjacents[index+1:]...)
			break
		}
	}
	for index, vertex := range to.adjacents {
		if vertex == from {
			to.adjacents = append(to.adjacents[:index], to.adjacents[index+1:]...)
			break
		}
	}
	if isChangeDirection {
		g.AddEdge(from.key, to.key, true)
	}
}

func (g *Graph) PrintGraph() {
	fmt.Println("#### PRINT GRAPH ####")
	for _, vertex := range g.vertices {
		fmt.Printf("# Vertex %v : ", vertex.key)
		for _, vertex := range vertex.adjacents {
			fmt.Printf("%v ", vertex.key)
		}
		fmt.Println()
	}
	fmt.Println("#####################")
}

func (g *Graph) getVertex(key string) *Vertex {
	for _, vertex := range g.vertices {
		if vertex.key == key {
			return vertex
		}
	}
	return nil
}

func (g *Graph) containsVertex(key string) bool {
	for _, vertex := range g.vertices {
		if vertex.key == key {
			return true
		}
	}
	return false
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
