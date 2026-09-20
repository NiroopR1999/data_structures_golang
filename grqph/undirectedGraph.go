package graph

import (
	"fmt"
	"slices"
)

type Graph struct {
	list map[string][]string
}

func NewGraph() *Graph {
	return &Graph{
		// We initialize the map because writing to a nil map causes a panic.
		list: make(map[string][]string),
	}
}

func (g *Graph) AddVertex(vertex string) {
	// A vertex should only be added if it doesn't already exist.
	// This prevents accidentally replacing its existing adjacency list.
	if _, exists := g.list[vertex]; !exists {
		g.list[vertex] = []string{}
	}
}

func (g *Graph) AddEdge(vertex1, vertex2 string) {
	// Make sure both vertices exist before creating the edge.
	g.AddVertex(vertex1)
	g.AddVertex(vertex2)

	// We check first because the same edge should not be added twice.
	if !g.HasEdge(vertex1, vertex2) {

		// Since this is an undirected graph, vertex2 is a neighbor of vertex1.
		g.list[vertex1] = append(g.list[vertex1], vertex2)

		// And vertex1 is a neighbor of vertex2.
		// We therefore store the edge in both adjacency lists.
		g.list[vertex2] = append(g.list[vertex2], vertex1)
	}
}

func (g *Graph) Display() {
	for vertex, neighbors := range g.list {
		fmt.Println(vertex, "->", neighbors)
	}
}

func (g *Graph) HasEdge(vertex1, vertex2 string) bool {
	// In an adjacency-list representation, an edge exists
	// if vertex2 appears in vertex1's list of neighbors.
	return slices.Contains(g.list[vertex1], vertex2)
}

func (g *Graph) RemoveEdge(vertex1, vertex2 string) {
	if !g.HasEdge(vertex1, vertex2) {
		return
	}

	// Find vertex2 inside vertex1's adjacency list.
	index1 := slices.Index(g.list[vertex1], vertex2)

	// Find vertex1 inside vertex2's adjacency list.
	// We need to remove both because this is an undirected edge.
	index2 := slices.Index(g.list[vertex2], vertex1)

	// Remove vertex2 from vertex1's neighbors.
	g.list[vertex1] = slices.Delete(
		g.list[vertex1],
		index1,
		index1+1,
	)

	// Remove vertex1 from vertex2's neighbors.
	g.list[vertex2] = slices.Delete(
		g.list[vertex2],
		index2,
		index2+1,
	)
}

func (g *Graph) RemoveVertex(vertex string) {
	if _, exists := g.list[vertex]; !exists {
		return
	}

	// Because this is an undirected graph, the vertex may appear
	// inside many other vertices' adjacency lists.
	//
	// We need to remove it from every neighbor's list.
	for key, neighbors := range g.list {
		index := slices.Index(neighbors, vertex)

		if index != -1 {
			g.list[key] = slices.Delete(
				neighbors,
				index,
				index+1,
			)
		}
	}

	// Finally remove the vertex itself from the graph.
	delete(g.list, vertex)
}