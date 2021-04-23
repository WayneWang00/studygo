package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	tree()
}

type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

// 添加结点
func (t *TreeNode) insert(node *TreeNode) {
	if node.data == t.data {
		return
	}

	if node.data > t.data {
		if t.right == nil {
			t.right = node
		} else {
			t.right.insert(node)
		}
	} else {
		if t.left == nil {
			t.left = node
		} else {
			t.left.insert(node)
		}
	}
}

// 查询结点
func (t *TreeNode) search(data int) *TreeNode {
	if t == nil {
		return nil
	}

	if data == t.data {
		return t
	}

	if data > t.data {
		return t.right.search(data)
	}

	if data < t.data {
		return t.left.search(data)
	}

	return nil
}

// 前序遍历
func (t *TreeNode) preOrder() {
	if t == nil {
		return
	}

	fmt.Print(t.data, " ")
	t.left.preOrder()
	t.right.preOrder()
}

// 中序遍历
func (t *TreeNode) midOrder() {
	if t == nil {
		return
	}

	t.left.midOrder()
	fmt.Print(t.data, " ")
	t.right.midOrder()
}

// 后序遍历
func (t *TreeNode) postOrder() {
	if t == nil {
		return
	}

	t.left.postOrder()
	t.right.postOrder()
	fmt.Print(t.data, " ")
}

//层次遍历
func (t *TreeNode) layerOrder() {
	if t == nil {
		return
	}

	var queue []*TreeNode
	var data [][]int
	queue = append(queue, t)

	for len(queue) != 0 {
		l := len(queue)
		newQueue := []*TreeNode{}
		res := []int{}

		for i := 0; i < l; i++ {
			node := queue[i]
			res = append(res, node.data)
			if node.left != nil {
				newQueue = append(newQueue, node.left)
			}
			if node.right != nil {
				newQueue = append(newQueue, node.right)
			}
		}

		queue = newQueue
		data = append(data, res)
	}

	fmt.Printf("data:%#v\n", data)
}

// 初始化结点
func newTreeNode(data int) *TreeNode {
	return &TreeNode{data: data}
}

// 结构-树
func tree() {
	// 创建，查询
	rand.Seed(time.Now().UnixNano())
	root := newTreeNode(50)
	for i := 0; i < 20; i++ {
		data := rand.Intn(100)
		log.Println(data)
		root.insert(newTreeNode(data))
	}
	fmt.Printf("tree: %#v\n", root)
	data := rand.Intn(100)
	fmt.Println("data:", data)
	fmt.Printf("search: %#v\n", root.search(data))

	// 前、中、后和层次序遍历
	fmt.Println("前序遍历：")
	root.preOrder()
	fmt.Println("\n中序遍历：")
	root.midOrder()
	fmt.Println("\n后序遍历：")
	root.postOrder()
	fmt.Println("\n层次遍历：")
	root.layerOrder()
}
