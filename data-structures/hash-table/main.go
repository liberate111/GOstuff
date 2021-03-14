package main

import "fmt"

// ArraySize is the size of the hash table array
const ArraySize = 10

// HashTable will hold an array
type HashTable struct {
	arr [ArraySize]*bucket
}

// bucket is a linked list in each slot of the hash table array
type bucket struct {
	head *bucketNode
}

// bucketNode is a linked list node that holds the key
type bucketNode struct {
	key  string
	next *bucketNode
}

// Insert will take in a key and add it to the hash table array
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.arr[index].insert(key)
}

// Search will take in a key and return true if that key is stored in the hash table
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.arr[index].search(key)
}

// Delete will take in a key and delete it from the hash table
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.arr[index].delete(key)
}

// insert will take in a key, create a node with the key and insert the node in the bucket
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(k, "already exist")
	}
}

// search will take in a key and return true if the bucket has that key
func (b *bucket) search(k string) bool {
	curNode := b.head
	for curNode != nil {
		if curNode.key == k {
			return true
		}
		curNode = curNode.next
	}
	return false
}

// delete will take in a key and delete the node from bucket
func (b *bucket) delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}

	prevNode := b.head
	for prevNode.next != nil {
		if prevNode.next.key == k {
			//delete
			prevNode.next = prevNode.next.next
			return
		}
		prevNode = prevNode.next
	}
}

// hash
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

// Init will create a bucket in each slot of the hash table
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.arr {
		result.arr[i] = &bucket{}
	}
	return result
}

func main() {
	hashTable := Init()
	list := []string{
		"ERIC",
		"STAN",
		"RANDY",
		"KYLE",
	}

	for _, v := range list {
		hashTable.Insert(v)
	}

	hashTable.Delete("ERIC")
}
