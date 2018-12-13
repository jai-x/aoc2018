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
	id string
}

func NewClaim(line string) Claim {
	parts := strings.Split(line, " ")

	xy := strings.Split(parts[2], ",")
	x, _ := strconv.Atoi(xy[0])
	y, _ := strconv.Atoi(strings.TrimSuffix(xy[1], ":"))

	wh := strings.Split(parts[3], "x")
	w, _ := strconv.Atoi(wh[0])
	h, _ := strconv.Atoi(wh[1])

	id := parts[0]

	return Claim{x, y, w, h, id}
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

	claims := []Claim{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		claims = append(claims, NewClaim(scanner.Text()))
	}


	fabric := [1000][1000]int{}

	for _, c := range claims {
		for i := 0; i < c.w; i++ {
			for j := 0; j < c.h; j++ {
				fabric[i + c.x][j + c.y]++
			}
		}
	}

	for _, c := range claims {
		overlap := false

		for i := 0; i < c.w; i++ {
			for j := 0; j < c.h; j++ {
				if fabric[i + c.x][j + c.y] > 1 {
					overlap = true
				}
			}
		}

		if !overlap {
			fmt.Println("Non overlapping ID:", c.id)
			os.Exit(0)
		}
	}

	fmt.Println("AHHH!")

}
