package main

import (
	"fmt"
	"os"

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
	foundPaths, err := mainGraph.FindAvailablePaths(copiedGraph, rooms.Ants)
	if err != nil {
		fmt.Printf("ERROR:%s\n", err)
		return
	}
	antsOnEachPath := algorithms.AntsOnEachPathCount(foundPaths, rooms.Ants)
	algorithms.PrintAnts(antsOnEachPath, foundPaths)
}
