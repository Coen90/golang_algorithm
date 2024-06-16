package main

import (
	"errors"
	"fmt"
	"log"
	"sort"
)

var dirs = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

type IntQueue[T any] struct {
	size  int
	queue []T
}

func (q *IntQueue[T]) IsEmpty() bool {
	return q.size == 0
}

func (q *IntQueue[T]) Enqueue(slice T) {
	q.queue = append(q.queue, slice)
	q.size += 1
}

func (q *IntQueue[T]) Dequeue() (T, error) {
	if q.size == 0 {
		var zeroValue T
		return zeroValue, errors.New("Empty queue")
	} else if q.size == 1 {
		defer func() {
			q.size = 0
			q.queue = []T{}
		}()
		return q.queue[0], nil
	}
	defer func() {
		q.size -= -1
		q.queue = q.queue[1:]
	}()
	return q.queue[0], nil
}

func main() {
	testCase1 := [][]int{
		{0, 0, 1, 0, 0, 0},
		{0, 2, 0, 0, 0, 1},
		{0, 0, 2, 1, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0},
	}
	k1 := 2
	ax1 := 1
	ay1 := 2
	testCase2 := [][]int{
		{0, 0, 0, 1},
		{0, 2, 0, 1},
		{2, 0, 0, 1},
		{0, 2, 0, 1},
	}
	k2 := 2
	ax2 := 2
	ay2 := 1
	fmt.Println(solution(testCase1, k1, ax1, ay1))
	fmt.Println(solution(testCase2, k2, ax2, ay2))
}

func solution(board [][]int, k, ax, ay int) int {
	boardSize := len(board)

	// 폭발 범위 구할 slice 생성
	explosion := make([][]bool, boardSize)
	for i := 0; i < boardSize; i++ {
		explosion[i] = make([]bool, boardSize)
	}

	//보드 순회하며 폭탄이 있는 경우 폭발범위 보드 만들기
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			// 폭탄인 경우
			if board[i][j] == 1 {
				explosion[i][j] = true
				//폭탄 반경 생성
				for _, dir := range dirs {
					for l := 1; l <= k; l++ {
						nx := i + dir[0]*l
						ny := j + dir[1]*l
						//보드를 초과하지 않고 벽이 아니면 폭발범위
						if isValid(boardSize, nx, ny) && !isWall(board, nx, ny) {
							explosion[nx][ny] = true
							continue
						}
						//아니면 break
						break
					}
				}
			}
		}
	}

	//BFS
	var results []int
	var queue IntQueue[[]int]
	queue.Enqueue([]int{ax, ay, 0})
	for !queue.IsEmpty() {
		//현재 위치 확인
		current, err := queue.Dequeue()
		if err != nil {
			log.Fatal(err)
		}
		x := current[0]
		y := current[1]
		cnt := current[2]
		//현재 위치가 폭탄 영향범위가 아니면 종료 및 list 추가
		if !explosion[x][y] {
			results = append(results, cnt)
			break
		}
		//방향별로 bfs 시작
		for _, dir := range dirs {
			nx := x + dir[0]
			ny := y + dir[1]
			//이동 가능한 방향이면
			if isValid(boardSize, nx, ny) && !isWall(board, nx, ny) && !isBomb(board, nx, ny) {
				queue.Enqueue([]int{nx, ny, cnt + 1})
			}
		}
	}

	var result int
	sort.Ints(results)

	if len(results) == 0 {
		result = -1
	} else {
		result = results[0]
	}
	return result
}

func isValid(boardSize, x, y int) bool {
	return x >= 0 && x < boardSize && y >= 0 && y < boardSize
}

func isWall(board [][]int, x, y int) bool {
	return board[x][y] == 2
}

func isBomb(board [][]int, x, y int) bool {
	return board[x][y] == 1
}

func isDestination(board [][]int, x, y int) bool {
	return board[x][y] == 1
}
