package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/lukemassa/go-sample-k-regular/pkg/graph"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	//edges := graph.SampleKRegularDirected(300, 10)
	start := time.Now()
	edges := graph.SampleKRegularUndirected(100, 10)
	fmt.Printf("Took %s\n", time.Since(start))

	for i := 0; i < len(edges); i++ {
		fmt.Printf("Edge: %v\n", edges[i])
	}
	fmt.Printf("Took %s\n", time.Since(start))
}
