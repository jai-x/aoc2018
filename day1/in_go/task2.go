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

	list := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			die(err)
		}
		list = append(list, val)
	}

	frequency := 0
	seen := map[int]bool{}
	loop := 1

	for {
		for _, diff := range list {
			fmt.Println("Loop:", loop)
			frequency += diff
			_, exists := seen[frequency]

			if exists {
				fmt.Println("Duplicate frequency found:", frequency)
				os.Exit(0)
			} else {
				seen[frequency] = true
			}
			loop++
		}
	}

}
