package main

import (
  "log"
  "fmt"
)

type Activity struct {
  Name string
  // in minutes
  Interval int
}

func (a Activity) ToString() string {
  return fmt.Sprintf("%v: %vmins", a.Name, a.Interval)
}

func main() {
  activities := []Activity{
    Activity{Name:"make bed", Interval:10},
    Activity{Name:"wash face", Interval:5},
  }

  for _, activity := range activities {
    log.Printf(activity.ToString())
  }
}
