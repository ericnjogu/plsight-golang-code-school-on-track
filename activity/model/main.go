package model

type Measurable interface {
  Add(amount int)
}

type Activity struct {
  Name string
  Interval int
}

func (a *Activity) Add(increment int) {
  a.Interval += increment
}

type Student struct {
  Name string
  Height int
}

func (s *Student) Add(increment int) {
  s.Height += increment
}
