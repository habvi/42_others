package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	argc := len(os.Args)
	if argc == 1 {
		fmt.Println()
		return
	}
	var (
		s = flag.String("s", "", "separator (default ' ')") // ' ' -> " "
		n = flag.String("n", "", "omit trailing newline")   // change none type..!?
	)
	flag.Parse()
	args := flag.Args()
	size := len(args)
	for i, v := range args {
		if i == 0 && *n != "" {
			fmt.Print(*n)
		}
		if i != 0 {
			if *s != "" {
				fmt.Print(*s)
			} else if i != size {
				fmt.Print(" ")
			}
		}
		fmt.Print(v)
	}
	if size == 0 && *n != "" {
		fmt.Print(*n)
	} else {
		fmt.Println()
	}
}
