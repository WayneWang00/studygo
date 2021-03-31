package main

import (
	"fmt"
	"os"
)

// 判断链表中是否包含环，并输出环的起始节点
/*
	方法1：通过map中key的唯一性来判断是否有环，且能输出环的起始节点
	方法2：定义快慢两个指针，当快的指针等于慢的指针时，说明存在环。但是，不能确定环的起始节点。
*/

type node struct {
	val  int
	next *node
}

func main() {
	node1 := new(node)
	node2 := new(node)
	node3 := new(node)
	node4 := new(node)
	node5 := new(node)
	node1.val = 1
	node2.val = 2
	node3.val = 3
	node4.val = 4
	node5.val = 5
	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5
	node5.next = node3

	fmt.Println(hasCycle(node1))

	work := detectCycle(node1)
	if work == nil {
		fmt.Println("no cycle")
		os.Exit(0)
	}

	fmt.Println(work.val)
}

func CreateLinkList(list []int) *node {
	return createLinkList(list, 0)
}

func createLinkList(list []int, index int) *node {
	if index >= len(list) {
		return nil
	}

	n := &node{
		val:  list[index],
		next: createLinkList(list, index+1),
	}

	return n
}

// 判断是否有环，并输出环的起始节点
func detectCycle(head *node) *node {
	m := make(map[*node]struct{})
	work := head

	for work != nil {
		_, ok := m[work]
		if ok {
			return work
		} else {
			m[work] = struct{}{}
		}

		work = work.next
	}

	return nil
}

// 判断是否有环
func hasCycle(head *node) *node {
	fast := head
	slow := head

	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next

		if slow == fast {
			return slow
		}
	}

	return nil
}
