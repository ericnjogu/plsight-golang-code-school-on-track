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
  var badArgs []string
	for _, arg := range os.Args[1:] {
    num, err := strconv.Atoi(arg)
		if err == nil {
      sum += num
    } else {
      badArgs = append(badArgs, arg)
    }
	}
  log.Printf("invalid args: %v", badArgs)
	fmt.Println(sum)
}
