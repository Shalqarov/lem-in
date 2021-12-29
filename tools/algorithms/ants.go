package algorithms

import (
	"fmt"
	"strconv"
)

func AntsOnEachPathCount(foundPaths [][]*Vertex, ants int) [][]string {
	antsOnPath := make([]int, len(foundPaths))
	L := make([][]string, len(foundPaths))
	antsOnPath[0]++
	antCounter := 1
	L[0] = append(L[0], "L"+strconv.Itoa(antCounter))
	antCounter++
	ants--
	if len(foundPaths) > 1 {
		for i := 0; ants > 0; {
			if i+1 >= len(foundPaths) {
				i = 0
			}
			antID := fmt.Sprintf("L%v", antCounter)
			if len(foundPaths[i])+antsOnPath[i] == len(foundPaths[i+1])+antsOnPath[i+1] {
				antsOnPath[i]++
				L[i] = append(L[i], antID)
				antCounter++
				ants--
				continue
			} else if len(foundPaths[i])+antsOnPath[i] < len(foundPaths[i+1])+antsOnPath[i+1] {
				antsOnPath[i]++
				L[i] = append(L[i], antID)
				antCounter++
				ants--
				i = 0
				continue
			}
			antsOnPath[i+1]++
			L[i+1] = append(L[i+1], antID)
			antCounter++
			ants--
			i++
		}
	} else {
		antsOnPath[0] += ants
		for i := 1; i < antsOnPath[0]; i++ {
			L[0] = append(L[0], "L"+strconv.Itoa(antCounter))
			antCounter++
		}
	}
	return L
}

func PrintAnts(antsOnPath [][]string, foundPaths [][]*Vertex) {
	if len(foundPaths[0]) == 1 {
		count := len(antsOnPath[0])
		for i := 1; count > 0; i++ {
			fmt.Printf("L%d ", i)
			count--
		}
		fmt.Println()
		return
	}
	maxLen := len(antsOnPath[0])
	for _, v := range antsOnPath {
		if len(v) > maxLen {
			maxLen = len(v)
		}
	}

	res := make([][]string, 1)
	for index, element, stack := 0, 0, 0; index < len(antsOnPath); index++ {
		for resIndex, vertex := range foundPaths[index] {
			if element >= len(antsOnPath[index]) {
				break
			}
			if resIndex+stack >= len(res) {
				res = append(res, []string{})
			}
			res[resIndex+stack] = append(res[resIndex+stack], antsOnPath[index][element]+"-"+vertex.GetKey())
		}
		if index+1 >= len(antsOnPath) {
			index = -1
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
