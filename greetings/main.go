package main

import (
  "time"
  "fmt"
)

func Greet(hour24 int) string {
  switch {
  case hour24 < 12:
    return "Good Morning"
  case hour24 >= 12 && hour24 <  18:
    return "Good Afternoon"
  default://case hour24 >= 18:
    return "Good Evening"
  }
}

func main() {
  fmt.Println(Greet(time.Now().Hour()))
}
