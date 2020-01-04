package main

import (
  "fmt"
  "os"
  "strconv"
  "log"
)

/*
* returns a string that is composed of the given character
*/
func Line(char string, length int) string {
  line := ""
  for len := 0; len < length; len++ {
    line = line + char
  }
  return line
}

func main() {
  length, _ := strconv.Atoi(os.Args[2])
  log.Printf("%v", os.Args)
  fmt.Println(Line(os.Args[1], length))
}
