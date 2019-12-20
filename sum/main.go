package main

import (
	"fmt"
	"os"
)

/*
sum up all args
*/
func main() {
	var sum int
	for _, arg := range os.Args[1:] {
		sum += arg
	}
	fmt.Println(sum)
}
