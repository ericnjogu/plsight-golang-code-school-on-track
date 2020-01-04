package main

import (
  "testing"
)

func TestLine(t *testing.T) {
  line := Line("*", 3)
  if line != "***" {
    t.Errorf("line should not be %v", line)
  }
}
