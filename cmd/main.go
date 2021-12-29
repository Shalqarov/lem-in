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

/*
L1-gilfoyle L2-dinish
L1-peter L2-jimYoung L3-gilfoyle L4-dinish
L2-peter L3-peter L4-jimYoung L5-gilfoyle L6-dinish
L4-peter L5-peter L6-jimYoung L7-gilfoyle L8-dinish
L6-peter L7-peter L8-jimYoung L9-gilfoyle
L8-peter L9-peter L10-gilfoyle
L10-peter
*/
