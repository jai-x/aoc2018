package main

import (
	"os"
	"fmt"
	"bufio"
	"unicode"
)

func die(v ...interface{}) {
	fmt.Println(v...)
	os.Exit(1)
}

func invertCase(b byte) byte {
	r := rune(b)
	var res rune
	if unicode.IsUpper(r) {
		res = unicode.ToLower(r)
	} else {
		res = unicode.ToUpper(r)
	}
	return byte(res)
}

func react(polymer string) string {
	if len(polymer) < 2 {
		return polymer
	}

	if polymer[0] == invertCase(polymer[1]) {
		return react(polymer[2:])
	}

	return string(polymer[0]) + react(polymer[1:])
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

	var polymer string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		polymer = scanner.Text()
	}

	fmt.Println("Initial chain length:", len(polymer))

	lastLen := 0
	i := 1

	for {
		polymer = react(polymer)

		fmt.Println("Iteration:", i, "Chain length:", len(polymer))
		i++

		if lastLen == len(polymer) {
			fmt.Println("Fully reacted, Chain length:", len(polymer))
			return
		}

		lastLen = len(polymer)
	}
}
