package main

import "fmt"

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

func (q *IntQueue[T]) Dequeue() T {
	defer func() {
		q.queue = q.queue[1:]
	}()
	return q.queue[0]
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
	fmt.Println(solution(testCase1, k1, ax1, ay1))
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
	var queue IntQueue[[]int]
	queue.Enqueue([]int{ay, ax, 0})
	for !queue.IsEmpty() {
		//내일하자~
	}

	return 1
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
