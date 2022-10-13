// go test
// go test -run TestMainStack01
// go test -tags=01_tests -v

package example01_test

import (
	"fmt"
	"testing"

	s01 "github.com/zinuhe/golang-data-structure/stack/example01"
)


func TestMainStack01(t *testing.T) {
	fmt.Println("Test example stack01")

	var stack s01.Stack

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
