package main

import (
	"fmt"
)

type IntTree struct {
	val         int
	left, right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}
	if val < it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		it.right = it.right.Insert(val)
	}
	return it
}

func (it *IntTree) Contains(val int) bool {
	if it == nil {
		return false
	}
	if val < it.val {
		return it.left.Contains(val)
	} else if val > it.val {
		return it.right.Contains(val)
	}
	return true
}

func main() {
	// tree := &IntTree{}
	var tree *IntTree
	fmt.Printf("%p\n", tree)
	fmt.Println(tree == nil)
	tree = tree.Insert(5)
	fmt.Printf("%p\n", tree)
	tree = tree.Insert(3)
	tree = tree.Insert(7)

	fmt.Println(tree.Contains(3))
	fmt.Println(tree.Contains(4))

}
