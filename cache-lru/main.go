package main

import "fmt"

type Cache struct {
	head *Node
	tail *Node

	keys map[string]*Node

	size int
}

type Node struct {
	key   string
	value interface{}

	previous *Node
	next     *Node
}

func main() {
	cache := Cache{
		size: 5,
		keys: map[string]*Node{},
	}

	cache.Append("key1", "val1")
	cache.Append("key2", "val2")
	cache.Append("key3", "val3")
	cache.Append("key4", "val4")
	cache.Append("key5", "val5")

	cache.Get("key1")
	fmt.Println(cache.GetKeys())

	cache.Append("key6", "val6")
	fmt.Println(cache.GetKeys())
}

func (c *Cache) Append(key, value string) {
	if c.Get(key) == nil && len(c.keys) >= c.size {
		lru := c.tail
		if lru != nil {
			delete(c.keys, lru.key)

			if lru.previous != nil {
				lru.previous.next = nil
				c.tail = lru.previous
			} else {
				c.head = nil
				c.tail = nil
			}
		}
	}

	newNode := &Node{
		key:   key,
		value: value,
	}

	if c.head == nil {
		c.head = newNode
		c.tail = newNode

		c.keys[key] = newNode
		return
	}

	if c.head.key == key {
		c.head.value = value
		return
	}

	c.head.previous = newNode
	newNode.next = c.head

	c.head = newNode
	c.keys[key] = newNode
}

func (c *Cache) MoveNodeToHead(node *Node) {
	//if the list is empty
	if c.head == nil {
		return
	}

	//if we have only 1 element
	if c.head.next == nil {
		return
	}

	//if node is already the head of the list
	if node.previous == nil {
		return
	}

	//if node is the tail
	if node.next == nil {
		node.previous.next = nil
		c.tail = node.previous
	}

	node.previous = nil
	node.next = c.head
	c.head.previous = node
	c.head = node
}

func (c *Cache) Get(key string) interface{} {
	if node, found := c.keys[key]; found {
		c.MoveNodeToHead(node)
		return node.value
	}

	return nil
}

func (c *Cache) GetKeys() []string {
	keys := []string{}

	if c.keys != nil {
		for key := range c.keys {
			keys = append(keys, key)
		}
	}

	return keys
}
