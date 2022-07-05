package main

type LRUCache struct {
	Head *Node
	Tail *Node
	Map  map[int]*Node
	Cap  int
}

type Node struct {
	Key  int
	Val  int
	Prev *Node
	Next *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{Map: make(map[int]*Node), Cap: capacity}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.Map[key]
	if ok {
		this.Remove(node)
		this.Add(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.Map[key]

	if ok {
		this.Remove(node)
		node.Val = value
		this.Add(node)
	} else {
		if len(this.Map) == this.Cap {
			delete(this.Map, this.Tail.Key)
			this.Remove(this.Tail)
		}
		node = &Node{Key: key, Val: value}
		this.Map[key] = node
		this.Add(node)
	}
}

func (this *LRUCache) Add(node *Node) {
	node.Prev = nil
	node.Next = this.Head

	if this.Head != nil {
		this.Head.Prev = node
	}
	this.Head = node
	if this.Tail == nil {
		this.Tail = node
	}

	// 	headNext := node.Next
	// 	this.Head.Next = node

	// 	node.Prev = this.Head
	// 	node.Next = headNext

	// 	headNext.Prev = node

}

func (this *LRUCache) Remove(node *Node) {
	// 	nextNode := node.Next
	// 	prevNode := node.Prev

	// 	nextNode.Prev = prevNode
	// 	nextNode.Next = nextNode

	if node != this.Head {
		node.Prev.Next = node.Next
	} else {
		this.Head = node.Next
	}
	if node != this.Tail {
		node.Next.Prev = node.Prev
	} else {
		this.Tail = node.Prev
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
