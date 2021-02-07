package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	fmt.Println(args, args[1], args[2])
	if len(args) == 0 {
		startLogger("")
	} else {
		startLogger(args[2])
	}
}
