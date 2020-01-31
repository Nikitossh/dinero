package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInput() {
	// To create dynamic array
	arr := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter Value: ")
		scanner.Scan()
		// Holds the string that scanned
		text := scanner.Text()
		if len(text) != 0 {
			fmt.Println(text)
			arr = append(arr, text)
		} else {
			break
		}
	}
	// Use collected inputs
	fmt.Println(arr)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile() {
	file, err := os.Open("/tmp/costs")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	readFile()
}