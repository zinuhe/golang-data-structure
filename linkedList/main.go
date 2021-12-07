package main

import (
<<<<<<< HEAD
         "fmt"
)

func main() {
=======
  "fmt"
)


type node struct {
  data int
  next *node
}

type linkedList struct {
  head *node
  length int
}

// pointer receiver
func (l *linkedList) prepend (n *node) {
  second := l.head
  l.head = n
  l.head.next = second
  l.length++
}

// value receiver
func (l linkedList) printListData() {
  toPrint := l.head
  for l.length != 0 {
    fmt.Println("%d ", toPrint.data)
    toPrint = toPrint.next
    l.length--
  }
}

func (l *linkedList) deleteWithValue(value int) {
  // If the list is empty
  if l.length == 0 {
    return
  }

  // Delete the header of the list
  if l.head.data == value {
    l.head = l.head.next
    l.length--
    return
  }

  previousToDelete := l.head
  for previousToDelete.next.data != value {
    if previousToDelete.next.next == nil {  // Value not in the List
      return
    }
    previousToDelete = previousToDelete.next //switch the pointer to the next node
  }
  previousToDelete.next = previousToDelete.next.next
  l.length--
}

func main() {
  myList := linkedList{}
  node1 := &node{data:11}
  node2 := &node{data:22}
  node3 := &node{data:33}
  node4 := &node{data:44}
  node5 := &node{data:55}

  myList.prepend(node1)
  myList.prepend(node2)
  myList.prepend(node3)
  myList.prepend(node4)
  myList.prepend(node5)

  myList.printListData()
  myList.deleteWithValue(33)
  myList.printListData()

  myList.deleteWithValue(100) //Deleting a value that not exists

  //empty List
  emptyList := linkedList{}
  emptyList.deleteWithValue(10)

>>>>>>> a38d56c4ed165a88a8c750537d4926b11d3ac7ec
}
