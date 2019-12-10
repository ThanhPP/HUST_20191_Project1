package main

import (
	"container/list"
	"fmt"
)

// import (
// 	"fmt"
// 	"math/rand"
// )

// //start Map function
// const size int = 5

// var m [size][size]int

// func genMap() {
// 	for i := 0; i < size; i++ {
// 		for j := 0; j < size; j++ {
// 			if i >= j {
// 				m[i][j] = 0
// 				continue
// 			}
// 			m[i][j] = rand.Intn(5)
// 		}
// 	}
// }

// func showMap() {
// 	fmt.Println()
// 	fmt.Println("------------------SHOW MAP------------------")
// 	fmt.Printf("    ")
// 	for k := 65; k < 65+size; k++ {
// 		fmt.Printf(" %-3c", k)
// 	}
// 	fmt.Println()
// 	k := 65
// 	for i := 0; i < size; i++ {
// 		fmt.Printf(" %-3c", k)
// 		k++
// 		for j := 0; j < size; j++ {
// 			fmt.Printf(" %-3d", m[i][j])
// 		}
// 		fmt.Println()
// 	}
// 	fmt.Println("--------------------------------------------")
// 	fmt.Println()
// }

// //end map function

type node struct {
	name    string
	friends map[string]*node
}

func bfs(n *node) []string {
	//keep track of visited nodes
	visited := make(map[string]*node)
	nodes := make([]string, 0)

	//queue of nodes need to visit
	queue := list.New()
	queue.PushBack(n)
	visited[n.name] = n
	nodes = append(nodes, n.name)
	for queue.Len() > 0 {
		qnode := queue.Front()
		for id, node := range qnode.Value.(*node).friends {
			if _, ok := visited[id]; !ok {
				visited[id] = node
				nodes = append(nodes, node.name)
				queue.PushBack(node)
			}
		}
		queue.Remove(qnode)
	}

	// for _, node := range visited {
	// 	nodes = append(nodes, node.name)
	// }
	return nodes
}

func main() {
	// genMap()
	// showMap()
	node1 := &node{
		name:    "A",
		friends: make(map[string]*node),
	}
	node2 := &node{
		name:    "B",
		friends: make(map[string]*node),
	}
	node3 := &node{
		name:    "C",
		friends: make(map[string]*node),
	}
	node4 := &node{
		name:    "D",
		friends: make(map[string]*node),
	}
	node5 := &node{
		name:    "E",
		friends: make(map[string]*node),
	}
	node1.friends[node2.name] = node2
	node1.friends[node4.name] = node4
	node2.friends[node3.name] = node3
	node2.friends[node5.name] = node5
	nodes := bfs(node1)
	for _, node := range nodes {
		fmt.Println(node)
	}
}
