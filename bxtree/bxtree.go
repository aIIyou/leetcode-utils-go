package bxtree

import (
	"strconv"
)

// //TreeNode 对外输出一个接口，可以在外部assert成想要的类型
// type TreeNode interface{}

//TreeNode 二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	index         = 0
	parentQueue   = make([]*TreeNode, 0, 10)
	childrenQueue = make([]*TreeNode, 0, 10)
)

//Array2BXTree 将数组中的元素构造成二叉树
func Array2BXTree(arr []string) *TreeNode {
	if len(arr) <= 0 || arr[0] == "null" {
		return nil
	}

	rootVal, _ := strconv.ParseInt(arr[0], 10, 0)
	root := &TreeNode{Val: int(rootVal)}
	parentQueue = append(parentQueue, root)

	index++

	bfs(root, arr)
	return root
}

func bfs(root *TreeNode, arr []string) {
	for index < len(arr) {
		n := 2 * len(parentQueue)
		n += index
		if len(arr) < n {
			n = len(arr)
		}
		for index < n {
			item := arr[index]
			if item == "null" {
				childrenQueue = append(childrenQueue, nil)
			} else {
				nodeVal, _ := strconv.ParseInt(arr[index], 10, 0)
				node := &TreeNode{Val: int(nodeVal)}
				childrenQueue = append(childrenQueue, node)
			}
			index++
		}
		for i := 0; i < len(parentQueue); i++ {
			parentQueue[i].Left = childrenQueue[2*i]
			parentQueue[i].Right = childrenQueue[2*i+1]
		}
		parentQueue = make([]*TreeNode, 0, 10)
		for i := 0; i < len(childrenQueue); i++ {
			if childrenQueue[i] != nil {
				parentQueue = append(parentQueue, childrenQueue[i])
			}
		}
		childrenQueue = make([]*TreeNode, 0, 10)
	}
}
