package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"git.01.alem.school/MangoMango/lem-in/tools/algorithms"
)

type Lem struct {
	StartRoom string
	EndRoom   string
	Rooms     []string
	Edges     []string
	Ants      int
}

func main() {
	lem := Lem{}
	args := os.Args[1:]

	file, err := os.Open("examples/" + args[0])
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
		if scanner.Text()[:1] == "#" {
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

	g := &algorithms.Graph{}
	r := &algorithms.Graph{}

	g.AddVertex(lem.StartRoom)
	r.AddVertex(lem.StartRoom)
	g.AddVertex(lem.EndRoom)
	r.AddVertex(lem.EndRoom)

	for _, v := range lem.Rooms {
		g.AddVertex(v)
		r.AddVertex(v)
	}

	for _, v := range lem.Edges {
		temp := strings.Split(v, "-")
		g.AddEdge(temp[0], temp[1])
		r.AddEdge(temp[0], temp[1])
	}

	g.PrintGraph()

	foundPaths := g.FindAvailablePaths(r, lem.StartRoom, lem.EndRoom)

	if len(foundPaths) == 0 {
		fmt.Println("Paths not found")
	} else {
		fmt.Println(len(foundPaths))
		for _, v := range foundPaths {
			algorithms.PrintPath(v)
		}
	}

}

// g := &path.Graph{}
// 	for i := 'a'; i <= 'n'; i++ {
// 		g.AddVertex(string(i))
// 	}
// 	g.AddEdge("b", "a")
// 	g.AddEdge("d", "a")
// 	g.AddEdge("d", "i")
// 	g.AddEdge("j", "i")
// 	g.AddEdge("b", "e")
// 	g.AddEdge("b", "c")
// 	g.AddEdge("f", "c")
// 	g.AddEdge("e", "h")
// 	g.AddEdge("e", "g")
// 	g.AddEdge("h", "l")
// 	g.AddEdge("n", "l")
// 	g.AddEdge("n", "m")
// 	g.AddEdge("g", "f")
// 	g.AddEdge("g", "j")
// 	g.AddEdge("g", "k")
// 	g.AddEdge("m", "k")
// 	g.AddEdge("j", "m")

// 	r := &path.Graph{}
// 	for i := 'a'; i <= 'n'; i++ {
// 		r.AddVertex(string(i))
// 	}
// 	r.AddEdge("b", "a")
// 	r.AddEdge("d", "a")
// 	r.AddEdge("d", "i")
// 	r.AddEdge("j", "i")
// 	r.AddEdge("b", "e")
// 	r.AddEdge("b", "c")
// 	r.AddEdge("f", "c")
// 	r.AddEdge("e", "h")
// 	r.AddEdge("e", "g")
// 	r.AddEdge("h", "l")
// 	r.AddEdge("n", "l")
// 	r.AddEdge("n", "m")
// 	r.AddEdge("g", "f")
// 	r.AddEdge("g", "j")
// 	r.AddEdge("g", "k")
// 	r.AddEdge("m", "k")
// 	r.AddEdge("j", "m")
