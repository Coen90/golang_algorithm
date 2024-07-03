package main

import (
	"fmt"
	"sync"
)

func main() {
	arr1 := []int{1, 2, 1, 3, 1, 4, 1, 2}
	fmt.Println(solution(arr1))
	arr2 := []int{1, 2, 3, 1, 4}
	fmt.Println(solution(arr2))
}

// first solution : time out
func solution2(topping []int) int {
	var result int
	var wg sync.WaitGroup
	var mutex sync.Mutex
	if len(topping) == 1 {
		return 0
	}
	for i := 1; i < len(topping)-1; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			oMap := make(map[int]bool)
			yMap := make(map[int]bool)
			olderTopping := topping[0:i]
			youngerTopping := topping[i:]
			for j := range olderTopping {
				oMap[olderTopping[j]] = true
			}
			for j := range youngerTopping {
				yMap[youngerTopping[j]] = true
			}
			oCount := 0
			for _, v := range oMap {
				if v {
					oCount += 1
				}
			}
			yCount := 0
			for _, v := range yMap {
				if v {
					yCount += 1
				}
			}
			if oCount == yCount {
				mutex.Lock()
				result += 1
				mutex.Unlock()
			}
		}(i)
	}
	wg.Wait()

	return result
}

// 두번째 풀이. 이건 시간 괜찮음
func solution(topping []int) int {
	var result int

	older := make(map[int]int)
	younger := make(map[int]int)

	for i := range topping {
		younger[topping[i]] = younger[topping[i]] + 1
	}

	for i := range topping {
		older[topping[i]] = older[topping[i]] + 1
		if younger[topping[i]]-1 == 0 {
			delete(younger, topping[i])
		} else {
			younger[topping[i]] = younger[topping[i]] - 1
		}
		if len(older) == len(younger) {
			result += 1
		}
	}

	return result
}
