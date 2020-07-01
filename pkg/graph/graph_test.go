package graph

import "testing"

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
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			edges := SampleKRegularDirected(tc.n, tc.k)
			if len(edges) != tc.n*tc.k {
				t.Errorf("Expected %d edges, got %d", tc.n*tc.k, len(edges))
			}
			// How many times an entry appears in the lefthand/righthand side of a edge
			entryInFrom := make(map[int]int)
			entryInTo := make(map[int]int)
			for _, edge := range edges {
				entryInFrom[edge.To]++
				entryInTo[edge.From]++
				if edge.To == edge.From {
					t.Errorf("Graph contains a loop at %d", edge.To)
				}
			}
			for entry, freq := range entryInFrom {
				if freq != tc.k {
					t.Errorf("Expected %d instances of entry %d as a 'from' in an edge, found %d", tc.k, entry, freq)
				}
			}

			for entry, freq := range entryInTo {
				if freq != tc.k {
					t.Errorf("Expected %d instances of entry %d as a 'to' in an edge, found %d", tc.k, entry, freq)
				}
			}
		})
	}
}
