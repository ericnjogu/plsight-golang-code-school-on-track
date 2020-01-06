package main

import "log"

type Activity struct {
  Name string
  // in minutes
  Interval int
}

func main() {
  activities := []Activity{
    Activity{"make bed", 10},
    Activity{"wash face", 5},
  }

  for _, activity := range activities {
    log.Printf("%v: %vmins", activity.Name, activity.Interval)
  }
}
