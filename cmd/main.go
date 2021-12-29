package main

import (
	"fmt"
	"os"

	"git.01.alem.school/MangoMango/lem-in/tools/algorithms"
	"git.01.alem.school/MangoMango/lem-in/tools/structs"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("ERROR: invalid number of arguments: please enter 'go run ./cmd/ filename.txt'")
		fmt.Println("all examples in 'examples' folder")
		return
	}
	rooms, err := structs.FileRead(args[0])
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return
	}
	mainGraph, copiedGraph := algorithms.SetGraphs(rooms)
	foundPaths, err := mainGraph.FindAvailablePaths(copiedGraph, rooms.Ants)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return
	}
	antsOnEachPath := algorithms.AntsOnEachPathCount(foundPaths, rooms.Ants)
	algorithms.PrintAnts(antsOnEachPath, foundPaths)
}
