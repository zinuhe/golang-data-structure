package example01

import (
	"fmt"
)

type Stack []string

func (s *Stack) Push(value string) {
  *s = append(*s, value)
}

func (s *Stack) Pop() (string) {
  index := len(*s) - 1
  if index >= 0 {
      value := (*s)[index]
      *s = (*s)[:index]
      return value
  }
  return ""
}

func (s *Stack) Print() {
  fmt.Println("Stack:", *s)
}
