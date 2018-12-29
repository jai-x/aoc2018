package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
	"strings"
	"image"
	"image/color"
	"image/png"
	"math"
)

func die(v ...interface{}) {
	fmt.Println(v...)
	os.Exit(1)
}

func abs(i int) int {
	if i < 0 {
		i = -i
	}
	return i
}

func manhatten(p, q image.Point) int {
	return abs(p.X - q.X) + abs(p.Y - q.Y)
}

func parsePoints() []image.Point {
	if len(os.Args) != 2 {
		die("Please specify input text file as only argument")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		die(err)
	}
	defer file.Close()

	out := []image.Point{}

	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		parts := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		y, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		out = append(out, image.Pt(x, y))
	}

	return out
}

func xyMax(points []image.Point) (int, int) {
	xMax, yMax := 0, 0
	for i := range points {
		if points[i].X > xMax {
			xMax = points[i].X
		}

		if points[i].Y > yMax {
			yMax = points[i].Y
		}
	}
	return xMax + 2, yMax + 2
}

func initGrid(x, y int) [][]int {
	out := make([][]int, x)
	for i := range out {
		out[i] = make([]int, y)
	}
	return out
}

func gridToImg(x, y, min, max int, grid [][]int) {
	img := image.NewGray16(image.Rect(0, 0, x, y))

	for x := range grid {
		for y := range grid[x] {
			val := ((grid[x][y] - min) * math.MaxUint16) / max
			col := color.Gray16{uint16(val)}
			img.SetGray16(x, y, col)
		}
	}

	name := "task2.png"

	file, err := os.Create(name)
	if err != nil {
		die(err)
	}

	err = png.Encode(file, img)
	if err != nil {
		die(err)
	}

	fmt.Println("Saved visualisation of input as:", name)
}

func main() {
	points := parsePoints()

	xMax, yMax := xyMax(points)

	grid := initGrid(xMax, yMax)

	distMax := 0
	distMin := 9999999999999999

	safeZones := 0

	for x := range grid {
		for y := range grid[x] {
			dist := 0

			for i := range points {
				dist += manhatten(points[i], image.Pt(x, y))
			}

			if dist > distMax {
				distMax = dist
			}

			if dist < distMin {
				distMin = dist
			}

			if dist < 10000 {
				safeZones++
			}

			grid[x][y] = dist
		}
	}

	fmt.Println("Safe area:", safeZones)

	gridToImg(xMax, yMax, distMin, distMax, grid)
}
