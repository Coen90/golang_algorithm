package double_linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushBack(t *testing.T) {
	var l LinkedList[int]

	assert.Nil(t, l.root)
	assert.Nil(t, l.tail)

	l.PushBack(1)
	assert.NotNil(t, l.root)
	assert.Equal(t, 1, l.Front().Value)
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

	assert.Nil(t, l.GetAt(5))

	assert.Equal(t, 3, l.GetAt(2).Value)

	l.PushFront(100)
	assert.Equal(t, 100, l.GetAt(0).Value)
	assert.Equal(t, 100, l.Front().Value)
	assert.Equal(t, 4, l.Count())

	node := l.GetAt(1)
	l.InsertAfter(node, 4)
	assert.Equal(t, 4, l.GetAt(2).Value)
	assert.Equal(t, 1, l.GetAt(1).Value)
	assert.Equal(t, 2, l.GetAt(3).Value)
	assert.Equal(t, 5, l.Count())
	assert.Equal(t, 1, node.Value)
	assert.Equal(t, 4, node.next.Value)
	assert.Equal(t, 2, node.next.next.Value)
	assert.Equal(t, 4, node.next.next.prev.Value)

	notIncludedNode := &Node[int]{
		Value: 1,
	}
	l.InsertAfter(notIncludedNode, 1000000)
	assert.Equal(t, 5, l.Count())

	l.InsertBefore(node, 5)
	assert.Equal(t, 5, l.GetAt(1).Value)
	assert.Equal(t, 1, l.GetAt(2).Value)
	assert.Equal(t, 2, l.GetAt(4).Value)
	assert.Equal(t, 6, l.Count())
	assert.Equal(t, 5, node.prev.Value)
	assert.Equal(t, 1, node.prev.next.Value)

	l.InsertAfter(l.Back(), 10)
	assert.Equal(t, 10, l.tail.Value)

	l.InsertBefore(l.Front(), 1000)
	assert.Equal(t, 1000, l.root.Value)
	assert.Equal(t, 1000, l.Front().Value)

	n := l.PopFront()
	assert.Equal(t, 1000, n.Value)
	assert.Equal(t, 100, l.Front().Value)

	l.Remove(l.GetAt(0))
	assert.Equal(t, 5, l.Front().Value)
}
