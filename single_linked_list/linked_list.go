package single_linked_list

type Node[T any] struct {
	next  *Node[T]
	Value T
}

type LinkedList[T any] struct {
	root  *Node[T]
	tail  *Node[T]
	count int
}

func (l *LinkedList[T]) PushBack(value T) {
	node := &Node[T]{
		Value: value,
	}
	l.count++
	if l.root == nil {
		l.root = node
		l.tail = node
		return
	}

	l.tail.next = node
	l.tail = node
}

func (l *LinkedList[T]) PushFront(value T) {
	node := &Node[T]{
		Value: value,
	}
	l.count++
	if l.root == nil {
		l.root = node
		l.tail = node
		return
	}
	node.next = l.root
	l.root = node
}

func (l *LinkedList[T]) Front() *Node[T] {
	return l.root
}

func (l *LinkedList[T]) Back() *Node[T] {
	return l.tail
}

// Deprecated
func (l *LinkedList[T]) Count() int {
	node := l.root
	cnt := 0

	for ; node != nil; node = node.next {
		cnt++
	}
	return cnt
}

func (l *LinkedList[T]) Count2() int {
	node := l.root
	cnt := 0

	for ; node != nil; node = node.next {
		cnt++
	}
	return cnt
}

func (l *LinkedList[T]) GetAt(idx int) *Node[T] {
	if l.count <= idx {
		return nil
	}
	node := l.root
	cnt := 0
	for ; node != nil; node = node.next {
		if cnt == idx {
			return node
		}
		cnt++
	}
	return nil
}

func (l *LinkedList[T]) InsertAfter(node *Node[T], value T) {
	if node == nil {
		return
	}
	if node == l.tail {
		l.PushBack(value)
	}
	var prevNode *Node[T]
	if node == l.root {
		prevNode = node
	} else {
		prevNode = l.findPrevNode(node)
		if prevNode == nil {
			return
		}
	}

	newNode := &Node[T]{
		Value: value,
	}

	node.next, newNode.next = newNode, node.next
	l.count++
}

func (l *LinkedList[T]) InsertBefore(node *Node[T], value T) {
	if node == nil {
		return
	}
	if node == l.root {
		l.PushFront(value)
		return
	}
	prevNode := l.findPrevNode(node)
	newNode := &Node[T]{
		Value: value,
	}
	if prevNode == nil {
		return
	}
	prevNode.next, newNode.next = newNode, node
	l.count++
}

func (l *LinkedList[T]) findPrevNode(node *Node[T]) *Node[T] {
	inner := l.root
	for ; inner != nil; inner = inner.next {
		if inner.next == node {
			return inner
		}
	}
	return nil
}

func (l *LinkedList[T]) PopFront() {
	if l.root == nil {
		return
	}
	l.root.next, l.root = nil, l.root.next
	if l.root == nil {
		l.tail = nil
	}
	l.count--
}

func (l *LinkedList[T]) Remove(node *Node[T]) {
	if node == l.root {
		l.PopFront()
		return
	}
	prev := l.findPrevNode(node)
	if prev == nil {
		return
	}
	prev.next = node.next
	node.next = nil
	if node == l.tail {
		l.tail = prev
	}
	l.count--
}
