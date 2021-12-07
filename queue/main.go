package main

import (
	"fmt"
)

type Queue []string

func (q *Queue) enQueue(value string) {
    *q = append(*q, value)
}

func (q *Queue) deQueue() (string, bool) {
    if len(*q) > 0 {
        result := (*q)[0]
        *q = (*q)[1:]
        return result, false
    }
    return "", true
}

func (q *Queue) print() {
    fmt.Println("queue:", *q)
}

func main() {
    var queue Queue
    queue.print()
    queue.enQueue("one")
    queue.print()
    queue.enQueue("two")
    queue.print()
    queue.enQueue("three")
    queue.print()

    if res, err := queue.deQueue(); err == false {
        fmt.Println("Value from Queue:", res)
        queue.print()
    } else {
        fmt.Println("No more values in the Queue")
    }

    if res, err := queue.deQueue(); err == false {
        fmt.Println("Value from Queue:", res)
        queue.print()
    } else {
        fmt.Println("No more values in the Queue")
    }
}
