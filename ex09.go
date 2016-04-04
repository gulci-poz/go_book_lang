package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
    occurs:= make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, occurs, "stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "duplicate: %v\n", err)
				continue
			}
			countLines(f, counts, occurs, arg)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, occurs[line], line)
		}
	}
}

func countLines(f *os.File, counts map[string]int, occurs map[string]string, src string) {

    input := bufio.NewScanner(f)
    sep := " "

	for input.Scan() {
        line := input.Text()
		counts[line]++

        // na razie uzupełniam stringa, można zrobić listę stringów i za każdym razem sprawdzać czy nazwa pliku nie została juz na nią dodana

        if counts[line] == 1 {
            occurs[line] += src
        } else {
            occurs[line] += sep + src
        }
    }
}
