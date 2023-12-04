package main

import (
	"os"
	"strings"
	"testing"
)

func TestOne(t *testing.T) {
	fo, _ := os.ReadFile("1.sample")
	fileSlice := strings.Split(string(fo), "\n")
	var totalSum int
	for i := 0; i < len(fileSlice); i++ {
		wlslice := fileSlice[i]

		num := getNum(getFirst(wlslice), getLast(wlslice))
		totalSum += num

	}
	result := totalSum
	expected := 281
	if result != expected {
		t.Errorf("Result: %v, Expected: %v", result, expected)
	}
}
