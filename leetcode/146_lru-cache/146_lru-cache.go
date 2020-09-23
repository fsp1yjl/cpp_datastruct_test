/*
link: https://leetcode.com/problems/lru-cache/

Design and implement a data structure for Least Recently Used (LRU) cache. It should support the following operations: get and put.

get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
put(key, value) - Set or insert the value if the key is not already present. When the cache reached its capacity, it should invalidate the least recently used item before inserting a new item.

The cache is initialized with a positive capacity.

Follow up:
Could you do both operations in O(1) time complexity?

Example:

LRUCache cache = new LRUCache( 2  );

cache.put(1, 1);
cache.put(2, 2);
cache.get(1);       // returns 1
cache.put(3, 3);    // evicts key 2
cache.get(2);       // returns -1 (not found)
cache.put(4, 4);    // evicts key 1
cache.get(1);       // returns -1 (not found)
cache.get(3);       // returns 3
cache.get(4);       // returns 4

*/

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

/*
提交结果：
Runtime: 124 ms, faster than 76.77% of Go online submissions for LRU Cache.
Memory Usage: 17.4 MB, less than 34.85% of Go online submissions for LRU Cache.

反思：
后面可以考虑将head， tail节点不做实际到存储节点，以简化链表的操作

*/
package main

import "fmt"

func main() {
	LRU := Constructor(3)
	// fmt.Println(LRU.Get(3))
	LRU.Put(1, 101)
	// LRU.Console()
	LRU.Put(2, 102)
	LRU.Console()
	LRU.Get(1)
	LRU.Console()

	LRU.Put(3, 103)

	LRU.Put(4, 104)
	LRU.Console()
	// fmt.Println(LRU.Get(1))

}

type Node struct {
	key   int
	value int
	pre   *Node
	next  *Node
}

type LRUCache struct {
	cap    int
	Bucket map[int]*Node
	Head   *Node
	Tail   *Node
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cap:    capacity,
		Bucket: make(map[int]*Node, capacity),
	}

	return l

}

func (this *LRUCache) Get(key int) int {
	if item, ok := this.Bucket[key]; ok {
		this.RefreshNode(item)
		return item.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {

	if item, ok := this.Bucket[key]; ok {
		//如果已经存在，则只需要更新bucket对应的item内容
		item.value = value
		this.RefreshNode(this.Bucket[key])

	} else {
		//如果不存在，则先判断是否需要剔除head
		if len(this.Bucket) >= this.cap {
			this.RemoveNode(this.Head)
		}

		node := &Node{
			key:   key,
			value: value,
		}

		this.AddNode(node)
		// this.Bucket[key] = node

	}

}

func (this *LRUCache) RefreshNode(node *Node) {
	if this.Tail == node {
		return
	}

	this.RemoveNode(node)
	// 注意，这里在重新把节点添加到尾部到时候需要先把节点前后指针信息清空
	node.pre = nil
	node.next = nil
	this.AddNode(node)
}

func (this *LRUCache) AddNode(node *Node) {

	if this.Tail != nil {
		this.Tail.next = node
		node.pre = this.Tail
	}
	this.Tail = node

	if this.Head == nil {
		this.Head = node
	}

	this.Bucket[node.key] = node

}

func (this *LRUCache) RemoveNode(node *Node) {
	pre := node.pre
	next := node.next

	// if node == this.Head {
	// 	this.Head = this.Head.next
	// }
	// if node == this.Tail {
	// 	this.Tail = this.Tail.pre
	// }

	if pre != nil && next != nil {
		//删除中间元素
		pre.next = next
		next.pre = pre
	} else if pre == nil && next == nil {
		// 仅含有一个元素，且删除之
		this.Head = nil
		this.Tail = nil
	} else if pre != nil && next == nil {
		//大于一个元素，删除尾部
		this.Tail = this.Tail.pre
	} else if pre == nil && next != nil {
		// 大于一个元素，删除头部
		next.pre = nil
		this.Head = next

	}
	delete(this.Bucket, node.key)
}

func (this *LRUCache) Console() {
	fmt.Println("start-----------")
	n := this.Head
	for n != nil {
		fmt.Println("k,v", n.key, n.value)
		n = n.next
	}
	fmt.Println("end-----------")

}
