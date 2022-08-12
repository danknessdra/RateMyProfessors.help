package main

import (
	"github.com/RyanCarrier/dijkstra"
)

func Graph() *dijkstra.Graph {
	graph := dijkstra.NewGraph()
	//Add the 3 verticies
	graph.AddVertex(0)
	graph.AddVertex(1)
	graph.AddVertex(2)
	//Add the arcs
	graph.AddArc(0, 1, 1)
	graph.AddArc(0, 2, 1)
	graph.AddArc(1, 2, 2)
	return graph
}
