package main

import "math"

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) FindClosestValue(target int) int {
	// for every branch, check delta of target with left and right
	// whichever is smaller, go there (left or right)
	// store the delta in variable called min
	// if the next checked delta is smaller than min, replace min with that delta
	// and go there (left or right)
	// stop when the next checked delta is bigger than min

	min := int(math.Abs(float64(target - tree.Value)))
	var rightDelta int
	var leftDelta int

	for tree.Left != nil || tree.Right != nil {
		if tree.Right != nil {
			rightDelta = int(math.Abs(float64(target - tree.Right.Value)))
		}

		if tree.Left != nil {
			leftDelta = int(math.Abs(float64(target - tree.Left.Value)))
		}

		if rightDelta > min && leftDelta > min {
			break
		}

		if rightDelta < leftDelta && tree.Right != nil {
			tree = tree.Right
			min = rightDelta
		} else if leftDelta < rightDelta && tree.Left != nil {
			tree = tree.Left
			min = leftDelta
		}
	}
	return tree.Value
}
