// Run: go run ex1_4.go text1 text2
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2 : %v\n", err)
				continue
			}
			fmt.Println(arg)
			countLines(f, counts)
			for line, n := range counts {
				if n > 1 {
					fmt.Printf("%d\t%s\n", n, line)
				}
			}
			f.Close()
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
