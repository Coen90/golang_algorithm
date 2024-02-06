package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

/*
	여기서는 우선순위 큐로 풀었지만, 힙으로도 풀 수 있다.
	힙으로 풀면 더 빠르다.
*/

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Queue[T Integer] interface {
	Push(x T)
	Pop() T
}

type PriorityQueue[T Integer] struct {
	length int
	elem   []T
}

func (pq *PriorityQueue[T]) Push(x T) {
	pq.length++
	pq.elem = append(pq.elem, x)
	sort.Slice(pq.elem, func(i, j int) bool {
		return pq.elem[i] > pq.elem[j]
	})
}

func (pq *PriorityQueue[T]) Pop() (result T) {
	if pq.length == 0 {
		return
	}
	result = pq.elem[0]
	pq.elem = pq.elem[1:]
	pq.length--
	return
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var pq Queue[int]
	pq = &PriorityQueue[int]{}

	var cnt int
	fmt.Fscanln(r, &cnt)
	for i := 0; i < cnt; i++ {
		var num string
		fmt.Fscanln(r, &num)
		n, _ := strconv.Atoi(num)
		log.Println(n)
		if n == 0 {
			fmt.Fprintln(w, pq.Pop())
			continue
		}
		pq.Push(n)
	}
}
