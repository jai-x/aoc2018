package main

import (
	"os"
	"fmt"
	"bufio"
)

type CommonRune map[rune]int

func NewCommonRune(line string) CommonRune {
	out := make(map[rune]int)
	for _, v := range line {
		out[v]++
	}
	return out
}

func (c CommonRune) Twos() bool {
	for _, v := range c {
		if v == 2 {
			return true
		}
	}
	return false
}

func (c CommonRune) Threes() bool {
	for _, v := range c {
		if v == 3 {
			return true
		}
	}
	return false
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

	twos := 0
	threes := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		common := NewCommonRune(scanner.Text())
		fmt.Println(scanner.Text(), "2:", common.Twos(), "3:", common.Threes())

		if common.Twos() {
			twos++
		}

		if common.Threes() {
			threes++
		}
	}

	fmt.Println("Checksum:", twos * threes)
}
