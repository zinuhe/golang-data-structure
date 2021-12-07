package main

import "fmt"

var count int

// BinaryNode represents the components of a binary search tree
type BinaryNode struct {
    Data  int
    Left  *BinaryNode  //No need to defined it as pointer?
    Right *BinaryNode
}

// Insert will add a node to the tree
// the data to add should not be already in the tree
func (n *BinaryNode) Insert(d int) {
    if n == nil {
        return
    } else if n.Data < d {
        //move to the right
        //if the right node doesn't exist
        if n.Right == nil {
            n.Right = &BinaryNode{Data: d}
        } else {
            n.Right.Insert(d)
        }
    } else if n.Data > d {
        //move to the left
        //if the left node doesn't exist
        if n.Left == nil {
            n.Left = &BinaryNode{Data: d}
        } else {
            n.Left.Insert(d)
        }
    }
}

// Search will take in a data value
// and return true if there is a node with that value
func (n *BinaryNode)Search(d int) bool {
    count++
    if n == nil {
        // the end of the tree without a match
        return false
    }
    if n.Data < d {
        //move to the right
        return n.Right.Search(d)
    } else if n.Data > d {
        //move to the left
        return n.Left.Search(d)
    }
    return true //there is a match
}

//TODO
// Delete function

// Print functions traverse
func (n *BinaryNode) printPreOrder() {
	if n == nil {
		return
	} else {
		fmt.Printf("%d ", n.data)
		printPreOrder(n.left)
		printPreOrder(n.right)
	}
}
func (n *BinaryNode) printPostOrder() {
	if n == nil {
		return
	} else {
		printPostOrder(n.left)
		printPostOrder(n.right)
		fmt.Printf("%c ", n.data)
	}
}
func (n *BinaryNode) printInOrder() {
	if n == nil {
		return
	} else {
		printInOrder(n.left)
		fmt.Printf("%c ", n.data)
        printInOrder(n.right)
	}
}

func main() {
    tree := &BinaryNode {Data: 100}
    fmt.Println(tree)
    tree.Insert(500)
    fmt.Println(tree)

    fmt.Println(tree.Search(500))
    // fmt.Println(tree.Search(5000))
    fmt.Println("count:", count)
}
