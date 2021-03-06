package main

import (
	"algorithm/internal/basicsort"
	"fmt"
	"math/rand"
	"time"
)

/**
# 선택 정렬

- 다음과 같은 순서를 반복하며 정렬하는 알고리즘
   	1. 주어진 데이터 중, 최소값을 찾음
   	2. 해당 최소값을 데이터 맨 앞에 위치하 값과 교체함
	3. 맨 앞의 위치를 뺀 나머지 데이터를 동일한 방법으로 반복함

## 링크
https://visualgo.net/en/sorting
*/
func main() {
	rand.Seed(time.Now().UnixNano())

	var data []int
	for range [50]int{} {
		data = append(data, rand.Intn(100))
	}

	fmt.Println("before: ", data)
	result := basicsort.SelectionSort(data)
	fmt.Println("after : ", result)
}
