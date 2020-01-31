package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
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
	var skipped int = 0
	var added int = 0
	var total int = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cost := scanner.Text()
		total++
		if isValidCost(cost) {
			fmt.Println(cost)
			added++
		} else {
			skipped++
		}
	}
	fmt.Printf("Total lines in file %d \t\t Correct: %d  \t\t Skipped: %d", total, added, skipped)
}

func isValidCost(s string) bool {
	// valid cost by regex, must be smth like:
	// YYYY-MM-DD category value comment
	trimmed := strings.TrimSpace(s)
	var validCost = regexp.MustCompile(`^((19|20)\d\d-(0[1-9]|1[012])-(0[1-9]|[12][0-9]|3[01]))\s+(\w+)\s+([0-9]{1,9})\s+(...{1,255}?)$`)

	if validCost.MatchString(trimmed) {
		return true
	} else {
		return false
	}
}

func main() {
	readFile()
}
