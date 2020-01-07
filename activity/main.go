package main

import (
  "log"
  "math/rand"
  "github.com/kunadawa/plsight-golang-code-school/activity/model"
)

func main() {
  list := []model.Measurable {
    &model.Activity{Name:"make bed", Interval:10},
    &model.Activity{Name:"wash face", Interval:5},
    &model.Student{"Ahadi", 100},
    &model.Student{"Myles", 100},
  }

  for _, item := range list {
    item.Add(rand.Intn(20))
    log.Printf("%#v",item)
  }
}
