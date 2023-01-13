package main

import (
	"fmt"
	"os"
	"strconv"
)

const ERROR_MSG string = "Arguments is invalid."

func main() {
	argc := len(os.Args)
	if argc != 2 {
		fmt.Println(ERROR_MSG)
		return
	}
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(ERROR_MSG)
		return
	}
	star := 1
	for star <= num {
		for i := 0; i < star; i++ {
			fmt.Print("*")
			num--
		}
		fmt.Println()
		star++
	}
}
