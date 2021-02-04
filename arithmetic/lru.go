package main

import (
	"container/list"
	"errors"
	"fmt"
	"sync"
)

type Node struct {
	Key   interface{}
	Value interface{}
	pre   *Node
	next  *Node
}

type LRUCache struct {
	limit   int
	HashMap map[interface{}]*Node
	head    *Node
	end     *Node
}

func (l *LRUCache) removeNode(node *Node) interface{} {
	if node == l.end {
		l.end = l.end.pre
		l.end.next = nil
	} else if node == l.head {
		l.head = l.head.next
		l.head.pre = nil
	} else {
		node.pre.next = node.next
		node.next.pre = node.pre
	}

	return node.Key
}

func (l *LRUCache) addNode(node *Node) {
	if l.end != nil {
		l.end.next = node
		node.pre = l.end
		node.next = nil
	}
	l.end = node
	if l.head == nil {
		l.head = node
	}
}

func (l *LRUCache) refreshNode(node *Node) {
	if node == l.end {
		return
	}
	l.removeNode(node) // 从链表中的任意位置移除原来的位置
	l.addNode(node)    // 添加到链表的尾部
}

// 构造
func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{limit: capacity}
	lruCache.HashMap = make(map[interface{}]*Node, capacity)
	return lruCache
}

// 获取
func (l *LRUCache) Get(key interface{}) interface{} {
	if v, ok := l.HashMap[key]; ok {
		l.refreshNode(v)
		return v.Value
	} else {
		return -1
	}
}

func (l *LRUCache) Put(key, value interface{}) {
	if v, ok := l.HashMap[key]; !ok {
		if len(l.HashMap) >= l.limit {
			oldKey := l.removeNode(l.head)
			delete(l.HashMap, oldKey)
		}
		node := Node{Key: key, Value: value}
		l.addNode(&node)
		l.HashMap[key] = &node
	} else {
		v.Value = value
		l.refreshNode(v)
	}
}

func (l *LRUCache) getCache() {
	for n := l.head; n != nil; n = n.next {
		fmt.Println(n.Key, n.Value)
	}
}

func lru1(n int) {
	cache := Constructor(n)
	cache.Put(11, 1)
	fmt.Println("add 11:")
	cache.getCache()
	cache.Put(22, 2)
	fmt.Println("add 22:")
	cache.getCache()
	cache.Put(33, 3)
	fmt.Println("add 33:")
	cache.getCache()
	cache.Put(44, 4)
	fmt.Println("add 44:")
	cache.getCache()
	v := cache.Get(33)
	fmt.Println(v)
	fmt.Println("========== 获取数据之后 ===============")
	cache.getCache()
}

func main() {
	//lru1(3)
	lru2(3)
}

func lru2(n int) {
	lru := NewLru(n)
	lru.Add(11, 1)
	fmt.Println("add 11:")
	getCache(lru.GetAll())
	lru.Add(22, 2)
	fmt.Println("add 22:")
	getCache(lru.GetAll())
	lru.Add(33, 3)
	fmt.Println("add 33:")
	getCache(lru.GetAll())
	lru.Add(44, 4)
	fmt.Println("add 44:")
	getCache(lru.GetAll())
	v, _ := lru.Get(33)
	fmt.Println("get 33:", v)
	fmt.Println("end:")
	getCache(lru.GetAll())
}

func getCache(nodes []*Node2) {
	for i := 0; i < len(nodes); i++ {
		fmt.Println(nodes[i].Key, nodes[i].Val)
	}
}

type Lru struct {
	max   int
	l     *list.List
	Call  func(key interface{}, value interface{})
	cache map[interface{}]*list.Element
	mu    *sync.Mutex
}
type Node2 struct {
	Key interface{}
	Val interface{}
}

func NewLru(len int) *Lru {
	return &Lru{
		max:   len,
		l:     list.New(),
		cache: make(map[interface{}]*list.Element),
		mu:    new(sync.Mutex),
	}
}
func (l *Lru) Add(key interface{}, val interface{}) error {
	if l.l == nil {
		return errors.New("not init NewLru")
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	if e, ok := l.cache[key]; ok {
		e.Value.(*Node2).Val = val
		l.l.MoveToFront(e)
		return nil
	}
	ele := l.l.PushFront(&Node2{
		Key: key,
		Val: val,
	})
	l.cache[key] = ele
	if l.max != 0 && l.l.Len() > l.max {
		if e := l.l.Back(); e != nil {
			l.l.Remove(e)
			node := e.Value.(*Node2)
			delete(l.cache, node.Key)
			if l.Call != nil {
				l.Call(node.Key, node.Val)
			}
		}
	}
	return nil
}
func (l *Lru) Get(key interface{}) (val interface{}, ok bool) {
	if l.cache == nil {
		return
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	if ele, ok := l.cache[key]; ok {
		l.l.MoveToFront(ele)
		return ele.Value.(*Node2).Val, true
	}
	return
}
func (l *Lru) GetAll() []*Node2 {
	l.mu.Lock()
	defer l.mu.Unlock()
	var data []*Node2
	for _, v := range l.cache {
		data = append(data, v.Value.(*Node2))
	}

	return data
}
func (l *Lru) Del(key interface{}) {
	if l.cache == nil {
		return
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	if ele, ok := l.cache[key]; ok {
		delete(l.cache, ele)
		if e := l.l.Back(); e != nil {
			l.l.Remove(e)
			delete(l.cache, key)
			if l.Call != nil {
				node := e.Value.(*Node2)
				l.Call(node.Key, node.Val)
			}
		}
	}
}
