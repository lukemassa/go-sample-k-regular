package graph

import (
	"fmt"
	"math/rand"

	"github.com/soniakeys/graph"
)

var (
	SampleKRegularUndirected = sampleKRegularUndirectedNew
)

// Code for generating random edges on a graph

// DirectedEdge An edge of a directed graph
type DirectedEdge struct {
	From int
	To   int
}

// UndirectedEdge An edge of an undirected graph
type UndirectedEdge struct {
	A int
	B int
}

// SampleKRegularDirectedNew better
func sampleKRegularUndirectedNew(n, k int) [][]int {
	res := graph.Gnm3Undirected(n, (n*k)/2, nil)
	edges := make([][]int, n)
	for i := 0; i < len(res.AdjacencyList); i++ {
		edges[i] = make([]int, len(res.AdjacencyList[i]))
		for j := 0; j < len(res.AdjacencyList[i]); j++ {
			edges[i][j] = int(res.AdjacencyList[i][j])
		}
	}
	return edges
}

// SampleKRegularDirectedOriginal From n elements, sample a random k-connected graph, and return its edges
// See https://www.rdocumentation.org/packages/igraph/versions/1.2.5/topics/sample_k_regular
// See also: https://github.com/igraph/igraph/blob/c2ccebab0be6bac56d9787cbfa4130eec64b76b2/src/games.c
func sampleKRegularDirectedOriginal(n, k int) [][]int {
	if k > n-1 {
		panic("k cannot be larger than n - 1")
	}
	LeftStubs := make([]int, n*k)
	RightStubs := make([]int, n*k)
	for i := 0; i < n; i++ {
		for j := 0; j < k; j++ {
			LeftStubs[i*k+j] = i
			RightStubs[i*k+j] = i
		}
	}

	for true {

		edges := make([][]int, n)

		rand.Shuffle(n*k, func(i, j int) {
			RightStubs[i], RightStubs[j] = RightStubs[j], RightStubs[i]
		})

		found := true
		for i := 0; i < n*k; i++ {

			if LeftStubs[i] == RightStubs[i] {
				found = false
				break
			}
			if edges[LeftStubs[i]] == nil {
				edges[LeftStubs[i]] = make([]int, 0)
			}
			edges[LeftStubs[i]] = append(edges[LeftStubs[i]], RightStubs[i])
		}
		if found {
			return edges
		}
	}
	panic("Unreachable")
}
