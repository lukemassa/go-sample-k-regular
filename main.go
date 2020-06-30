package main

import (
	"fmt"
    "github.com/lukemassa/go-sample-k-regular/pkg/graph"

)

func main() {
	edges := graph.SampleKRegularDirected(2, 1)
	for i := 0; i < len(edges); i++ {
		fmt.Printf("Edge: %v\n", edges[i])
	}
}
