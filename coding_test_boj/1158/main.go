package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node[T any] struct {
	Value T
	prev  *Node[T]
	next  *Node[T]
}

type LinkedList[T any] struct {
	now   *Node[T]
	count int
}

func (l *LinkedList[T]) Push(val T) {
	node := &Node[T]{Value: val}
	l.count++
	if l.now == nil {
		l.now = node
		l.now.prev = node
		l.now.next = node
		return
	}
	now := l.now
	node.next = now
	node.prev = now.prev
	now.prev.next = node
	now.prev = node
}

func (l *LinkedList[T]) SetNow(n *Node[T]) {
	l.now = n
}

func (l *LinkedList[T]) Pop(n *Node[T]) *Node[T] {
	l.count--
	if l.count == 0 {
		l.now = nil
		return n
	}
	if l.count == 1 {
		l.now = n.next
		n.prev = n.next
		n.next = n.next
		return n
	}
	l.now = n.next
	n.prev.next = n.next
	n.next.prev = n.prev
	return n
}

func main() { // 7 3  or 10 5
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var length, cnt, counter int
	fmt.Fscan(r, &length, &cnt)

	list := &LinkedList[int]{}
	for i := 1; i <= length; i++ {
		list.Push(i)
	}
	node := list.now
	fmt.Fprint(w, "<")
	for ; list.now != nil; node = node.next {
		counter++
		if counter%cnt == 0 {
			list.SetNow(node)
			pop := list.Pop(node)
			fmt.Fprint(w, pop.Value)
			counter = 0
			if list.now != nil {
				fmt.Fprint(w, ", ")
			}
		}
	}
	fmt.Fprint(w, ">")
}
