package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main( ) {
  start1 := time.Now()
  var s, sep string
  for i := 1; i < len(os.Args); i++ {
    s += sep + os.Args[i]
    sep = " "
  }
  fmt.Println(s)
  fmt.Printf("ver1: %d Microseconds elapsed\n", time.Since(start1).Microseconds())

  start2 := time.Now()
  var s2 string
  for _,arg := range(os.Args[1:]) {
    s2 += arg + " "
  }
  fmt.Println(s2)
  fmt.Printf("ver2: %d Microseconds elapsed\n", time.Since(start2).Microseconds())

  start3 := time.Now()
  fmt.Println(strings.Join(os.Args[1:], " "))
  fmt.Printf("ver3: %d Microseconds elapsed\n", time.Since(start3).Microseconds())
}
