package main

import "fmt"

func main() {
	arr := []int{5, 1, 3, 2, 5, 2, 6, 8, 2, 0, 4, 5, 1, 6, 8, 2, 7, 9, 2, 1, 5, 6, 10}

	var count [11]int
	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
	}
	fmt.Println(count)

	// 방법1
	sorted := make([]int, 0, len(arr))
	for i := 0; i < len(count); i++ {
		for j := 0; j < count[i]; j++ {
			sorted = append(sorted, i)
		}
	}
	fmt.Println(sorted)

	// 방법2
	for i := 1; i < 11; i++ {
		count[i] += count[i-1]
	}

	fmt.Println(count)

	sorted2 := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		sorted2[count[arr[i]]-1] = arr[i]
		count[arr[i]]--
	}
	fmt.Println(sorted2)

	test02()
	test03()
}

func test02() {
	// 알파벳 소문자로 이루어진 문자열 중 가장 많이 나오는 알파벳 캐릭터를 출력하시오
	str := "ashfjkwlejhguiblkcjvlkzxcjviwjfkqwljgijaksldfjlaksdfjkasfjsdkfjeifjalksj"

	var count [26]int
	for i := 0; i < len(str); i++ {
		count[str[i]-'a']++
	}

	maxCount := 0
	var maxCh byte
	for i := 0; i < 26; i++ {
		if count[i] > maxCount {
			maxCount = count[i]
			maxCh = byte('a' + i)
		}
	}
	fmt.Printf("Max character : %c count: %d", maxCh, maxCount)
}

func test03() {
	// 한 반의 학생들 중 키가 특정 범위의 학생들만 출력하시오.
	// 범위 값은 여러번 입력받을 수 & 있습니다
	// 키는 소수점 1자리까지 주어집니다.
	students := []struct {
		Name   string
		Height float64
	}{
		{Name: "Kyle", Height: 173.4},
		{Name: "Ken", Height: 164.5},
		{Name: "Ryu", Height: 178.8},
		{Name: "Honda", Height: 154.2},
		{Name: "Hwarang", Height: 188.8},
		{Name: "Lebron", Height: 209.8},
		{Name: "Hodong", Height: 197.7},
		{Name: "Tom", Height: 164.8},
		{Name: "Kevin", Height: 164.8},
	}

	for i := 0; i < len(students); i++ {
		if students[i].Height >= 170.0 && students[i].Height < 180.0 {
			fmt.Println(students[i].Name, students[i].Height)
		}
	}

	var heightMap [3000][]string
	for i := 0; i < len(students); i++ {
		idx := int(students[i].Height * 10)
		heightMap[idx] = append(heightMap[idx], students[i].Name)
	}

	// 140 ~ 170
	for i := 1400; i < 1700; i++ {
		for _, name := range heightMap[i] {
			fmt.Println("name:", name, "height:", float64(i)/10)
		}
	}

	// 180 ~ 210
	for i := 1800; i < 2100; i++ {
		for _, name := range heightMap[i] {
			fmt.Println("name:", name, "height:", float64(i)/10)
		}
	}
}
