package graph

import (
	"testing"
)

// Code for generating random edges on a graph

func TestSampleKRegular(t *testing.T) {
	testCases := []struct {
		name string
		n    int
		k    int
	}{
		{
			name: "Two node, 1-graph",
			n:    2,
			k:    1,
		},
		{
			name: "Three node, 2-graph",
			n:    3,
			k:    2,
		},
		{
			name: "10 node, 2-graph",
			n:    10,
			k:    2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			edges := SampleKRegularUndirected(tc.n, tc.k)
			if len(edges) != tc.n {
				t.Errorf("Expected %d total edges, got %d", tc.n, len(edges))
			}
			entryInFrom := make(map[int]int)
			totalEdges := 0
			for left := 0; left < len(edges); left++ {
				//if len(edges[left]) != tc.k {
				//	t.Errorf("Expected %d instances of entry %d as a 'from' in an edge, found %d", tc.k, left, len(edges[left]))
				//}
				for _, right := range edges[left] {
					totalEdges++
					entryInFrom[right]++
				}
			}
			// How many times an entry appears in the lefthand/righthand side of a edge

			//for entry, freq := range entryInFrom {
			//	if freq != tc.k {
			//		t.Errorf("Expected %d instances of entry %d as a 'from' in an edge, found %d", tc.k, entry, freq)
			//	}
			//}

			if totalEdges != tc.n*tc.k {
				t.Errorf("Expected %d edges, got %d", tc.n*tc.k, len(edges))
			}
		})
	}
}
