package main

import (
	"os"
	"fmt"
	"bufio"
)

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

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

}
