package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

type Claim struct {
	x, y, w, h int
}

func NewClaim(line string) Claim {
	parts := strings.Split(line, " ")

	xy := strings.Split(parts[2], ",")
	x, _ := strconv.Atoi(xy[0])
	y, _ := strconv.Atoi(strings.TrimSuffix(xy[1], ":"))

	wh := strings.Split(parts[3], "x")
	w, _ := strconv.Atoi(wh[0])
	h, _ := strconv.Atoi(wh[1])

	return Claim{x, y, w, h}
}

func die(v ...interface{}) {
	fmt.Println(v...)
	os.Exit(1)
}

func main() {
	if len(os.Args) != 2 {
		die("Please specify input text file as only argument")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		die(err)
	}
	defer file.Close()

	fabric := [1000][1000]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		c := NewClaim(scanner.Text())

		for i := 0; i < c.w; i++ {
			for j := 0; j < c.h; j++ {
				fabric[i + c.x][j + c.y]++
			}
		}
	}

	area := 0

	for _, rows := range fabric {
		for _, unit := range rows {
			if unit > 1 {
				area++
			}
		}
	}

	fmt.Println("Overlapping area:", area)
}
