package main

import (
	"os"
	"fmt"
	"bufio"
	"unicode"
	"strings"
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

	results := map[string]int{}
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	for _, char := range alphabet {
		r := strings.NewReplacer(string(char), "", strings.ToUpper(string(char)), "")
		chain := r.Replace(polymer)

		fmt.Println("Replacing characters:", string(char), strings.ToUpper(string(char)))
		fmt.Println("Initial chain length:", len(chain))

		lastLen := 0

		for {
			chain = react(chain)

			fmt.Print(".")

			if lastLen == len(chain) {
				fmt.Println()
				fmt.Println("Fully reacted, Chain length:", len(chain))
				results[string(char)] = len(chain)
				break
			}

			lastLen = len(chain)
		}
	}

	fmt.Println(results)

	minLen := 999999999
	minLetter := ""

	for letter, length := range results {
		if length < minLetter {
			minLetter, minLen = letter, length
		}
	}

	fmt.Println("Smallest chain with removed letter:" minLetter, minLen)
}
