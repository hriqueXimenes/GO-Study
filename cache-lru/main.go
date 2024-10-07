package main

import "fmt"

type Cache struct {
	head *Node

	size int
}

type Node struct {
	Key   string
	Value interface{}

	Read int

	Previous *Node
	Next     *Node
}

func main() {
	cache := Cache{
		size: 2,
	}

	cache.Append("key1", "val1")
	cache.Get("key1")

	cache.Append("key2", "val2")
	cache.Append("key3", "val3")

	cache.Append("key3", "val3")
	cache.Append("key4", "val4")
	cache.Append("key5", "val5")

	fmt.Println(cache.GetKeys())
}

func (c *Cache) Reset() {
	c.head = nil
}

func (c *Cache) Append(key, value string) {
	if c.Get(key) == nil && len(c.GetKeys()) >= c.size {
		lru := c.GetLRU()

		if lru != nil {
			if lru.Previous != nil {
				lru.Previous.Next = lru.Next
			} else {
				//head
				c.head = lru.Next
			}

			if lru.Next != nil {
				lru.Next.Previous = lru.Previous
			} else if lru.Previous != nil {
				//tail
				lru.Previous.Next = nil
			}
		}
	}

	newNode := &Node{
		Key:   key,
		Value: value,
	}

	if c.head == nil {
		c.head = newNode
	}

	if c.head.Key == key {
		c.head.Value = value
		return
	}

	current := c.head
	for current.Next != nil {
		if current.Next.Key == key {
			current.Next.Value = value
			return
		}
		current = current.Next
	}

	newNode.Previous = current
	current.Next = newNode
}

func (c *Cache) Get(key string) interface{} {
	current := c.head
	for current != nil {
		if current.Key == key {
			current.Read++
			return current.Value
		}
		current = current.Next
	}

	return nil
}

func (c *Cache) GetLRU() *Node {
	current := c.head

	minRead := c.head.Read
	lru := current

	for current.Next != nil {

		current = current.Next
		if current.Read < minRead {
			lru = current
			minRead = current.Read
		}
	}

	return lru
}

func (c *Cache) GetKeys() []string {
	keys := []string{}

	current := c.head
	for current != nil {
		keys = append(keys, current.Key)
		current = current.Next
	}

	return keys
}
