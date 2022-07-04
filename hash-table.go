/**
* Hash Table implementation in Go
 */

package main

import "fmt"

const ARRAY_SIZE = 7

// HashTable will hold an array
type HashTable struct {
	Array [ARRAY_SIZE] *bucket
}

// Insert will take in a key and add it to the hash table array
func (h *HashTable) Insert (key string) {
	index := hash(key)
	h.Array[index].insert(key)
}

// Search will take in a key and return true if that key is stored in the hash table
func (h *HashTable) Search (key string) bool {
	index := hash(key)
	return h.Array[index].search(key)
}

// Delete will take in a key and delete it from the hash table
func (h *HashTable) Delete (key string) {
	index := hash(key)
	h.Array[index].delete(key)
}

// bucket is a linked list in each slot of the hash table array
type bucket struct {
	Head *bucketNode
}

// insert will take in a key, create a node with the key and insert the node in the bucket
func (b *bucket) insert (k string) {
	if !b.search(k) {

		newNode := &bucketNode{Key:k}
		newNode.Next = b.Head
		b.Head = newNode
	} else {
		fmt.Println(k, "ALREADY EXISTED IN HASH TABLE")
	}
}

// search will take in a key and return true if the bucket has that key
func (b *bucket) search (k string) bool {
	currentNode := b.Head
	for currentNode != nil {
		if currentNode.Key == k {
			return true
		}
		currentNode = currentNode.Next
	}
	return false
}

// delete will take in a key and delete the node from the bucket
func (b *bucket) delete (k string) {
	if b.Head.Key == k {
		b.Head = b.Head.Next
	}

	previousNode := b.Head
	for previousNode.Next != nil {
		if previousNode.Next.Key == k {
			previousNode.Next = previousNode.Next.Next
			return
		}

		previousNode = previousNode.Next
	}
}

// Init will create a bucket in each slot of the hash table
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.Array {
		result.Array[i] = &bucket{}
	}
	return result
}

// bucketNode structure
type bucketNode struct {
	Key string
	Next *bucketNode
}

// hash function
func hash (key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ARRAY_SIZE
}


func main() {
	hashTable := Init()
	list := []string {
		"ERIC",
		"KENNY",
		"KYLE",
		"STAN",
		"RANDY",
		"BUTTERS",
		"TOKEN",
	}

	for _, v := range list {
		fmt.Println("Inserting", v, "into hash table")
		hashTable.Insert(v)
	}

	fmt.Println("Is ERIC existing in hash table:",hashTable.Search("ERIC"))
	fmt.Println("Is ERIC existing in hash table:",hashTable.Search("ROSE"))
	fmt.Println("Is ERIC existing in hash table:",hashTable.Search("TOKEN"))
	fmt.Println("Is ERIC existing in hash table:",hashTable.Search("RANDY"))


	fmt.Println("Deleting RANDY from hash table")
	hashTable.Delete("RANDY")
	fmt.Println("Is ERIC existing in hash table:",hashTable.Search("RANDY"))
}