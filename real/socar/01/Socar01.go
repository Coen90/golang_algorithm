package main

import (
	"fmt"
	"strconv"
	"strings"
)

func solution(tickets []string, roll int, shop [][]string, money int) int {
	// 티켓 정보 파싱
	ticketPrices := make(map[string]int)
	for _, ticket := range tickets {
		parts := strings.Split(ticket, " ")
		name := parts[0]
		price, _ := strconv.Atoi(parts[1])
		ticketPrices[name] = price
	}

	// 최대 상점 방문 횟수
	maxShops := len(shop)

	// DP 테이블 초기화: dp[i][j]는 i번째 상점까지 방문하고 j만큼의 돈을 사용했을 때 얻을 수 있는 최대 황금 티켓 개수
	// dp[i][j]는 [각 티켓 보유 개수]를 저장하는 map으로 구성됩니다.
	dp := make([][]map[string]int, maxShops+1)
	for i := range dp {
		dp[i] = make([]map[string]int, money+1)
		for j := range dp[i] {
			dp[i][j] = make(map[string]int)
		}
	}

	// 다이나믹 프로그래밍
	for i := 1; i <= maxShops; i++ {
		for j := 0; j <= money; j++ {
			// 현재 상점에서 티켓을 구매하지 않는 경우
			dp[i][j] = copyMap(dp[i-1][j])

			// 현재 상점에서 티켓을 구매하는 경우
			for _, ticket := range shop[i-1] {
				price := ticketPrices[ticket]
				if j >= price {
					// 이전 상태에서 현재 티켓을 하나 추가하여 새로운 상태 생성
					newTickets := copyMap(dp[i][j-price])
					newTickets[ticket]++
					dp[i][j] = maxTickets(dp[i][j], newTickets)
				}
			}

			// 현재 상점에서 새로고침하는 경우
			if i < maxShops && j+roll <= money {
				dp[i+1][j+roll] = maxTickets(dp[i+1][j+roll], dp[i][j])
			}
		}
	}

	// 최대 황금 티켓 개수 계산
	maxGoldenTickets := 0
	for j := 0; j <= money; j++ {
		goldenTickets := calculateGoldenTickets(dp[maxShops][j])
		maxGoldenTickets = max(maxGoldenTickets, goldenTickets)
	}

	return maxGoldenTickets
}

func copyMap(m map[string]int) map[string]int {
	newMap := make(map[string]int)
	for k, v := range m {
		newMap[k] = v
	}
	return newMap
}

func maxTickets(a, b map[string]int) map[string]int {
	aGolden := calculateGoldenTickets(a)
	bGolden := calculateGoldenTickets(b)

	if aGolden > bGolden {
		return a
	} else if aGolden < bGolden {
		return b
	} else {
		// 황금 티켓 개수가 같으면 티켓 개수 합 비교
		aSum := sumTickets(a)
		bSum := sumTickets(b)
		if aSum >= bSum {
			return a
		} else {
			return b
		}
	}
}

func calculateGoldenTickets(tickets map[string]int) int {
	goldenTickets := 0
	for _, count := range tickets {
		goldenTickets += count / 3
	}
	return goldenTickets
}

func sumTickets(tickets map[string]int) int {
	sum := 0
	for _, count := range tickets {
		sum += count
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	tickets := []string{"A 1", "B 2", "C 5", "D 3"}
	roll := 10
	shop := [][]string{{"B", "C", "B", "C"}, {"A", "A", "A", "B"}, {"D", "D", "C", "D"}}
	money := 30
	fmt.Println(solution(tickets, roll, shop, money)) // 출력: 2
}
