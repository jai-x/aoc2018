package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
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

	frequency := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		diff, err := strconv.Atoi(scanner.Text())
		if err != nil {
			die(err)
		}
		frequency += diff
	}

	fmt.Println("Freqency:", frequency)
}
