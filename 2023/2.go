package main

import (
	"os"
	"strings"
)

func two() {
	fo, err := os.ReadFile("2.sample")
	if err != nil {
		panic(err)
	}
	fileSlice := strings.Split(string(fo), "\n")
}
