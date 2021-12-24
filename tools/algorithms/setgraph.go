package algorithms

import (
	"strings"

	"git.01.alem.school/MangoMango/lem-in/tools/structs"
)

func SetGraphs(rooms *structs.RoomsAndAnts) (*Graph, *Graph) {
	mainGraph := GraphInit()
	copiedGraph := GraphInit()

	mainGraph.AppendVertex(rooms.StartRoom)
	copiedGraph.AppendVertex(rooms.StartRoom)
	mainGraph.AppendVertex(rooms.EndRoom)
	copiedGraph.AppendVertex(rooms.EndRoom)

	for _, v := range rooms.Rooms {
		mainGraph.AppendVertex(v)
		copiedGraph.AppendVertex(v)
	}

	for _, v := range rooms.Edges {
		temp := strings.Split(v, "-")
		mainGraph.AddEdge(temp[0], temp[1], false)
		copiedGraph.AddEdge(temp[0], temp[1], false)
	}

	return mainGraph, copiedGraph
}
