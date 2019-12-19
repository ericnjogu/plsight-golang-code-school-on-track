package main

import (
  "fmt"
  "os"
)

func main() {
  default_msg := "Hello, I am Mugo"
  if len(os.Args) > 1 {
    fmt.Println(os.Args[1])
  } else {
    fmt.Println(default_msg)
  }
}
