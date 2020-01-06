package main

import "log"

type Activity struct {
  Name string
  // in minutes
  Interval int
}

func main() {
  activities := []Activity{
    Activity{Name:"make bed", Interval:10},
    Activity{Name:"wash face", Interval:5},
  }

  for _, activity := range activities {
    log.Printf("%v: %vmins", activity.Name, activity.Interval)
  }
}
