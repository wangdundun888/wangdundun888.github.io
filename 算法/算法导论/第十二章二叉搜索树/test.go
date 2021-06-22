package main

import (
	"container/list"
	"fmt"
)

// 本程序主要测试从递归遍历变为stack非递归遍历

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func traverse(root *Node) {
	if root != nil {
		traverse(root.Left)
		fmt.Print(root.Val, " ")
		traverse(root.Right)
	}
}

func main() {
	root := &Node{4, nil, nil}
	left := &Node{2, &Node{1, nil, nil}, &Node{3, nil, nil}}
	right := &Node{6, &Node{5, nil, nil}, &Node{7, nil, nil}}
	root.Left = left
	root.Right = right
	//递归遍历
	traverse(root)
	fmt.Println()

	//非递归遍历
	stack := list.New()
	stack.Init()
	stack.PushFront(root)
	for stack.Len() != 0 {
		e := stack.Front()
		currNode := e.Value.(*Node)
		if currNode.Left != nil {
			stack.PushFront(currNode.Left)
			currNode.Left = nil
		} else {
			fmt.Print(currNode.Val, " ")
			stack.Remove(e)
			if currNode.Right != nil {
				stack.PushFront(currNode.Right)
			}
		}

	}
	fmt.Println()
}
