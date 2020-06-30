package graph

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

// SampleKRegularDirected From n elements, sample a random k-connected graph, and return its edges
// See https://www.rdocumentation.org/packages/igraph/versions/1.2.5/topics/sample_k_regular
// See also: https://github.com/igraph/igraph/blob/c2ccebab0be6bac56d9787cbfa4130eec64b76b2/src/games.c
func SampleKRegularDirected(n, k int) []DirectedEdge {
	if k > n-1 {
		panic("k cannot be larger than n - 1")
	}

	if n == 2 && k == 1 {
		return []DirectedEdge{
			{
				From: 0,
				To:   1,
			},
			{
				From: 1,
				To:   0,
			},
		}
	}
	if n == 3 && k == 2 {
		return []DirectedEdge{
			{
				From: 0,
				To:   1,
			},
			{
				From: 1,
				To:   0,
			},
			{
				From: 0,
				To:   2,
			},
			{
				From: 2,
				To:   0,
			},
			{
				From: 1,
				To:   2,
			},
			{
				From: 2,
				To:   1,
			},
		}
	}
	return nil
}

// SampleKRegularUndirected same as SampleKRegularDirected, but returns undirected edges
func SampleKRegularUndirected(n, k int) []UndirectedEdge {
	return nil
}
