package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	argc := len(os.Args)
	if argc == 1 {
		fmt.Println("Invalid argument")
		return
	}
	re := regexp.MustCompile(`[\w\-\._]+@[\w\-\._]+\.[A-Za-z]+`)
	for _, v := range os.Args[1:] {
		if re.MatchString(v) {
			fmt.Printf("%v is a valid email address.\n", v)
		} else {
			fmt.Printf("%v is NOT a valid email address.\n", v)
		}
	}
}
