package main

import "fmt"

type Node struct {
	Key int
	left,
	right *Node
}

type Tree struct {
	root *Node
}

func (t *Tree) Insert(key int) {
	t.root = t.insertRec(t.root, key)
}

func (t *Tree) insertRec(root *Node, key int) *Node {
	if root == nil {
		root = &Node{Key: key}
		return root
	}

	if key < root.Key {
		root.left = t.insertRec(root.left, key)
	} else if key > root.Key {
		root.right = t.insertRec(root.right, key)
	}

	return root
}

func (t Tree) Contains(key int) bool {
	return t.contains(t.root, key)
}

func (t Tree) contains(root *Node, key int) bool {
	if root == nil {
		return false
	}

	if root.Key == key {
		return true
	}

	if key < root.Key {
		return t.contains(root.left, key)
	} else {
		return t.contains(root.right, key)
	}
}

func (t Tree) printInOrder(node *Node) {
	if node == nil {
		return
	}

	t.printInOrder(node.left)
	fmt.Println(node.Key, " ")
	t.printInOrder(node.right)
}

func (t Tree) Print() {
	t.printInOrder(t.root)
	fmt.Println()
}
