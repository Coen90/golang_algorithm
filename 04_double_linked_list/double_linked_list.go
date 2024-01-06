package double_linked_list

type Node[T any] struct {
	next *Node[T]
	prev *Node[T]

	Value T
}

type LinkedList[T any] struct {
	root  *Node[T]
	tail  *Node[T]
	count int
}

func (l *LinkedList[T]) Front() *Node[T] {
	return l.root
}

func (l *LinkedList[T]) Back() *Node[T] {
	return l.tail
}

func (l *LinkedList[T]) PushBack(val T) {
	n := &Node[T]{
		Value: val,
	}
	if l.tail == nil {
		l.root = n
		l.tail = n
		l.count = 1
		return
	}

	l.tail.next = n
	n.prev = l.tail
	l.tail = n
	l.count++
}

func (l *LinkedList[T]) Count() int {
	return l.count
}

func (l *LinkedList[T]) GetAt(idx int) *Node[T] {
	if l.count <= idx || idx < 0 {
		return nil
	}
	n := l.root
	cnt := 0
	for ; n != nil; n = n.next {
		if cnt == idx {
			return n
		}
		cnt++
	}
	return nil
}

func (l *LinkedList[T]) PushFront(val T) {
	n := &Node[T]{
		Value: val,
	}
	if l.tail == nil {
		l.root = n
		l.tail = n
		l.count = 1
		return
	}
	l.root.prev = n
	n.next = l.root
	l.root = n
	l.count++
}

func (l *LinkedList[T]) InsertAfter(node *Node[T], val T) {
	if !l.IsExists(node) {
		return
	}
	n := &Node[T]{
		Value: val,
	}

	n.prev = node
	if node.next == nil {
		l.tail = n
	}
	n.next = node.next
	if node.next != nil {
		node.next.prev = n
	}
	node.next = n
	l.count++
}

func (l *LinkedList[T]) IsExists(node *Node[T]) bool {
	n := l.root
	for ; n != nil; n = n.next {
		if n == node {
			return true
		}
	}
	return false
}

func (l *LinkedList[T]) InsertBefore(node *Node[T], val T) {
	if !l.IsExists(node) {
		return
	}
	n := &Node[T]{
		Value: val,
	}
	if node.prev != nil {
		node.prev.next = n
	}
	n.next = node
	n.prev = node.prev

	if node.prev == nil {
		l.root = n
	}
	node.prev = n
	l.count++
}
