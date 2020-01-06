package main

import (
  "testing"
  "reflect"
)

func TestFindDuplicates_01(t *testing.T) {
  actualDuplicates := map[string]int{
    "1":3,
    "2":2,
  }
  duplicates := FindDuplicates([]string{"1", "1", "1", "2", "3", "4", "2"})
  // https://www.geeksforgeeks.org/comparing-maps-in-golang/
  if ! reflect.DeepEqual(actualDuplicates, duplicates) {
    t.Errorf("duplicate results (%v) not match", duplicates)
  }
}
