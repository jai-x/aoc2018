package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
	"strings"
	"math/rand"
	"image"
	"image/color"
	"image/png"
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

func randRGBA() color.RGBA {
	return color.RGBA{
		R: uint8(rand.Intn(256)),
		G: uint8(rand.Intn(256)),
		B: uint8(rand.Intn(256)),
		A: 255,
	}
}

func pointColors(points []image.Point) map[image.Point]color.RGBA {
	out := map[image.Point]color.RGBA{}
	for _, p := range points {
		out[p] = randRGBA()
	}
	return out
}

func saveImage(img image.Image) {
	name := "task1.png"

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

	colors := pointColors(points)

	xMax, yMax := xyMax(points)

	img := image.NewRGBA(image.Rect(0, 0, xMax, yMax))

	white := color.RGBA{255, 255, 255, 255}

	// Consider each point
	for x := 0; x < xMax; x++ {
		for y := 0; y < yMax; y++ {
			// Consider manhatten distance for each point against this location
			dists := map[image.Point]int{}
			for i := range points {
				dists[points[i]] = manhatten(points[i], image.Pt(x,y))
			}

			// Find the minimum manhatten distance from the list
			minDist := 100000000
			var minPoint image.Point
			for p, d := range dists {
				if d < minDist {
					minPoint, minDist = p, d
				}
			}

			//  Equidistant to points
			eqd := false
			for p, d := range dists {
				if d == minDist && !p.Eq(minPoint) {
					eqd = true
				}
			}

			// White if equidistant, point color if not
			if eqd {
				img.SetRGBA(x, y, white)
			} else {
				img.SetRGBA(x, y, colors[minPoint])
			}
		}
	}

	edgeColors := map[color.RGBA]bool{}
	for x := 0; x < xMax; x++ {
		for y := 0; y < yMax; y++ {
			col := img.RGBAAt(x, y)

			edge := (x == 0) || (y == 0) || (x == (xMax-1)) || (y == (yMax-1))

			if edge && (col != white) {
				edgeColors[col] = true
			}
		}
	}

	finiteSpace := map[color.RGBA]int{}
	for x := 0; x < xMax; x++ {
		for y := 0; y < yMax; y++ {
			col := img.RGBAAt(x, y)

			_, edge := edgeColors[col]

			if !edge && (col != white) {
				finiteSpace[col]++
			}
		}
	}

	maxSpace := 0
	for _, space := range finiteSpace {
		if space > maxSpace {
			maxSpace = space
		}
	}
	fmt.Println("Max finite space:", maxSpace)

	// Draw crosses for each input point, for fun :P
	for _, p := range points {
		img.Set(p.X  , p.Y  , color.Black)
		img.Set(p.X-1, p.Y-1 , color.Black)
		img.Set(p.X+1, p.Y-1  , color.Black)
		img.Set(p.X-1, p.Y+1, color.Black)
		img.Set(p.X+1, p.Y+1, color.Black)
	}

	saveImage(img)
}
