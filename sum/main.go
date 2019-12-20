package main

import (
	"fmt"
	"os"
  "strconv"
  "log"
)

/*
sum up all args
*/
func main() {
	var sum int
	for _, arg := range os.Args[1:] {
    num, err := strconv.Atoi(arg)
		if err == nil {
      sum += num
    } else {
      log.Printf("num: %v, err: %v", num, err)
    }
	}
	fmt.Println(sum)
}
