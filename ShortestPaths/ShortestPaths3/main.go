package main

import (
	"fmt"
)

func main() {
	fmt.Println("Dijkstra")
	// Example
	g := newGraph()
	g.addEdge("A", "B", 4)
	g.addEdge("A", "C", 2)
	g.addEdge("B", "C", 1)
	g.addEdge("B", "D", 9)
	g.addEdge("C", "D", 8)
	g.addEdge("C", "E", 10)
	g.addEdge("D", "E", 2)
	g.addEdge("D", "F", 6)
	g.addEdge("E", "F", 2)
	fmt.Println(g.dijikstra("A", "F"))
}
