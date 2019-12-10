package main

import (
	"fmt"
	"time"
)

func main() {
	genMap()
	buildGraph()
	// showGraph()
	startNode := nodeTemplate{
		x: 1,
		y: 0,
	}

	desNode := nodeTemplate{
		x: 0,
		y: 19,
	}

	if map2d[startNode.x][startNode.y] == 1 || map2d[desNode.x][desNode.y] == 1 {
		fmt.Println("Blocked node")
		return
	}

	fmt.Println("START DIJIKSTRA")
	dijikstraTimer := time.Now()
	cost, sPath, visited := dijikstra(startNode.x, startNode.y, desNode.x, desNode.y)
	dijikstraTime := time.Since(dijikstraTimer)
	fmt.Println("dijikstra time = ", dijikstraTime)
	fmt.Println("dijikstra cost = ", cost)
	fmt.Println("dijikstra path : ", sPath)
	count := 0
	fmt.Print("visited :")
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if visited[i][j] {
				count++
				fmt.Printf(" {%d, %d}", i, j)
			}
		}
	}
	fmt.Println()
	fmt.Println("visited nodes = ", count)
	makePNGMap(visited, sPath, startNode, desNode, "img/dijikstra.png")
	fmt.Println()

	fmt.Println("START A-STAR")
	aStarTimer := time.Now()
	costA, sPathA, visitedA := aStar(startNode.x, startNode.y, desNode.x, desNode.y)
	aStarTime := time.Since(aStarTimer)
	fmt.Println("a-Star time = ", aStarTime)
	fmt.Println("a-star cost = ", costA)
	fmt.Println("a-star path : ", sPathA)
	countA := 0
	fmt.Print("visited a-star :")
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if visitedA[i][j] {
				countA++
				fmt.Printf(" {%d, %d}", i, j)
			}
		}
	}
	fmt.Println()
	fmt.Println("visited nodes = ", countA)
	makePNGMap(visitedA, sPathA, startNode, desNode, "img/aStars.png")
}
