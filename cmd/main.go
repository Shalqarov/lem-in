package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Lem struct {
	StartRoom string
	EndRoom   string
	Rooms     []string
	Edges     []string
	Ants      int
}

type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	key      string
	adj      []*Vertex
	visited  bool
	previous *Vertex
}

// func (g *Graph) deleteEdge(from, to *Vertex) {
// 	toVrtx := g.getVertex(to)
// }

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
	fromVrtx.adj = append(fromVrtx.adj, toVrtx)
	toVrtx.adj = append(toVrtx.adj, fromVrtx)
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
	fromVrtx.adj = append(fromVrtx.adj, toVrtx)
}

func (g *Graph) swapEdges(v *Vertex) {
	for i, val := range v.adj {
		if val == v {
			v.adj = append(v.adj[:i], v.adj[i+1:]...)
			break
		}
	}
	for i, val := range v.previous.adj {
		if val == v {
			v.adj = append(v.adj[:i], v.adj[i+1:]...)
			break
		}
	}
	g.AddEdgeOneDir(v.key, v.previous.key)
	v.previous = nil
}

// func (g *Graph) BFS(from, to *Vertex) {
// 	visited := map[*Vertex]bool{from: true}
// 	queue := []*Vertex{from}
// 	foundPaths := [][]*Vertex{}

// 	for len(queue) > 0 {
// 		current := queue[0]
// 		for _, v := range current.adj {
// 			if visited[v] {
// 				continue
// 			}
// 			visited[v] = true
// 			v.previous = current
// 			if v == to {
// 				foundPaths = append(foundPaths, g.getpath(v))
// 				visited[v] = false
// 			}
// 			queue = append(queue, v)
// 		}
// 		queue = queue[1:]
// 	}
// 	if len(foundPaths) == 0 {
// 		fmt.Println("Paths not found")
// 		return
// 	}
// 	for _, v := range foundPaths {
// 		printPath(v)
// 	}
// }

func (g *Graph) BFS(from, to *Vertex) {
	visited := map[*Vertex]bool{from: true}
	queue := []*Vertex{from}
	Path := []*Vertex{}

	for len(queue) > 0 {
		current := queue[0]
		for _, v := range current.adj {
			if visited[v] {
				continue
			}
			visited[v] = true
			v.previous = current
			if v == to {
				printPath(g.getpath(v))
				return
			}
			queue = append(queue, v)
		}
		queue = queue[1:]
	}
	if len(Path) == 0 {
		fmt.Println("Paths not found")
		return
	}

}

func printPath(path []*Vertex) {
	if len(path) == 0 {
		return
	}
	fmt.Print(path[0].key)
	for i := 1; i < len(path); i++ {
		fmt.Printf(" --> %s", path[i].key)
	}
	fmt.Println()
}

func (g *Graph) getpath(finish *Vertex) []*Vertex {
	reversed := []*Vertex{}
	for node := finish; node != nil; node = node.previous {
		reversed = append(reversed, node)
	}
	res := make([]*Vertex, len(reversed))
	for i, j := len(reversed)-1, 0; i >= 0; i, j = i-1, j+1 {
		res[j] = reversed[i]
	}
	for i := 1; i < len(res); i++ {
		fmt.Println(res[i].key)
		g.swapEdges(res[i])
	}
	return res
}

func (g *Graph) PrintGraph() {
	fmt.Println("#### PRINT GRAPH ####")
	for _, v := range g.vertices {
		fmt.Printf("# Vertex %v : ", v.key)
		for _, v := range v.adj {
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

func main() {
	lem := Lem{}
	args := os.Args[1:]
	file, err := os.Open("test/" + args[0])
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		if scanner.Text() == "##start" {
			fmt.Println(scanner.Text())
			scanner.Scan()
			temp := strings.Split(scanner.Text(), " ")
			lem.StartRoom = temp[0]
			fmt.Println(scanner.Text())
			continue
		}
		if scanner.Text() == "##end" {
			fmt.Println(scanner.Text())
			scanner.Scan()
			temp := strings.Split(scanner.Text(), " ")
			lem.EndRoom = temp[0]
			fmt.Println(scanner.Text())
			continue
		}
		if strings.Contains(scanner.Text(), "-") {
			lem.Edges = append(lem.Edges, scanner.Text())
			fmt.Println(scanner.Text())
			continue
		}
		if len(scanner.Text()) == 1 {
			temp, _ := strconv.Atoi(scanner.Text())
			lem.Ants = temp
			continue
		}
		temp := strings.Split(scanner.Text(), " ")
		lem.Rooms = append(lem.Rooms, temp[0])
		fmt.Println(scanner.Text())
	}
	g := Graph{}
	g.AddVertex(lem.StartRoom)
	g.AddVertex(lem.EndRoom)
	for _, v := range lem.Rooms {
		g.AddVertex(v)
	}
	for _, v := range lem.Edges {
		temp := strings.Split(v, "-")
		g.AddEdge(temp[0], temp[1])
	}

	g.PrintGraph()

	g.BFS(g.getVertex("1"), g.getVertex("0"))
	g.BFS(g.getVertex("1"), g.getVertex("0"))
	g.BFS(g.getVertex("1"), g.getVertex("0"))
}
