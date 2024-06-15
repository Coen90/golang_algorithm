package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type QueueInterface interface {
	Push(x int)
	Pop() int
	Size() int
	Empty() int
	Front() int
	Back() int
	GetElem() []int
}

type Queue struct {
	Count int
	Elem  []int
}

func (q *Queue) Push(x int) {
	q.Elem = append(q.Elem, x)
	q.Count++
}

func (q *Queue) Pop() int {
	if q.Count == 0 {
		return -1
	}
	q.Count--
	result := q.Elem[0]
	if q.Count == 0 {
		q.Elem = []int{}
	} else {
		q.Elem = q.Elem[1:]
	}
	return result
}

func (q *Queue) Size() int {
	return q.Count
}

func (q *Queue) Empty() int {
	if q.Count == 0 {
		return 1
	}
	return 0
}

func (q *Queue) Front() int {
	if q.Count == 0 {
		return -1
	}
	return q.Elem[0]
}

func (q *Queue) Back() int {
	if q.Count == 0 {
		return -1
	}
	return q.Elem[q.Count-1]
}

func (q *Queue) GetElem() []int {
	return q.Elem
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var q QueueInterface
	q = &Queue{}

	var cnt int
	fmt.Fscanln(r, &cnt)
	for i := 0; i < cnt; i++ {
		var str, number string
		fmt.Fscanln(r, &str, &number)
		s, num := parser(str, number)
		if s == "push" {
			q.Push(num)
			continue
		}
		fmt.Fprintln(w, commander(q, s, num))
	}
}

func parser(str, num string) (string, int) {
	if num != "" {
		number, _ := strconv.Atoi(num)
		return str, number
	}
	return str, -1
}

func commander(q QueueInterface, s string, num int) int {
	switch s {
	case "pop":
		return q.Pop()
	case "size":
		return q.Size()
	case "empty":
		return q.Empty()
	case "front":
		return q.Front()
	case "back":
		return q.Back()
	}
	return -1
}
