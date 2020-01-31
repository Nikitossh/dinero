package main

import "testing"

var failMessages = []string{"2020-07-02 food __ 200 testCase",
	"2020-07-02 food 200testCase",
	"2020-13-02 food testCase 222",
	"food __ 200 testCase 2020-13-02",
	"2020-07-0 food200 testCase",
	"202-07-02 food 2s00 testCase",
	"202-07-32 fos 200 testCase",
	"2020-07-02a food 200 testCase",
}

var correctMessages = []string{
	"2020-12-31 food 999 shop or smth else",
	"2020-02-25 communication 999 2line",
	"2020-01-15 food 999 shop or smth else",
	"2020-07-02 food 200 testCase",
}

func Test_isValidCost(t *testing.T) {
	for i, _ := range failMessages {
		if isValidCost(failMessages[i]) {
			t.Errorf("Invalid string is VALID. Check validator pattern or failMessages array")
		}
	}
	for i, _ := range correctMessages {
		if !isValidCost(correctMessages[i]) {
			t.Errorf("This should be valid. Check correctMessages or validator")
		}
	}

}
