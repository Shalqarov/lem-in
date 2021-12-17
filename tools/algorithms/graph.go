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

func GraphInit() *Graph {
	return &Graph{}
}

func (g *Graph) AddVertex(key string) {
	if g.contains(key) {
		fmt.Printf("Oh snap: '%s' is already exists", key)
	}
	g.vertices = append(g.vertices, &Vertex{key: key})
}

func (g *Graph) AddEdge(from, to string, isOneDirection bool) {
	fromVrtx := g.getVertex(from)
	if fromVrtx == nil {
		fmt.Printf("Oh snap: '%s' don't exists", from)
		return
	}
	toVrtx := g.getVertex(to)
	if toVrtx == nil {
		fmt.Printf("Oh snap: '%s' don't exists", to)
		return
	}
	fromVrtx.adjacents = append(fromVrtx.adjacents, toVrtx)
	if !isOneDirection {
		toVrtx.adjacents = append(toVrtx.adjacents, fromVrtx)
	}
}

func (g *Graph) deleteEdge(from, to *Vertex, isChangeDirection bool) {
	for i, val := range from.adjacents {
		if val == to {
			from.adjacents = append(from.adjacents[:i], from.adjacents[i+1:]...)
			break
		}
	}
	for i, val := range to.adjacents {
		if val == from {
			to.adjacents = append(to.adjacents[:i], to.adjacents[i+1:]...)
			break
		}
	}
	if isChangeDirection {
		g.AddEdge(from.key, to.key, true)
	}
}

func (g *Graph) PrintGraph() {
	fmt.Println("#### PRINT GRAPH ####")
	for _, v := range g.vertices {
		fmt.Printf("# Vertex %v : ", v.key)
		for _, v := range v.adjacents {
			fmt.Printf("%v ", v.key)
		}
		fmt.Println()
	}
	fmt.Println("#####################")
}

func (g *Graph) getVertex(key string) *Vertex {
	for _, v := range g.vertices {
		if v.key == key {
			return v
		}
	}
	return nil
}

func (g *Graph) contains(key string) bool {
	for _, v := range g.vertices {
		if v.key == key {
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
