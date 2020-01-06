package main

import "log"
/*
"Write a function that given an array of integers, it return a map object of integers that appear multiple times
and the number of times each appears.

Example 1
[1,2,3,4,6,6,7,8,9,5,2,6,1,8] #=> {1:2, 2:2, 6:3, 8:2}"

*/

func FindDuplicates(numbers []string) map[string]int{
  var duplicates map[string]int
  duplicates = make(map[string]int)
  for _, number := range numbers {
    count, present := duplicates[number]
    if present {
      duplicates[number] = count + 1
    } else {
      duplicates[number] = 1
    }
  }
  var notDuplicated []string
  for countedNumber, count := range duplicates {
    if count == 1 {
      delete(duplicates, countedNumber)
      notDuplicated = append(notDuplicated, countedNumber)
    }
  }
  log.Printf("Numbers not duplicated %v", notDuplicated)

  return duplicates
}

func main() {

}
