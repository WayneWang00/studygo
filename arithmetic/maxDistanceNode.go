package main

import "fmt"

// 算出二叉树中节点间的最大距离
/*
	总共有三种情况：
		1：根节点的左子树
		2：根节点的右子树
		3：在根节点的左右子树

	分别计算以上三种情况的最大距离，然后取最大值。
*/
type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

type result struct {
	maxDepth    int
	maxDistance int
}

func main() {
	t := []int{1, 2, -1, 3, 4, -1, 5, 4, -1, 3, 2, 1}
	node := CreateTree(t)
	max := GetMaxDistance(node)

	fmt.Println("max distance:", max)
}

func CreateTree(t []int) *treeNode {
	return createTree(t, 0)
}

func createTree(t []int, index int) *treeNode {
	if index >= len(t) {
		return nil
	}
	if t[index] == -1 {
		return nil
	}

	node := &treeNode{
		val:   t[index],
		left:  createTree(t, 2*index+1),
		right: createTree(t, 2*index+2),
	}

	return node
}

func GetMaxDistance(t *treeNode) int {
	ret := getMaxDistance(t)
	return ret.maxDistance
}

func getMaxDistance(t *treeNode) *result {
	if t == nil {
		return &result{
			maxDepth:    -1,
			maxDistance: 0,
		}
	}

	left := getMaxDistance(t.left)
	right := getMaxDistance(t.right)
	maxDepth := checkMax(left.maxDepth+1, right.maxDepth+1)
	maxDistance := checkMax(checkMax(left.maxDistance, right.maxDistance), left.maxDepth+right.maxDepth+2)

	return &result{
		maxDepth:    maxDepth,
		maxDistance: maxDistance,
	}
}

func checkMax(a, b int) int {
	max := a
	if b > max {
		max = b
	}

	return max
}
