package main

import (
	"fmt"
	"strconv"
)

func calc(expression string) (float64, error) {
	return strconv.ParseFloat(expression, 64)
}

func main() {
	var input string
	fmt.Scan(input)
	fmt.Print(calc(input))
}
