package main

import "fmt"

type nodeTemplate struct {
	x int
	y int
}

type edge struct {
	node   nodeTemplate
	weight int
}

var graph [size][size][]edge

func buildGraph() error {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			addEdge(i, j)
		}
	}
	return nil
}

func showGraph() error {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Println(graph[i][j])
		}
	}
	return nil
}

func addEdge(x int, y int) {
	if map2d[x][y] == 1 {
		return
	}
	realSize := size - 1
	if x == 0 && y == 0 {
		if map2d[x][y+1] == 0 {
			graph[x][y] = append(graph[x][y], edge{nodeTemplate{x, y + 1}, 1})
		}
		if map2d[x+1][y] == 0 {
			graph[x][y] = append(graph[x][y], edge{nodeTemplate{x + 1, y}, 1})
		}
		return
	}
	if x == realSize && y == 0 {
		if map2d[x][y+1] == 0 {
			graph[x][y] = append(graph[x][y], edge{nodeTemplate{x, y + 1}, 1})
		}
		if map2d[x-1][y] == 0 {
			graph[x][y] = append(graph[x][y], edge{nodeTemplate{x - 1, y}, 1})
		}
		return
	}
	if y == 0 {
		if map2d[x][y+1] == 0 {
			graph[x][y] = append(graph[x][y], edge{nodeTemplate{x, y + 1}, 1})
		}
		if map2d[x+1][y] == 0 {
			graph[x][y] = append(graph[x][y], edge{nodeTemplate{x + 1, y}, 1})
		}
		if map2d[x-1][y] == 0 {
			graph[x][y] = append(graph[x][y], edge{nodeTemplate{x - 1, y}, 1})
		}
		return
	}
	if x == 0 && y == realSize {
		if map2d[x+1][y] == 0 {
			graph[x][y] = append(graph[x][y], edge{nodeTemplate{x + 1, y}, 1})
		}
		if map2d[x][y-1] == 0 {
			graph[x][y] = append(graph[x][y], edge{nodeTemplate{x, y - 1}, 1})
		}
		return
	}
	if x == realSize && y == realSize {
		if map2d[x-1][y] == 0 {
			graph[x][y] = append(graph[x][y], edge{nodeTemplate{x - 1, y}, 1})
		}
		if map2d[x][y-1] == 0 {
			graph[x][y] = append(graph[x][y], edge{nodeTemplate{x, y - 1}, 1})
		}
		return
	}
	if y == realSize {
		if map2d[x+1][y] == 0 {
			graph[x][y] = append(graph[x][y], edge{nodeTemplate{x + 1, y}, 1})
		}
		if map2d[x-1][y] == 0 {
			graph[x][y] = append(graph[x][y], edge{nodeTemplate{x - 1, y}, 1})
		}
		if map2d[x][y-1] == 0 {
			graph[x][y] = append(graph[x][y], edge{nodeTemplate{x, y - 1}, 1})
		}
		return
	}
	if x == 0 {
		if map2d[x+1][y] == 0 {
			graph[x][y] = append(graph[x][y], edge{nodeTemplate{x + 1, y}, 1})
		}
		if map2d[x][y+1] == 0 {
			graph[x][y] = append(graph[x][y], edge{nodeTemplate{x, y + 1}, 1})
		}
		if map2d[x][y-1] == 0 {
			graph[x][y] = append(graph[x][y], edge{nodeTemplate{x, y - 1}, 1})
		}
		return
	}
	if x == realSize {
		if map2d[x-1][y] == 0 {
			graph[x][y] = append(graph[x][y], edge{nodeTemplate{x - 1, y}, 1})
		}
		if map2d[x][y+1] == 0 {
			graph[x][y] = append(graph[x][y], edge{nodeTemplate{x, y + 1}, 1})
		}
		if map2d[x][y-1] == 0 {
			graph[x][y] = append(graph[x][y], edge{nodeTemplate{x, y - 1}, 1})
		}
		return
	}
	if map2d[x+1][y] == 0 {
		graph[x][y] = append(graph[x][y], edge{nodeTemplate{x + 1, y}, 1})
	}
	if map2d[x-1][y] == 0 {
		graph[x][y] = append(graph[x][y], edge{nodeTemplate{x - 1, y}, 1})
	}
	if map2d[x][y+1] == 0 {
		graph[x][y] = append(graph[x][y], edge{nodeTemplate{x, y + 1}, 1})
	}
	if map2d[x][y-1] == 0 {
		graph[x][y] = append(graph[x][y], edge{nodeTemplate{x, y - 1}, 1})
	}
}

func getEdges(x, y int) []edge {
	return graph[x][y]
}
