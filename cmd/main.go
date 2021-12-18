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
		if scanner.Text() == "" {
			continue
		}
		if ants, err := strconv.Atoi(scanner.Text()); err == nil {
			lem.Ants = ants
			continue
		}
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
		temp := strings.Split(scanner.Text(), " ")
		lem.Rooms = append(lem.Rooms, temp[0])
		fmt.Println(scanner.Text())
	}
	fmt.Println()
	g := algorithms.GraphInit()
	r := algorithms.GraphInit()

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
		g.AddEdge(temp[0], temp[1], false)
		r.AddEdge(temp[0], temp[1], false)
	}

	foundPaths := g.FindAvailablePaths(r, lem.StartRoom, lem.EndRoom, lem.Ants)
	fmt.Println("############################################")
	if len(foundPaths) == 0 {
		fmt.Println("Paths not found")
		return
	}
	for _, v := range foundPaths {
		algorithms.PrintPath(v)
	}

	fmt.Println(lem.Ants)

	antsOnEachPath := make([]int, len(foundPaths))
	antsOnEachPath[0]++
	lem.Ants--
	if len(foundPaths) > 1 {
		for i := 0; lem.Ants > 0; {
			if i+1 >= len(foundPaths) {
				i = 0
			}
			// Если кол-во комнат + кол-во муравьев на этом пути меньше или равно
			// кол-вам комнат + кол-вам муравьев следующего пути
			// то засчитываю муравья на нынешнем пути...
			if len(foundPaths[i])+antsOnEachPath[i] <= len(foundPaths[i+1])+antsOnEachPath[i+1] {
				antsOnEachPath[i]++
				lem.Ants--
				continue
			}
			// ...иначе добавляю муравья на след путь
			antsOnEachPath[i+1]++
			lem.Ants--
			i++
		}
	} else { // если путь только один
		antsOnEachPath[0] = lem.Ants + 1 // + 1 муравей, потому что по умолчанию одного муравья ставлю в первый путь
	}
	for i, v := range antsOnEachPath {
		fmt.Println(i+1, "Path:", v, "ants")
	}

	// res := [][]string{}

}
