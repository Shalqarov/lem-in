package main

import (
	"fmt"
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
	reversed bool
	visited  bool
	key      string
	adj      []*Vertex
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

func (g *Graph) delEdges(v *Vertex) {
	for i, val := range v.adj {
		if val == v.previous {
			v.adj = append(v.adj[:i], v.adj[i+1:]...)
			break
		}
	}
	for i, val := range v.previous.adj {
		if val == v {
			v.previous.adj = append(v.previous.adj[:i], v.previous.adj[i+1:]...)
			break
		}
	}
	g.AddEdgeOneDir(v.key, v.previous.key)
	v.visited = true
	v.previous.visited = true
}

func (g *Graph) deleteEdge(from, to string) {
	fromVrtx := g.getVertex(from)
	fmt.Println(fromVrtx.key)
	toVrtx := g.getVertex(to)
	fmt.Println(toVrtx.key)

	for i, val := range fromVrtx.adj {
		if val == toVrtx {
			fromVrtx.adj = append(fromVrtx.adj[:i], fromVrtx.adj[i+1:]...)
			break
		}
	}
	for i, val := range toVrtx.adj {
		if val == fromVrtx {
			toVrtx.adj = append(toVrtx.adj[:i], toVrtx.adj[i+1:]...)
			break
		}
	}
}

func (g *Graph) AllPathsBFS(from, to *Vertex) {
	visited := map[*Vertex]bool{from: true}
	queue := []*Vertex{from}
	foundPaths := [][]*Vertex{}
	visitChecking := []map[*Vertex]bool{}

	for len(queue) > 0 {
		current := queue[0]
		for _, vertex := range current.adj {
			check := true
			if visited[vertex] {
				continue
			}
			visited[vertex] = true
			vertex.previous = current
			if vertex == to {
				tempPath, tempMapPath := g.getpath(vertex)
				if len(visitChecking) == 0 {
					visitChecking = append(visitChecking, tempMapPath)
					foundPaths = append(foundPaths, tempPath)
					visited[vertex] = false
					continue
				}
				for _, v := range visitChecking {
					if !checker(v, tempMapPath, from, to) {
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
	if len(foundPaths) == 0 {
		fmt.Println("Paths not found")
		return
	}
	for _, v := range foundPaths {
		printPath(v)
	}
}

func checker(path, currentpath map[*Vertex]bool, from, to *Vertex) bool {
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
		for _, v := range current.adj {
			if visited[v] {
				continue
			}
			visited[v] = true
			v.previous = current
			if v == to {
				temp, cross := g.reversepath(v)
				printPath(temp)
				return cross, true
			}
			queue = append(queue, v)
		}
		queue = queue[1:]
	}
	fmt.Println("All available paths has been found")
	return nil, false
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
	// lem := Lem{}
	// args := os.Args[1:]
	// file, err := os.Open("test/" + args[0])
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer file.Close()
	// for scanner := bufio.NewScanner(file); scanner.Scan(); {
	// 	if scanner.Text() == "##start" {
	// 		fmt.Println(scanner.Text())
	// 		scanner.Scan()
	// 		temp := strings.Split(scanner.Text(), " ")
	// 		lem.StartRoom = temp[0]
	// 		fmt.Println(scanner.Text())
	// 		continue
	// 	}
	// 	if scanner.Text() == "##end" {
	// 		fmt.Println(scanner.Text())
	// 		scanner.Scan()
	// 		temp := strings.Split(scanner.Text(), " ")
	// 		lem.EndRoom = temp[0]
	// 		fmt.Println(scanner.Text())
	// 		continue
	// 	}
	// 	if strings.Contains(scanner.Text(), "-") {
	// 		lem.Edges = append(lem.Edges, scanner.Text())
	// 		fmt.Println(scanner.Text())
	// 		continue
	// 	}
	// 	if len(scanner.Text()) == 1 {
	// 		temp, _ := strconv.Atoi(scanner.Text())
	// 		lem.Ants = temp
	// 		continue
	// 	}
	// 	temp := strings.Split(scanner.Text(), " ")
	// 	lem.Rooms = append(lem.Rooms, temp[0])
	// 	fmt.Println(scanner.Text())
	// }
	// g := Graph{}
	// g.AddVertex(lem.StartRoom)
	// g.AddVertex(lem.EndRoom)
	// for _, v := range lem.Rooms {
	// 	g.AddVertex(v)
	// }
	// for _, v := range lem.Edges {
	// 	temp := strings.Split(v, "-")
	// 	g.AddEdge(temp[0], temp[1])
	// }

	g := &Graph{}
	r := &Graph{}
	for i := '1'; i <= '8'; i++ {
		g.AddVertex(string(i))
	}
	g.AddVertex("9")
	r.AddVertex("9")
	for i := '1'; i <= '8'; i++ {
		r.AddVertex(string(i))
	}

	g.AddEdge("1", "2")
	g.AddEdge("1", "3")
	g.AddEdge("9", "3")
	g.AddEdge("4", "9")
	g.AddEdge("2", "4")
	g.AddEdge("5", "4")
	g.AddEdge("6", "4")
	g.AddEdge("6", "7")
	g.AddEdge("5", "7")

	r.AddEdge("1", "2")
	r.AddEdge("1", "3")
	r.AddEdge("9", "3")
	r.AddEdge("4", "9")
	r.AddEdge("2", "4")
	r.AddEdge("5", "4")
	r.AddEdge("6", "4")
	r.AddEdge("6", "7")
	r.AddEdge("5", "7")

	// mainG := g
	g.PrintGraph()

	// b := true
	res := []string{}
	for s, b := g.BFS(g.getVertex("1"), g.getVertex("7")); b; s, b = g.BFS(g.getVertex("1"), g.getVertex("7")) {
		// s, b := g.BFS(g.getVertex("1"), g.getVertex("7"))
		// if !b {
		// 	break
		// }
		for _, v := range s {
			res = append(res, v)
		}
	}
	for _, v := range res {
		r.deleteEdge(string(v[0]), string(v[2]))
	}
	r.AllPathsBFS(r.getVertex("1"), r.getVertex("7"))
}

// g := &Graph{}
// r := &Graph{}
// for i := 'a'; i <= 'n'; i++ {
// 	g.AddVertex(string(i))
// }
// for i := 'a'; i <= 'n'; i++ {
// 	r.AddVertex(string(i))
// }

// g.AddEdge("b", "a")
// g.AddEdge("d", "a")
// g.AddEdge("d", "i")
// g.AddEdge("j", "i")
// g.AddEdge("b", "e")
// g.AddEdge("b", "c")
// g.AddEdge("f", "c")
// g.AddEdge("e", "h")
// g.AddEdge("e", "g")
// g.AddEdge("h", "l")
// g.AddEdge("n", "l")
// g.AddEdge("n", "m")
// g.AddEdge("g", "f")
// g.AddEdge("g", "j")

// g.AddEdge("g", "k")
// g.AddEdge("m", "k")
// g.AddEdge("j", "m")

// r.AddEdge("b", "a")
// r.AddEdge("d", "a")
// r.AddEdge("d", "i")
// r.AddEdge("j", "i")
// r.AddEdge("b", "e")
// r.AddEdge("b", "c")
// r.AddEdge("f", "c")
// r.AddEdge("e", "h")
// r.AddEdge("e", "g")
// r.AddEdge("h", "l")
// r.AddEdge("n", "l")
// r.AddEdge("n", "m")
// r.AddEdge("g", "f")
// r.AddEdge("g", "j")

// r.AddEdge("g", "k")
// r.AddEdge("m", "k")
// r.AddEdge("j", "m")
