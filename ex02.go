package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string

	// nie musi być inicjalizacji ani post
	// bez warunku mamy z kolei pętlę nieskończoną
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println("program name:", os.Args[0])
	fmt.Println("args:", s)
	fmt.Println("total number of args", len(os.Args))
}
