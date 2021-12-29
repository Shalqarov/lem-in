package algorithms

import "fmt"

type Graph struct {
	vertices map[string]*Vertex
	Start    *Vertex
	End      *Vertex
}

type Vertex struct {
	reversed  bool
	key       string
	adjacents []*Vertex
	previous  *Vertex
}

func (g *Graph) SetStart(key string) {
	g.Start = g.GetVertex(key)
}
func (g *Graph) SetEnd(key string) {
	g.End = g.GetVertex(key)
}

func (v *Vertex) GetKey() string {
	return v.key
}

func GraphInit() *Graph {
	temp := &Graph{
		vertices: make(map[string]*Vertex),
	}
	return temp
}
func (g *Graph) AppendVertex(key string) error {
	temp := &Vertex{key: key}
	if _, isHave := g.vertices[key]; isHave {
		return fmt.Errorf("already has vertex")
	}
	g.vertices[key] = temp
	return nil
}

func (g *Graph) AddEdge(from, to *Vertex) {
	from.adjacents = append(from.adjacents, to)
	to.adjacents = append(to.adjacents, from)
}

func (g *Graph) AddOneDirectedEdge(from, to *Vertex) {
	from.adjacents = append(from.adjacents, to)
}

func (g *Graph) deleteEdge(from, to *Vertex) {
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
}

func (g *Graph) GetVertex(key string) *Vertex {
	return g.vertices[key]
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
