package main

import (
  "fmt"
  "github.com/kunadawa/plsight-golang-code-school/activity/model"
  "sync"
  "strconv"
  "os"
)

func AddUp(activity *model.Activity, limit int, waitGroup *sync.WaitGroup) {
  for activity.Interval < limit {
    activity.Interval++
  }
  waitGroup.Done()
}

func Monitor(activity *model.Activity, limit int, waitGroup *sync.WaitGroup) {
  for activity.Interval < limit {
    // 'print-in-place credit - https://socketloop.com/tutorials/golang-overwrite-previous-output-with-count-down-timer'
    fmt.Printf("\r%v", activity.Interval)
  }
  waitGroup.Done()
}

func main() {
  limit, _ := strconv.Atoi(os.Args[1])
  var wg sync.WaitGroup
  wg.Add(2)
  activity := &model.Activity{"adding up", 0}
  go AddUp(activity, limit, &wg)
  go Monitor(activity, limit, &wg)
  wg.Wait()
  fmt.Println()
}
