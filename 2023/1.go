package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getFirst(x string) int {
	var charInt int
	for i := 0; i < len(x); i++ {
		char := string(x[i])
		if isInt(char) {
			charInt, _ = strconv.Atoi(char)
			break
		}

	}
	return charInt
}

func getLast(x string) int {
	var charInt int
	for i := 1; i <= len(x); i++ {
		char := string(x[len(x)-i])
		if isInt(char) {
			charInt, _ = strconv.Atoi(char)
			break
		}

	}
	return charInt
}

func isInt(x string) bool {
	_, err := strconv.Atoi(x)
	if err != nil {
		return false
	}
	return true
}

func getNum(x, y int) int {
	strNum := strconv.Itoa(x) + strconv.Itoa(y)
	num, _ := strconv.Atoi(strNum)
	return num
}

func main() {
	fo, err := os.ReadFile("1.input")
	if err != nil {
		fmt.Print(err)
	}

	fileSlice := strings.Split(string(fo), "\n")
	fmt.Println(fileSlice)
	var totalSum int
	for i := 0; i < len(fileSlice); i++ {
		wslice := fileSlice[i]

		num := getNum(getFirst(wslice), getLast(wslice))
		totalSum += num
		fmt.Println(wslice, "", num)
	}
	fmt.Println(totalSum)
}
