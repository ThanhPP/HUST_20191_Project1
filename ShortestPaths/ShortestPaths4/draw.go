package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func colorXYBlue(x int, y int, img *image.RGBA) {
	xEnd := (x) * scale
	yEnd := (y) * scale
	xStart := (x+1)*scale - 1
	yStart := (y+1)*scale - 1
	for i := xStart; i > xEnd; i-- {
		for j := yStart; j > yEnd; j-- {
			img.Set(i, j, color.RGBA{0, 0, 204, 0xff})
		}
	}
}

func colorXYRed(x int, y int, img *image.RGBA) {
	xEnd := (x) * scale
	yEnd := (y) * scale
	xStart := (x+1)*scale - 1
	yStart := (y+1)*scale - 1
	for i := xStart; i > xEnd; i-- {
		for j := yStart; j > yEnd; j-- {
			img.Set(i, j, color.RGBA{255, 0, 0, 0xff})
		}
	}
}

func colorXYLightGreen(x int, y int, img *image.RGBA) {
	xEnd := (x) * scale
	yEnd := (y) * scale
	xStart := (x+1)*scale - 1
	yStart := (y+1)*scale - 1
	for i := xStart; i > xEnd; i-- {
		for j := yStart; j > yEnd; j-- {
			img.Set(i, j, color.RGBA{102, 255, 102, 0xff})
		}
	}
}

func colorXYGreen(x int, y int, img *image.RGBA) {
	xEnd := (x) * scale
	yEnd := (y) * scale
	xStart := (x+1)*scale - 1
	yStart := (y+1)*scale - 1
	for i := xStart; i > xEnd; i-- {
		for j := yStart; j > yEnd; j-- {
			img.Set(i, j, color.RGBA{0, 204, 0, 0xff})
		}
	}
}

func makePNGMap(visited [size][size]bool, sPath []nodeTemplate, startNode nodeTemplate, desNode nodeTemplate, filePath string) {
	width := size*scale + 1
	height := size*scale + 1

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, color.White)
			if x%(scale) == 0 || y%scale == 0 || x == width-1 || y == height-1 {
				img.Set(x, y, color.Black)
			}
		}
	}

	// Color visited = blue
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if visited[i][j] {
				colorXYBlue(i, j, img)
			}
		}
	}

	// Color path = green
	for _, i := range sPath {
		colorXYLightGreen(i.x, i.y, img)
	}
	colorXYGreen(startNode.x, startNode.y, img)
	colorXYGreen(desNode.x, desNode.y, img)

	// Color block = red
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if map2d[i][j] == 1 {
				colorXYRed(i, j, img)
			}
		}
	}

	// Encode as PNG.
	f, _ := os.Create(filePath)
	png.Encode(f, img)
}
