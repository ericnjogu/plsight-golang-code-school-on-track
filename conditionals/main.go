package main

import (
  "fmt"
  "os"
)

func main() {
  default_msg := "Hello, I am Mugo"
  args := os.Args
  if len(args) > 1 {
    fmt.Println(args[1])
  } else {
    fmt.Println(default_msg)
  }
}
