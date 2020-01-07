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

func (a *Activity) UpdateName(newName string) {
  a.Name = newName
}

func main() {
  activities := []Activity{
    Activity{Name:"make bed", Interval:10},
    Activity{Name:"wash face", Interval:5},
  }

  for _, activity := range activities {
    activity.UpdateName(fmt.Sprintf("Updated: %v", activity.Name))
    log.Printf(activity.ToString())
  }
}
