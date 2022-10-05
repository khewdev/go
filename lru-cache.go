package main

import "fmt"

type LRUCache struct {
    Head *Node
    Tail *Node
    HT map[int]*Node
    Cap int
}

type Node struct {
    Key int
    Val int
    Prev *Node
    Next *Node
}

func Constructor(capacity int) LRUCache {
    return LRUCache{HT: make(map[int]*Node), Cap: capacity}
}

func (this *LRUCache) Get(key int) int {
    node, ok := this.HT[key]
    if ok {
        this.Remove(node)
        this.Add(node)
        return node.Val
    }
    return -1
}

func (this *LRUCache) Put(key int, value int)  {
    node, ok := this.HT[key]
    if ok {
        node.Val = value
        this.Remove(node)
        this.Add(node)
        return
    } else {
        node = &Node{Key: key, Val: value}
        this.HT[key] = node
        this.Add(node)
    }
    if len(this.HT) > this.Cap {
        delete(this.HT, this.Tail.Key)
        this.Remove(this.Tail)
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
	fmt.Println("node", node)
	fmt.Println("this.Head", this.Head)
	fmt.Println("this.Tail", this.Tail)
	fmt.Println()
}

func (this *LRUCache) Remove(node *Node) {
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

func main() {
	lru := Constructor(2)
	fmt.Println(lru)
	// fmt.Println("Adding 1")
	// lru.Add(&Node{Key:1, Val:1})
	// fmt.Println("Adding 2")
	// lru.Add(&Node{Key:2, Val:2})
	// fmt.Println("Adding 3")
	// lru.Add(&Node{Key:3, Val:3})
    // fmt.Println("Adding 4")
	// lru.Add(&Node{Key:4, Val:4})


	fmt.Println("Putting 1")
    lru.Put(1,1);
    fmt.Println("Putting 2")
    lru.Put(2,2);
    fmt.Println("Putting 3")
    lru.Put(3,3);
    fmt.Println("Putting 4")
    lru.Put(4,4);
    fmt.Println("Putting 5")
    lru.Put(5,5);
}