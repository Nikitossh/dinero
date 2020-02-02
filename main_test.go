package main

import (
	"fmt"
	"os"
	"testing"
)

var failMessages = []string{"2020-07-02 food __ 200 testCase",
	"2020-07-02 food 200testCase",
	"2020-13-02 food testCase 222",
	"food __ 200 testCase 2020-13-02",
	"2020-07-0 food200 testCase",
	"202-07-02 food 2s00 testCase",
	"202-07-32 fos 200 testCase",
	"2020-07-02 food 200s testCase",
}

var correctMessages = []string{
	"2020-12-31 food 999 shop or smth else",
	"2020-02-25 communication 999 2line",
	"2020-01-15 food 999 shop or smth else",
	"2020-07-02 food 200 testCase",
}

func Test_isValidCost(t *testing.T) {
	for i := range failMessages {
		if isValidCost(failMessages[i]) {
			t.Errorf("Invalid string is VALID. Check validator pattern or failMessages array")
		}
	}
	for i := range correctMessages {
		if !isValidCost(correctMessages[i]) {
			t.Errorf("This should be valid. Check correctMessages or validator")
		}
	}
	fmt.Println("Success test")
}

func Test_CostsFromFile(t *testing.T) {
	filename := "costs_test"
	costs := make([]*Cost, 0)

	// create file if it is not exists
	var _, err = os.Stat(filename)
	if os.IsNotExist(err) {
		var file, err = os.Create(filename)
		if isError(err) {
			return
		}
		defer file.Close()
	}
	defer os.Remove(filename)

	// fill in both with correct and fail data
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()
	for i := range correctMessages {
		file.WriteString(correctMessages[i] + "\n")
	}
	for i := range failMessages {
		file.WriteString(failMessages[i] + "\n")
	}

	// synchronize changes
	err = file.Sync()
	if isError(err) {
		return
	}

	costs = CostsFromFile(filename)
	for i := range costs {
		fmt.Println(costs[i])
	}

}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return err != nil
}
