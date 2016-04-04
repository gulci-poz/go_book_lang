package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println("program name:", os.Args[0])
	fmt.Println("args:", s)
	fmt.Println("total number of args", len(os.Args))
}
