package main

import (
	"container/list"
	"fmt"
)

//graph

//Node an element of a graph
type Node struct {
	name string
}

type graph struct {
	nodes []*Node
	edges map[Node][]*Node
}

func (g *graph) addNode(n *Node) {
	g.nodes = append(g.nodes, n)
}

func (g *graph) addEdge(n1, n2 *Node) {
	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}
	g.edges[*n1] = append(g.edges[*n1], n2)
	g.edges[*n2] = append(g.edges[*n2], n1)
}

func (g *graph) show() {
	for i := 0; i < len(g.nodes); i++ {
		fmt.Printf("%v ->", g.nodes[i].name)
		near := g.edges[*g.nodes[i]]
		for j := 0; j < len(near); j++ {
			fmt.Printf(" %v", near[j].name)
		}
		fmt.Printf("\n")
	}
}

//end graph

//start BFS
func bfs(g *graph, n *Node, des *Node) []string {
	bfsRoute := make([]string, 0)
	// fmt.Println()
	// fmt.Println("===START BFS===")
	visited := make(map[string]bool)
	bfsQueue := list.New()
	bfsQueue.PushBack(n)
	visited[n.name] = true

	if n.name == des.name {
		return bfsRoute
	}

	for bfsQueue.Len() > 0 {
		qNode := bfsQueue.Front()
		// fmt.Println(qNode.Value.(*Node).name)
		visited[qNode.Value.(*Node).name] = true
		bfsRoute = append(bfsRoute, qNode.Value.(*Node).name)
		near := g.edges[*qNode.Value.(*Node)]
		// fmt.Println(near)

		for i := 0; i < len(near); i++ {
			j := near[i]
			if visited[j.name] == false {
				bfsQueue.PushBack(j)
				visited[j.name] = true
			}
		}
		bfsQueue.Remove(qNode)
	}
	return (bfsRoute)
	// fmt.Println("route : ", routes)
	// fmt.Println("===END BFS===")
	// fmt.Println()
}

//end BFS

//start DFS
func dfs(g *graph, n *Node, des *Node) (dfsRoute []string) {
	visited := make(map[string]bool)
	dfsQueue := list.New()
	dfsQueue.PushFront(n)
	visited[n.name] = true
	if n.name == des.name {
		return dfsRoute
	}

	for dfsQueue.Len() > 0 {
		qNode := dfsQueue.Front()
		visited[qNode.Value.(*Node).name] = true
		dfsRoute = append(dfsRoute, qNode.Value.(*Node).name)
		if qNode.Value.(*Node).name == des.name {
			return dfsRoute
		}
		near := g.edges[*qNode.Value.(*Node)]
		for i := len(near) - 1; i >= 0; i-- {
			j := near[i]
			if visited[j.name] == false {
				dfsQueue.PushFront(j)
				visited[j.name] = true
			}
		}
		dfsQueue.Remove(qNode)
	}
	return dfsRoute
}

//end DFS

func main() {
	var g graph

	nA := &Node{name: "A"}
	nB := &Node{name: "B"}
	nC := &Node{name: "C"}
	nD := &Node{name: "D"}
	nE := &Node{name: "E"}
	nF := &Node{name: "F"}

	g.addNode(nA)
	g.addNode(nB)
	g.addNode(nC)
	g.addNode(nD)
	g.addNode(nE)
	g.addNode(nF)

	g.addEdge(nA, nB)
	g.addEdge(nB, nC)
	g.addEdge(nB, nE)
	g.addEdge(nC, nD)
	g.addEdge(nE, nD)
	g.addEdge(nD, nF)

	g.show()

	fmt.Println()
	fmt.Println("===START BFS===")
	BFSroute := bfs(&g, nA, nF)
	fmt.Println("BFS route = ", BFSroute)
	fmt.Println("===END BFS===")
	fmt.Println()

	fmt.Println()
	fmt.Println("===START DFS===")
	DFSroute := dfs(&g, nA, nF)
	fmt.Println("DFS route = ", DFSroute)
	fmt.Println("===END DFS===")
	fmt.Println()
}
