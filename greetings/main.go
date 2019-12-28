package main

import (
  "time"
  "fmt"
)

func Greet(hour24 int) (string, string) {
  var greeting string
  var error string
  switch {
  case hour24 < 7:
    error = "It's too early"
  case hour24 < 12:
    greeting = "Good Morning"
  case hour24 >= 12 && hour24 <  18:
    greeting = "Good Afternoon"
  default://case hour24 >= 18:
    greeting = "Good Evening"
  }
  return greeting, error
}

func main() {
  hour24 := time.Now().Hour()
  greeting, error := Greet(hour24)
  if len(greeting) > 0 {
      fmt.Println(greeting)
  } else {
    panic(error)
  }
}
