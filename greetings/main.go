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
  hour24 := time.Now().Hour()
  if hour24 < 10 {
    panic("It's too early")
  }
  fmt.Println(Greet(hour24))
}
