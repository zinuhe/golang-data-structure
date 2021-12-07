package main

import "fmt"

var count int

// Node represents the components of a binary search tree
type Node struct {
    Key int
    Left *Node  //No need to defined it as pointer?
    Right *Node
}

// Insert will add a node to the tree
// the key to add should not be already in the tree
func (n *Node) Insert (k int) {
    if n.Key < k {
        //move to the right
        //if the right node doesn't exist
        if n.Right == nil {
            n.Right = &Node{Key: k}
        } else {
            n.Right.Insert(k)
        }
    } else if n.Key > k {
        //move to the left
        //if the left node doesn't exist
        if n.Left == nil {
            n.Left = &Node{Key: k}
        } else {
            n.Left.Insert(k)
        }
    }
}

// Search will take in a key value
// and return true if there is a node with that value
func (n *Node)Search(k int) bool {
    count++
    if n == nil {
        // the end of the tree without a match
        return false
    }
    if n.Key < k {
        //move to the right
        return n.Right.Search(k)
    } else if n.Key > k {
        //move to the left
        return n.Left.Search(k)
    }
    return true //there is a match
}

//TODO
// Delete function

// Print function

func main() {
    tree := &Node {Key: 100}
    fmt.Println(tree)
    tree.Insert(500)
    fmt.Println(tree)

    fmt.Println(tree.Search(500))
    // fmt.Println(tree.Search(5000))
    fmt.Println("count:", count)
}
