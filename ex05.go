package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	for i, args := range os.Args[1:] {
		fmt.Println(i, args)
	}
	elapsed := time.Since(start)
	fmt.Println("running time for range in for", elapsed)

	start = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	elapsed = time.Since(start)
	fmt.Println("running time strings.Join", elapsed)
}
