package main

import "fmt"

type LinkedList struct {
	Head *Node
}

type Node struct {
	Key  string
	Next *Node
}

func main() {
	list := LinkedList{}
	list.Append("key1")
	list.Append("key2")
	list.Append("key3")

	fmt.Printf("All Keys: %v \n", list.AllKeys())
	list.Delete("key1")
	fmt.Printf("All Keys After Delete: %v", list.AllKeys())
}

func (l *LinkedList) Append(key string) {
	newNode := Node{
		Key: key,
	}

	if l.Head == nil {
		l.Head = &newNode
		return
	}

	current := l.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = &newNode
}

func (l *LinkedList) Delete(key string) {
	current := l.Head

	if current == nil {
		return
	}

	//If the node has the key.
	if current.Key == key {
		l.Head = current.Next
		return
	}

	//Check all nodes
	for current.Next != nil {
		nextNode := current.Next
		if nextNode != nil && nextNode.Key == key {
			current.Next = nextNode.Next
			return
		}

		current = current.Next
	}
}

func (l *LinkedList) AllKeys() []string {
	keys := []string{}

	current := l.Head
	if current == nil {
		return []string{}
	}

	keys = append(keys, current.Key)

	for current.Next != nil {
		keys = append(keys, current.Next.Key)
		current = current.Next
	}

	return keys
}
