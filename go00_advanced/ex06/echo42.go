package main

import (
	"flag"
	"fmt"
	"os"
)

// go build echo42.go
// ./echo42 12 34 555 | cat -e
// ./echo42 -s === a bc def | cat -e
// ./echo42 -n 12 34 567 | cat -e
// ./echo42 -s === -n a bc def | cat -e
// ./echo42 -help

func main() {
	argc := len(os.Args)
	if argc == 1 {
		fmt.Println()
		return
	}
	var (
		n = flag.Bool("n", false, "omit trailing newline")
		s = flag.String("s", "", "separator (default \" \")")
	)
	flag.Parse()
	args := flag.Args()
	size := len(args)
	for i, v := range args {
		if i != 0 {
			if *s != "" {
				fmt.Print(*s)
			} else if i != size {
				fmt.Print(" ")
			}
		}
		fmt.Print(v)
	}
	if !*n {
		fmt.Println()
	}
}
