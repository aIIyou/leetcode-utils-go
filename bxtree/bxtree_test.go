package bxtree

import (
	"fmt"
	"testing"
)

func TestArray2BXTree(t *testing.T) {
	arr := []string{"1", "2", "3", "4", "5", "null", "6", "7", "null", "null", "null", "null", "8"}

	root := Array2BXTree(arr)

	show(root)
}

func show(root *TreeNode) {
	if root != nil {
		fmt.Printf("%d:{", root.Val)
		if root.Left != nil {
			fmt.Printf("%d:", root.Left.Val)
		} else {
			fmt.Printf("null:")
		}
		if root.Right != nil {
			fmt.Printf("%d}\n", root.Right.Val)
		} else {
			fmt.Printf("null}\n")
		}
		show(root.Left)
		show(root.Right)
	}
}
