package BufferPool

import (
	"errors"
	"fmt"
)

var (
	ErrorMaxCapacity = errors.New("Maximum Capacity Reached.")
)

type node struct {
	key   any
	value any
	prev  *node
	next  *node
}

type circularList struct {
	head     *node
	tail     *node
	size     int
	capacity int
}

// traverse and find the key
func (c *circularList) find(key any) *node {
	ptr := c.head
	for i := 0; i < c.size; i++ {
		if ptr.key == key {
			return ptr
		}

		ptr = ptr.next
	}

	return nil
}

// checks if the list contains key
func (c *circularList) hasKey(key any) bool {
	return c.find(key) != nil
}

// inserts the key value pair into the circular list
func (c *circularList) insert(key any, value any) error {
	if c.size == c.capacity {
		return ErrorMaxCapacity
	}

	newNode := &node{key, value, nil, nil}
	if c.size == 0 {
		newNode.next = newNode
		newNode.prev = newNode
		c.head = newNode
		c.tail = newNode
		c.size++
		return nil
	}

	node := c.find(key)
	if node != nil {
		node.value = value
		return nil
	}

	newNode.next = c.head
	newNode.prev = c.tail

	c.tail.next = newNode
	if c.head == c.tail {
		c.head.next = newNode
	}

	c.tail = newNode
	c.head.prev = c.tail

	c.size++

	return nil
}

// finds and replaces the key
func (c *circularList) remove(key any) {
	node := c.find(key)
	if node == nil {
		return
	}

	if c.size == 1 {
		c.head = nil
		c.tail = nil
		c.size--
		return
	}

	if node == c.head {
		c.head = c.head.next
	}

	if node == c.tail {
		c.tail = c.tail.prev
	}

	node.next.prev = node.prev
	node.prev.next = node.next

	c.size--
}

// checks if the list is full
func (c *circularList) isFull() bool {
	return c.size == c.capacity
}

func (c *circularList) print() {
	if c.size == 0 {
		fmt.Println(nil)
	}
	ptr := c.head
	for i := 0; i < c.size; i++ {
		fmt.Println(ptr.key, ptr.value, ptr.prev.key, ptr.next.key)
		ptr = ptr.next
	}
}

func newCircularList(maxSize int) *circularList {
	return &circularList{nil, nil, 0, maxSize}
}
