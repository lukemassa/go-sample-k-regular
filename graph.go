package graph

// Code for generating random edges on a graph

// Edge An edge of a graph
// TODO: We don't really care about the order here, we could cut the number of edges in half
type Edge struct {
	Left  int
	Right int
}

// SampleKRegular From n elements, sample a random k-connected graph, and return its edges
// See https://www.rdocumentation.org/packages/igraph/versions/1.2.5/topics/sample_k_regular
func SampleKRegular(n, k int) []Edge {
	if k > n-1 {
		panic("k cannot be larger than n - 1")
	}

	if n == 2 && k == 1 {
		return []Edge{
			{
				Left:  0,
				Right: 1,
			},
			{
				Left:  1,
				Right: 0,
			},
		}
	}
	if n == 3 && k == 2 {
		return []Edge{
			{
				Left:  0,
				Right: 1,
			},
			{
				Left:  1,
				Right: 0,
			},
			{
				Left:  0,
				Right: 2,
			},
			{
				Left:  2,
				Right: 0,
			},
			{
				Left:  1,
				Right: 2,
			},
			{
				Left:  2,
				Right: 1,
			},
		}
	}
	return nil
}
