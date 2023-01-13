package main

import (
	"fmt"
	"os"
	"strconv"
)

const ERROR_MSG string = "Arguments is invalid."

func calculationStr(args []string) (string, bool) {
	var s string
	var ok bool
	if len(os.Args) != 3 {
		return s, ok
	}
	num1, err := strconv.Atoi(args[1])
	if err != nil {
		return s, ok
	}
	num2, err := strconv.Atoi(args[2])
	if err != nil || num2 == 0 {
		return s, ok
	}
	a := "sum: " + strconv.Itoa(num1+num2) + "\n"
	b := "difference: " + strconv.Itoa(num1-num2) + "\n"
	c := "product: " + strconv.Itoa(num1*num2) + "\n"
	d := "quotient: " + strconv.Itoa(num1/num2) + "\n"
	s = a + b + c + d
	ok = true
	return s, ok
}

func main() {
	s, ok := calculationStr(os.Args)
	if !ok {
		fmt.Println(ERROR_MSG)
		os.Exit(1)
	}
	fmt.Print(s)
}
