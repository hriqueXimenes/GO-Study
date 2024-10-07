package main

import "fmt"

type LinkedList struct {
	Head *Node
}

type Node struct {
	Key      string
	Previous *Node
	Next     *Node
}

func main() {
	list := LinkedList{}
	list.Append("key1")
	list.Append("key2")
	list.Append("key3")

	fmt.Printf("All Keys: %v \n", list.AllKeys())
	list.Delete("key1")

	fmt.Printf("All Keys: %v \n", list.AllKeys())
	fmt.Printf("All Keys Reverse: %v", list.AllKeysReverse())
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
		newNode.Previous = current
		current = current.Next
	}

	newNode.Previous = current
	current.Next = &newNode
}

func (l *LinkedList) Delete(key string) {
	current := l.Head

	if current == nil {
		return // Lista vazia
	}

	if current.Key == key {
		l.Head = current.Next
		if l.Head != nil {
			l.Head.Previous = nil
		}
		return
	}

	for current != nil {
		if current.Key == key {
			if current.Previous != nil {
				current.Previous.Next = current.Next
			}
			if current.Next != nil {
				current.Next.Previous = current.Previous
			}
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

func (l *LinkedList) AllKeysReverse() []string {
	keys := []string{}

	current := l.Head
	if current == nil {
		return []string{}
	}

	for current.Next != nil {
		current = current.Next
	}

	keys = append(keys, current.Key)
	if current.Previous == nil {
		return []string{}
	}

	for current.Previous != nil {
		keys = append(keys, current.Previous.Key)
		current = current.Previous
	}

	return keys
}
