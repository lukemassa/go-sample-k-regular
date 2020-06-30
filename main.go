package main

import (
	"fmt"

	"github.com/lukemassa/go-sample-k-regular/graph"
)

func main() {
	edges := graph.SampleKRegular(2, 1)
	for i := 0; i < len(edges); i++ {
		fmt.Println("Edge: %v", edges[i])
	}
}
