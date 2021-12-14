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

func (g *Graph) AddVertex(key string) {
	if g.contains(key) {
		fmt.Printf("Oh snap: '%s' is already exists", key)
	}
	g.vertices = append(g.vertices, &Vertex{key: key})
}

func (g *Graph) AddEdge(from, to string) {
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
	toVrtx.adjacents = append(toVrtx.adjacents, fromVrtx)
}

func (g *Graph) AddEdgeOneDir(from, to string) {
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
}

func (g *Graph) swapEdges(v *Vertex) {
	for i, val := range v.adjacents {
		if val == v {
			v.adjacents = append(v.adjacents[:i], v.adjacents[i+1:]...)
			break
		}
	}
	for i, val := range v.previous.adjacents {
		if val == v {
			v.adjacents = append(v.adjacents[:i], v.adjacents[i+1:]...)
			break
		}
	}
	g.AddEdgeOneDir(v.key, v.previous.key)
	v.previous = nil
}

func (g *Graph) delEdges(v *Vertex) {
	for i, val := range v.adjacents {
		if val == v.previous {
			v.adjacents = append(v.adjacents[:i], v.adjacents[i+1:]...)
			break
		}
	}
	for i, val := range v.previous.adjacents {
		if val == v {
			v.previous.adjacents = append(v.previous.adjacents[:i], v.previous.adjacents[i+1:]...)
			break
		}
	}
	g.AddEdgeOneDir(v.key, v.previous.key)
}

func (g *Graph) deleteEdge(from, to string) {
	fromVrtx := g.getVertex(from)
	toVrtx := g.getVertex(to)

	for i, val := range fromVrtx.adjacents {
		if val == toVrtx {
			fromVrtx.adjacents = append(fromVrtx.adjacents[:i], fromVrtx.adjacents[i+1:]...)
			break
		}
	}
	for i, val := range toVrtx.adjacents {
		if val == fromVrtx {
			toVrtx.adjacents = append(toVrtx.adjacents[:i], toVrtx.adjacents[i+1:]...)
			break
		}
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