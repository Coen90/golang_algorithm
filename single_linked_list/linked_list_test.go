package single_linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPushBack(t *testing.T) {
	var l LinkedList[int]

	assert.Nil(t, l.root)
	assert.Nil(t, l.tail)

	l.PushBack(1)

	assert.NotNil(t, l.root)
	assert.Equal(t, 1, l.Front().Value)

	assert.NotNil(t, l.tail)
	assert.Equal(t, 1, l.Back().Value)

	l.PushBack(2)
	assert.NotNil(t, l.root)
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 2, l.Back().Value)

	l.PushBack(3)
	assert.NotNil(t, l.root)
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 3, l.Back().Value)

	assert.Equal(t, 3, l.Count())
	assert.Equal(t, 3, l.Count2())
}

func TestPushFront(t *testing.T) {
	var l LinkedList[int]

	assert.Nil(t, l.root)
	assert.Nil(t, l.tail)

	l.PushFront(1)

	assert.NotNil(t, l.root)
	assert.Equal(t, 1, l.Front().Value)

	assert.NotNil(t, l.tail)
	assert.Equal(t, 1, l.Back().Value)

	l.PushFront(2)
	assert.NotNil(t, l.root)
	assert.Equal(t, 2, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)

	l.PushFront(3)
	assert.NotNil(t, l.root)
	assert.Equal(t, 3, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)

	assert.Equal(t, 3, l.Count())
	assert.Equal(t, 3, l.Count2())
}

func TestLinkedList_GetAt(t *testing.T) {
	var l LinkedList[int]

	l.PushFront(1)
	assert.Equal(t, 1, l.GetAt(0).Value)

	l.PushFront(2)
	assert.Equal(t, 2, l.GetAt(0).Value)

	l.PushFront(3)
	assert.Equal(t, 3, l.GetAt(0).Value)

	assert.Nil(t, l.GetAt(3))
}

func TestLinkedList_InsertAfter(t *testing.T) {
	var l LinkedList[int]
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	node := l.GetAt(1)
	l.InsertAfter(node, 100)

	assert.Equal(t, 100, l.GetAt(2).Value)

	lastNode := l.GetAt(l.count - 1)
	l.InsertAfter(lastNode, 1000)
	assert.Equal(t, 1000, l.GetAt(l.Count2()-1).Value)

	l.InsertAfter(l.GetAt(0), 1000)
	assert.Equal(t, 1000, l.GetAt(1).Value)

	l.InsertAfter(l.GetAt(1000), 1000)
}

func TestLinkedList_InsertBefore(t *testing.T) {
	var l LinkedList[int]
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	node := l.GetAt(1)
	l.InsertBefore(node, 100)

	assert.Equal(t, 4, l.Count())
	assert.Equal(t, 4, l.Count2())
	assert.Equal(t, 100, l.GetAt(1).Value)

	l.InsertBefore(l.GetAt(1000), 1000)

	l.InsertBefore(l.GetAt(0), 1000)
	assert.Equal(t, 1000, l.GetAt(0).Value)

	l.InsertBefore(&Node[int]{Value: 9999}, 9999)
}

func TestLinkedList_PopFront(t *testing.T) {
	var l LinkedList[int]
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	l.PopFront()

	assert.Equal(t, 2, l.Count2())
	assert.Equal(t, 2, l.Front().Value)
	assert.Equal(t, 3, l.GetAt(1).Value)

}

func TestLinkedList_Remove(t *testing.T) {
	var l LinkedList[int]
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	l.Remove(l.GetAt(1))
	assert.Equal(t, 2, l.Count())
	assert.Equal(t, 2, l.Count2())
	assert.Equal(t, 1, l.GetAt(0).Value)
	assert.Equal(t, 3, l.GetAt(1).Value)

	l.Remove(&Node[int]{Value: 1000})

	l.Remove(l.GetAt(0))
	l.Remove(l.GetAt(0))

	assert.Nil(t, l.GetAt(0))
}
