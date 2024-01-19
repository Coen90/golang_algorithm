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
		return
	}
	if l.now.prev == nil {
		l.now.prev = node
		l.now.next = node
		node.next = l.now
		node.prev = l.now
		return
	}
	prev := l.now.prev
	prev.next = node
	node.prev = prev
	node.next = l.now
	l.now.prev = node
}

func (l *LinkedList[T]) Pop(n *Node[T]) {
	l.count--
	if l.count == 0 {
		l.now = nil
		return
	} else if l.count == 1 {
		l.now = n.next
		l.now.prev = nil
		l.now.next = nil
		return
	}
	prev := n.prev
	next := n.next
	prev.next = next
	next.prev = prev
	l.now = next
}

func main() { //10 3 \b 1 2 3
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var length, cnt, result int
	fmt.Fscan(r, &length, &cnt)

	popList := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		fmt.Fscan(r, &popList[i])
	}

	list := &LinkedList[int]{}
	for i := 1; i <= length; i++ {
		list.Push(i)
	}

	for _, v := range popList {
		result += Counter[int](list, v)
	}
	fmt.Fprintln(w, result)
}

func Counter[T comparable](list *LinkedList[T], target T) int {
	now := list.now
	if now.Value == target {
		list.Pop(now)
		return 0
	}
	var nextCount, prevCount int
	var node = now
	for ; node.Value != target; node = node.next {
		nextCount++
	}
	result := nextCount
	node = now
	for ; node.Value != target; node = node.prev {
		prevCount++
	}
	if result > prevCount {
		result = prevCount
	}
	list.Pop(node)
	return result
}
