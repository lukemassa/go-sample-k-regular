package model

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
			edges := SampleKRegular(tc.n, tc.k)
			if len(edges) != tc.n*tc.k {
				t.Errorf("Expected %d edges, got %d", tc.n*tc.k, len(edges))
			}
			// How many times an entry appears in the lefthand/righthand side of a edge
			entryInLeft := make(map[int]int)
			entryInRight := make(map[int]int)
			for _, edge := range edges {
				entryInLeft[edge.Left]++
				entryInRight[edge.Right]++
			}
			for entry, freq := range entryInLeft {
				if freq != tc.k {
					t.Errorf("Expected %d instances of entry %d on the left, found %d", tc.k, entry, freq)
				}
			}

			for entry, freq := range entryInRight {
				if freq != tc.k {
					t.Errorf("Expected %d instances of entry %d on the right, found %d", tc.k, entry, freq)
				}
			}
		})
	}
}
