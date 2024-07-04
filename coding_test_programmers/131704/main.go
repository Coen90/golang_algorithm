package main

import "fmt"

func main() {
	order1 := []int{4, 3, 1, 2, 5}
	fmt.Println(solution(order1))
	order2 := []int{5, 4, 3, 2, 1}
	fmt.Println(solution(order2))
}

func solution(order []int) int {
	var result int
	nextNum := 1
	s := &Stack{}
	for _, v := range order {
		//요소가 지금 실을 짐이면
		if v == nextNum {
			nextNum += 1
			result += 1
			continue
		}
		s.add(v)
	}
	for !s.isEmpty() {
		now := s.pop()
		if now == nextNum {
			nextNum += 1
			result += 1
			continue
		}
		break
	}

	return result
}

type Stack struct {
	size int
	arr  []int
}

func (s *Stack) add(i int) {
	s.size += 1
	s.arr = append(s.arr, i)
}

func (s *Stack) pop() int {
	var result int
	if s.size == 0 {
		return 0
	}
	if s.size == 1 {
		result = s.arr[0]
		s.size = 0
		s.arr = []int{}
		return result
	}
	result = s.arr[s.size-1]
	s.arr = s.arr[0 : s.size-1]
	s.size -= 1
	return result
}

func (s *Stack) get() int {
	if s.size == 0 {
		return 0
	}
	return s.arr[s.size-1]
}

func (s *Stack) isEmpty() bool {
	return s.size == 0
}
