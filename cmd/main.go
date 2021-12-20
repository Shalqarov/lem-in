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

	fmt.Println("Ants:", lem.Ants)

	antsOnEachPath := make([]int, len(foundPaths))
	L := make([][]string, len(foundPaths))
	antsOnEachPath[0]++
	cnt := 1
	L[0] = append(L[0], "L"+strconv.Itoa(cnt))
	cnt++
	lem.Ants--
	if len(foundPaths) > 1 {

		for i := 0; lem.Ants > 0; {
			if i+1 >= len(foundPaths) {
				i = 0
			}
			a := fmt.Sprintf("L%v", cnt)
			// Если кол-во комнат + кол-во муравьев на этом пути меньше или равно
			// кол-вам комнат + кол-вам муравьев следующего пути
			// то засчитываю муравья на нынешнем пути...
			if len(foundPaths[i])+antsOnEachPath[i] == len(foundPaths[i+1])+antsOnEachPath[i+1] {
				antsOnEachPath[i]++
				L[i] = append(L[i], a)
				cnt++
				lem.Ants--
				continue
			} else if len(foundPaths[i])+antsOnEachPath[i] < len(foundPaths[i+1])+antsOnEachPath[i+1] {
				antsOnEachPath[i]++
				L[i] = append(L[i], a)
				cnt++
				lem.Ants--
				i = 0
				continue
			}
			// ...иначе добавляю муравья на след путь
			antsOnEachPath[i+1]++
			L[i+1] = append(L[i+1], a)
			cnt++
			lem.Ants--
			i++
		}
	} else { // если путь только один
		antsOnEachPath[0] += lem.Ants
		for i := 1; i < antsOnEachPath[0]; i++ {
			L[0] = append(L[0], "L"+strconv.Itoa(cnt))
			cnt++
		}
	}
	for i, v := range antsOnEachPath {
		fmt.Println(i+1, "Path:", v, "ants")
	}
	// maxLen := len(L[0])
	// for _, v := range L {
	// 	fmt.Println(v)
	// 	if len(v) > maxLen {
	// 		maxLen = len(v)
	// 	}
	// }

	// res := [][]string{}
	// for _, v := range L {
	// 	for i := 0; i < maxLen; i++ {
	// 		if i < len(v) {

	// 		}
	// 	}
	// }

}
