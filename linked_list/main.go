package main

import "fmt"

type Node[T any] struct {
	next *Node[T]
	//prev *Node[T] //이게 있으면 Double Linked List 없으면 Single Linked List
	val T
}

func Append[T any](node *Node[T], next *Node[T]) {
	node.next = next
}

func main() {
	root := &Node[int]{nil, 10}
	root.next = &Node[int]{nil, 20}
	root.next.next = &Node[int]{nil, 30}

	for n := root; n != nil; n = n.next {
		fmt.Printf("node val : %d\n", n.val)
	}

	fmt.Println()
	fmt.Println()

	node := &Node[int]{nil, 100}
	originalNext := root.next.next
	Append(root.next, node)
	node.next = originalNext

	for n := root; n != nil; n = n.next {
		fmt.Printf("node val : %d\n", n.val)
	}
}
