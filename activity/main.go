package main

import (
  "log"
  "math/rand"
)

type Measurable interface {
  Add(amount int)
}

type Activity struct {
  Name string
  Interval int
}

func (a *Activity) Add(increment int) {
  a.Interval += increment
}

type Student struct {
  Name string
  Height int
}

func (s *Student) Add(increment int) {
  s.Height += increment
}

func main() {
  list := []Measurable {
    &Activity{Name:"make bed", Interval:10},
    &Activity{Name:"wash face", Interval:5},
    &Student{"Ahadi", 100},
    &Student{"Myles", 100},
  }

  for _, item := range list {
    item.Add(rand.Intn(20))
    log.Printf("%#v",item)
  }
}
