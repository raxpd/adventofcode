package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isInt(x string) bool {
	_, err := strconv.Atoi(x)
	if err != nil {
		return false
	}
	return true
}

func getNum(x, y string) int {
	strNum := x + y
	num, _ := strconv.Atoi(strNum)
	return num
}

func getFirst(x string) string {
	str := strings.ToLower(x)
	for i := 0; i < len(str); i++ {
		switch {
		case isInt(string(str[i])):
			return string(str[i])
		case len(str) >= i+3 && str[i:i+3] == "one":
			return "1"
		case len(str) >= i+3 && str[i:i+3] == "two":
			return "2"
		case len(str) >= i+5 && str[i:i+5] == "three":
			return "3"
		case len(str) >= i+4 && str[i:i+4] == "four":
			return "4"
		case len(str) >= i+4 && str[i:i+4] == "five":
			return "5"
		case len(str) >= i+3 && str[i:i+3] == "six":
			return "6"
		case len(str) >= i+5 && str[i:i+5] == "seven":
			return "7"
		case len(str) >= i+5 && str[i:i+5] == "eight":
			return "8"
		case len(str) >= i+4 && str[i:i+4] == "nine":
			return "9"

		}
	}
	return ""
}

func getLast(x string) string {
	str := strings.ToLower(x)
	for i := 0; i < len(str); i++ {
		j := len(str) - i
		switch {
		case isInt(string(str[j-1])):
			return string(str[j-1])
		case j-3 >= 0 && str[j-3:j] == "one":
			return "1"
		case j-3 >= 0 && str[j-3:j] == "two":
			return "2"
		case j-5 >= 0 && str[j-5:j] == "three":
			return "3"
		case j-4 >= 0 && str[j-4:j] == "four":
			return "4"
		case j-4 >= 0 && str[j-4:j] == "five":
			return "5"
		case j-3 >= 0 && str[j-3:j] == "six":
			return "6"
		case j-5 >= 0 && str[j-5:j] == "seven":
			return "7"
		case j-5 >= 0 && str[j-5:j] == "eight":
			return "8"
		case j-4 >= 0 && str[j-4:j] == "nine":
			return "9"

		}
	}
	return ""
}

func main() {
	fo, err := os.ReadFile("1.input")
	if err != nil {
		fmt.Print(err)
	}

	fileSlice := strings.Split(string(fo), "\n")
	var totalSum int
	for i := 0; i < len(fileSlice); i++ {
		wlslice := fileSlice[i]

		num := getNum(getFirst(wlslice), getLast(wlslice))
		fmt.Println(wlslice, ":", num)
		totalSum += num

	}
	fmt.Println(totalSum)
}
