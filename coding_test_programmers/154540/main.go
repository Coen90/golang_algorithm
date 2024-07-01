package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var dirs = [][]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func main() {
	maps1 := []string{"X591X", "X1X5X", "X231X", "1XXX1"}
	fmt.Println(solution(maps1))
	maps2 := []string{"XXX", "XXX", "XXX"}
	fmt.Println(solution(maps2))
}

func solution(maps []string) []int {
	var result []int
	//방문 map 생성 및 maps []int로 변환
	dayMaps := make([][]int, len(maps))
	visited := make([][]bool, len(maps))
	for i := range maps {
		s := strings.Split(maps[i], "")
		visited[i] = make([]bool, len(s))
		dayMaps[i] = make([]int, len(s))
		for j := range s {
			if s[j] == "X" {
				dayMaps[i][j] = 0
				continue
			}
			tmp, _ := strconv.Atoi(s[j])
			dayMaps[i][j] = tmp
		}
	}

	fmt.Println(dayMaps)

	//maps 순회하며 X 아닌 경우 상하좌우 찾아가기 bfs
	for i := range dayMaps {
		for j := range dayMaps[i] {
			if visited[i][j] {
				continue
			}
			if dayMaps[i][j] != 0 {
				q := &Queue{}
				q.Push([]int{i, j, 0}) //이동좌표, 이동좌표, 누적day
				for !q.IsEmpty() {
					pop := q.Pop()
					nowX, nowY, days := pop[0], pop[1], pop[2]
					if days != 0 && nowX == i && nowY == j {
						result = append(result, days)
						break
					}
					days += dayMaps[i][j]
					for _, dir := range dirs {
						nx := nowX + dir[0]
						ny := nowY + dir[1]
						if IsValid(dayMaps, nx, ny) {
							if i != nx && j != ny {
								q.Push([]int{nx, ny, days + dayMaps[nx][ny]})
								visited[nx][ny] = true
							}
						}
					}
				}
			} else {
				visited[i][j] = true
			}
		}
	}

	//시작점부터 bfs하며 들린 위치 숫자 저장하기
	//시작점 도착하면 총 머물수 있는 시간 list에 저장

	//list sort하기
	if len(result) == 0 {
		return []int{-1}
	}
	sort.Ints(result)
	return result
}

func IsValid(maps [][]int, x, y int) bool {
	if x < 0 || y < 0 || len(maps) >= x || len(maps[0]) >= y {
		return false
	}
	//X는 못감
	if maps[x][y] == 0 {
		return false
	}
	return true
}

type Queue struct {
	size  int
	queue [][]int //[]int{시작x, 시작y, 현재까지 총합}
}

func (q *Queue) Push(arr []int) {
	q.size += 1
	q.queue = append(q.queue, arr)
}

func (q *Queue) Pop() []int {
	var result []int
	if q.IsEmpty() {
		return []int{}
	}
	if q.size == 1 {
		q.size = 0
		result = q.queue[0]
		q.queue = [][]int{}
		return result
	}
	q.size -= 1
	result = q.queue[0]
	q.queue = q.queue[1:]
	return result
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}
