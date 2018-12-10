package main

import (
	"os"
	"fmt"
	"bufio"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Number of characters the strings differ by
func diff(a, b string) int {
	if len(a) != len(b) {
		return abs(len(a) - len(b))
	}

	diffs := 0
	for i := range a {
		if a[i] != b[i] {
			diffs++
		}
	}
	return diffs
}

func merge(a, b string) string {
	if len(a) != len(b) {
		return ""
	}

	out := []byte{}
	for i := range a {
		if a[i] == b[i] {
			out = append(out, a[i])
		}
	}
	return string(out)
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

	list := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}

	for i, _ := range list {
		for j, _ := range list {
			if diff(list[i], list[j]) == 1 {
				fmt.Printf("IDs found:\n %s\n %s\n %s\n", list[i], list[j], merge(list[i], list[j]))
				os.Exit(0)
			}
		}
	}

	fmt.Println("No matching IDSs")
}
