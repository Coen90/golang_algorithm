package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

var miro [][]string
var visited [][]bool
var min = math.MaxInt
var row, col int
var dir = [][]int{
	[]int{1, 0},  //d
	[]int{0, -1}, //l
	[]int{-1, 0}, //u
	[]int{0, 1},  //r
}

func main() {
	start := time.Now().UTC()
	r := bufio.NewReader(os.Stdin)
	fmt.Fscanln(r, &row, &col)
	for i := 0; i < row; i++ {
		var tmp string
		fmt.Fscanln(r, &tmp)
		split := strings.Split(tmp, "")
		miro = append(miro, split)
		visited = append(visited, make([]bool, col))
	}
	BFS(0, 0, 1)
	fmt.Print(min)
	end := time.Now().UTC()
	fmt.Println(start)
	fmt.Println(end)
}

func BFS(y, x, cnt int) {
	if cnt > min {
		return
	}
	if y == row-1 && x == col-1 {
		min = cnt
		return
	}
	visited[y][x] = true
	for _, v := range dir {
		ny, nx := v[1], v[0]
		if ny+y < 0 || ny+y >= row || nx+x < 0 || nx+x >= col {
			continue
		}
		if miro[ny+y][nx+x] == "0" {
			continue
		}
		if visited[ny+y][nx+x] {
			continue
		}
		BFS(ny+y, nx+x, cnt+1)
		visited[y][x] = false
	}
	return
}
