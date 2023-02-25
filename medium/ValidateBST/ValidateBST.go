package main

import (
	Queue "ExercisingAlgoAndDs/queue"
)

type BST struct {
	leftChild, rightChild *BST
	value                 int
	visited               bool
}

func (t *BST) RecurBST() bool {
	if t.value < t.leftChild.value && t.value >= t.rightChild.value {
		return false
	} else if t.rightChild != nil && !t.rightChild.RecurBST() {
		return false
	} else if t.leftChild != nil && !t.leftChild.RecurBST() {
		return false
	}
	return true
}

// ValidateBST Validating BST using queue
func ValidateBST(t *BST) bool {
	queue := Queue.NewQueue[*BST]()
	queue.Enqueue(t)
	for !queue.IsEmpty() {
		node, err := queue.Dequeue()
		if err != nil {
			panic(err)
		}
		if !node.visited {
			if node.value < node.rightChild.value && node.value >= node.leftChild.value {
				return false
			}
			node.visited = true
			queue.Enqueue(node.leftChild)
			queue.Enqueue(node.rightChild)
		}
	}
	return true
}
