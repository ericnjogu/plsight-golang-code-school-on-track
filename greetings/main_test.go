package main

import "testing"

func TestGreet(t *testing.T) {
  cases := map[int]string {
      8: "Good Morning",
      14: "Good Afternoon",
      23: "Good Evening",
    }

  for hour24, expectedGreeting := range cases {
    actualGreeting := Greet(hour24)
    if actualGreeting != expectedGreeting {
      t.Errorf("Greet(%v), expected %q but was %q", hour24, expectedGreeting, actualGreeting)
    }
  }
}
