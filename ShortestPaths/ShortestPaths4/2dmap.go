package main

import (
	"fmt"
	"math/rand"
)

var map2d [size][size]int

func showMap() {
	fmt.Println()
	fmt.Println("---- Show map ----")
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Printf(" %-4d", map2d[j][i])
		}
		fmt.Println()
	}
}

func genMap() {
	for i := 0; i < size; i++ {
		for j := 0; j <
			size; j++ {
			temp := rand.Intn(10)
			if temp > 1 {
				map2d[i][j] = 0
				continue
			}
			map2d[i][j] = 1
		}
	}
	showMap()
}
