package main

import (
	"fmt"
	"os"
	"strconv"

	"git.01.alem.school/MangoMango/lem-in/tools/algorithms"
	"git.01.alem.school/MangoMango/lem-in/tools/structs"
)

func main() {

	args := os.Args[1:]
	rooms, err := structs.FileRead(args[0])
	if err != nil {
		fmt.Printf("ERROR:%s\n", err)
		return
	}
	mainGraph, copiedGraph := algorithms.SetGraphs(rooms)

	foundPaths, err := mainGraph.FindAvailablePaths(copiedGraph, rooms.StartRoom, rooms.EndRoom, rooms.Ants)
	if err != nil {
		fmt.Printf("ERROR:%s\n", err)
		return
	}
	if len(foundPaths) == 0 {
		fmt.Println("Paths not found")
		return
	}
	for _, v := range foundPaths {
		algorithms.PrintPath(v)
	}

	antsOnEachPath := make([]int, len(foundPaths))
	L := make([][]string, len(foundPaths))
	antsOnEachPath[0]++
	antCounter := 1
	L[0] = append(L[0], "L"+strconv.Itoa(antCounter))
	antCounter++
	rooms.Ants--
	if len(foundPaths) > 1 {
		for i := 0; rooms.Ants > 0; {
			if i+1 >= len(foundPaths) {
				i = 0
			}
			a := fmt.Sprintf("L%v", antCounter)
			if len(foundPaths[i])+antsOnEachPath[i] == len(foundPaths[i+1])+antsOnEachPath[i+1] {
				antsOnEachPath[i]++
				L[i] = append(L[i], a)
				antCounter++
				rooms.Ants--
				continue
			} else if len(foundPaths[i])+antsOnEachPath[i] < len(foundPaths[i+1])+antsOnEachPath[i+1] {
				antsOnEachPath[i]++
				L[i] = append(L[i], a)
				antCounter++
				rooms.Ants--
				i = 0
				continue
			}
			antsOnEachPath[i+1]++
			L[i+1] = append(L[i+1], a)
			antCounter++
			rooms.Ants--
			i++
		}
	} else {
		antsOnEachPath[0] += rooms.Ants
		for i := 1; i < antsOnEachPath[0]; i++ {
			L[0] = append(L[0], "L"+strconv.Itoa(antCounter))
			antCounter++
		}
	}
	maxLen := len(L[0])
	for _, v := range L {
		if len(v) > maxLen {
			maxLen = len(v)
		}
	}

	res := make([][]string, 1)

	element := 0
	stack := 0
	for idx := 0; idx < len(L); idx++ {
		for resIndex, vertex := range foundPaths[idx] {
			if element >= len(L[idx]) {
				break
			}
			if resIndex+stack >= len(res) {
				res = append(res, []string{})
			}
			res[resIndex+stack] = append(res[resIndex+stack], L[idx][element]+"-"+vertex.GetKey())
		}
		if idx+1 >= len(L) {
			idx = -1
			element++
			stack++
		}
		if element >= maxLen {
			break
		}
	}
	for _, stack := range res {
		for _, ant := range stack {
			fmt.Printf("%s ", ant)
		}
		fmt.Println()
	}
}

/*
L1-gilfoyle L2-dinish
L1-peter L2-jimYoung L3-gilfoyle L4-dinish
L2-peter L3-peter L4-jimYoung L5-gilfoyle L6-dinish
L4-peter L5-peter L6-jimYoung L7-gilfoyle L8-dinish
L6-peter L7-peter L8-jimYoung L9-gilfoyle
L8-peter L9-peter L10-gilfoyle
L10-peter
*/
