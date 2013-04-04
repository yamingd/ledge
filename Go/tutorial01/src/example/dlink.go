package main

import (
	"fmt"
)

type Link struct {
	root *Node
	tail *Node
	size int
}

type Node struct {
	next *Node
	prev *Node
	val  int
}

func (link *Link) Append(val int) {
	tmp := new(Node)
	tmp.next = nil
	tmp.val = val

	if link.root == nil {
		link.root = tmp
	} else {
		link.tail.next = tmp
		tmp.prev = link.tail
	}
	link.size += 1
	link.tail = tmp
}

func (link *Link) Remove(val int) {
	node := link.root
	for node != nil {
		if node.val == val {
			next := node.next
			prev := node.prev
			if prev != nil {
				prev.next = next
			}
			next.prev = prev
			if link.tail == node {
				link.tail = prev
			}
			link.size = link.size - 1
			break
		} else {
			node = node.next
		}
	}
}

func (link *Link) Pop() (val int) {
	node := link.tail
	if node == nil {
		return
	}
	val = node.val
	link.tail = node.prev
	link.tail.next = nil
	node.prev.next = nil
	link.size -= 1
	return
}

func (link *Link) Println() {
	node := link.root
	fmt.Printf("Link Size (%v): ", link.size)
	for node != nil {
		fmt.Printf("%v(%v) --> ", node.val, node.next == nil)
		node = node.next
	}
	fmt.Println()
}

func main() {
	link := new(Link)

	link.Append(10)
	link.Println()

	link.Append(11)
	link.Println()

	link.Append(22)
	link.Println()

	link.Remove(33)
	link.Println()

	fmt.Println("Remove 11")
	link.Remove(11)
	//fmt.Printf("tail is %v, %v \n", link.tail.val, link.tail.next)
	link.Println()

	val := link.Pop()
	//fmt.Printf("tail is %v, %v \n", link.tail.val, link.tail.next)
	fmt.Printf("Pop %v\n", val)
	link.Println()
}
