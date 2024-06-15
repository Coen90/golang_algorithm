package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	//w := bufio.NewWriter(os.Stdout)
	//defer func() {
	//	if err := w.Flush(); err != nil {
	//		panic(err)
	//	}
	//}()

	var nodeCnt, edgeCnt, start int
	_, err := fmt.Fscanln(r, &nodeCnt, &edgeCnt, &start)
	if err != nil {
		panic(err)
	}

	graph := make(map[int]*TreeNode[int], nodeCnt)
	for i := 0; i < edgeCnt; i++ {
		var parent, child int
		fmt.Fscanln(r, &parent, &child)
		parentNode := &TreeNode[int]{Value: parent}
		if _, ok := graph[parent]; !ok {
			graph[parent] = parentNode
		}
		if _, ok := graph[child]; !ok {
			childNode := &TreeNode[int]{Value: child}
			graph[child] = childNode
		}
		graph[parent].Children = append(graph[parent].Children, graph[child])
		graph[child].Children = append(graph[child].Children, graph[parent])
	}
	if graph[start] != nil {
		for k, _ := range graph {
			sort.Slice(graph[k].Children, func(i, j int) bool {
				return graph[k].Children[i].Value > graph[k].Children[j].Value
			})
		}
		graph[start].DFSByStack(nodeCnt)
	} else {
		fmt.Println(start)
	}
	if graph[start] != nil {
		fmt.Println()
		for k, _ := range graph {
			sort.Slice(graph[k].Children, func(i, j int) bool {
				return graph[k].Children[i].Value < graph[k].Children[j].Value
			})
		}
		graph[start].BFS(nodeCnt)
	} else {
		fmt.Println(start)
	}
}

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type TreeNode[T Integer] struct {
	Value    T
	Children []*TreeNode[T]
}

func (t *TreeNode[T]) Add(val T) *TreeNode[T] {
	t.Children = append(t.Children, &TreeNode[T]{Value: val})
	return t
}

func (t *TreeNode[T]) BFS(nodeCnt int) {
	visited := make(map[T]bool, nodeCnt)
	queue := make([]*TreeNode[T], 0)
	queue = append(queue, t)
	for len(queue) > 0 {
		target := queue[0]
		queue = queue[1:]
		if v, ok := visited[target.Value]; ok && v {
			continue
		}
		visited[target.Value] = true
		for _, v := range target.Children {
			if b := visited[v.Value]; !b {
				queue = append(queue, v)
			}
		}
		fmt.Print(target.Value, " ")
	}
}

func (t *TreeNode[T]) DFSByStack(nodeCnt int) {
	visited := make(map[T]bool, nodeCnt)
	stack := make([]*TreeNode[T], 0)
	stack = append(stack, t)
	for len(stack) > 0 {
		target := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if v, ok := visited[target.Value]; ok && v {
			continue
		}
		visited[target.Value] = true
		for _, child := range target.Children {
			if !visited[child.Value] {
				stack = append(stack, child)
			}
		}
		fmt.Print(target.Value, " ")
	}
}
