package main

import(
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

func main() {
    fmt.Println("Simple Stack implementation")
    var stack Stack
    stack.Print()
    stack.Push("one")
    stack.Print()
    stack.Push("two")
    stack.Print()
    stack.Push("three")
    stack.Print()
    fmt.Println("pop value:", stack.Pop())
    stack.Print()
    fmt.Println("pop value:", stack.Pop())
    stack.Print()
    fmt.Println("pop value:", stack.Pop())
    stack.Print()
    fmt.Println("pop value:", stack.Pop())
    stack.Print()
}
