package ClosestValue

type BST struct {
	LeftChild, RightChild *BST
	value                 int
}

func FindTheClosestValue(t *BST, key int) int {

	if key >= t.value {
		if t.RightChild == nil {
			return t.value
		}
		FindTheClosestValue(t.RightChild, key)
	} else {
		if t.LeftChild == nil {
			return t.value
		}
		FindTheClosestValue(t.LeftChild, key)
	}
	return -1
}
