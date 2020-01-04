package main

import (
  "fmt"
  "os"
  "strconv"
)

/*
* returns a string that is composed of the given character
*/
func Line(char string, length int) string {
  line := ""
  len := 0
  for  {
    if len > length {
      break
    }
    line = line + char
    len++
  }
  return line
}

func main() {
  length, _ := strconv.Atoi(os.Args[2])
  fmt.Println(Line(os.Args[1], length))
}
