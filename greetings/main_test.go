package main

import "testing"

func TestGreetWithinTime(t *testing.T) {
  cases := map[int]string {
      8: "Good Morning",
      14: "Good Afternoon",
      23: "Good Evening",
    }

  for hour24, expectedGreeting := range cases {
    actualGreeting, error := Greet(hour24)
    if len(error) > 0 {
      t.Errorf("There should be no error within time")
    }
    if actualGreeting != expectedGreeting {
      t.Errorf("Greet(%v), expected %q but was %q", hour24, expectedGreeting, actualGreeting)
    }
  }
}

func TestGreetOutsideTime(t *testing.T) {
  actualGreeting, error := Greet(3)
  if len(error) == 0 {
    t.Errorf("There should be an error outside time")
  }
  if len(actualGreeting) > 0 {
    t.Errorf("There should be no greeting")
  }
}
