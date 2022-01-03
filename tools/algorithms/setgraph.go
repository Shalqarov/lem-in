package algorithms

import (
	"strings"

	"git.01.alem.school/MangoMango/lem-in/tools/structs"
)

func SetGraphs(rooms *structs.RoomsAndAnts) (*Graph, *Graph) {
	mainGraph := GraphInit()
	copiedGraph := GraphInit()

	mainGraph.AppendVertex(rooms.StartRoom)
	mainGraph.AppendVertex(rooms.EndRoom)

	copiedGraph.AppendVertex(rooms.StartRoom)
	copiedGraph.AppendVertex(rooms.EndRoom)

	mainGraph.SetStart(rooms.StartRoom)
	mainGraph.SetEnd(rooms.EndRoom)

	copiedGraph.SetStart(rooms.StartRoom)
	copiedGraph.SetEnd(rooms.EndRoom)

	for _, v := range rooms.Rooms {
		err := mainGraph.AppendVertex(v)
		if err != nil {
			continue
		}
		copiedGraph.AppendVertex(v)
	}

	for _, v := range rooms.Edges {
		temp := strings.Split(v, "-")
		mainGraph.AddEdge(mainGraph.GetVertex(temp[0]), mainGraph.GetVertex(temp[1]))
		copiedGraph.AddEdge(copiedGraph.GetVertex(temp[0]), copiedGraph.GetVertex(temp[1]))
	}

	return mainGraph, copiedGraph
}
