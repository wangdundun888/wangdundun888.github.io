package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getRandomSlice(size int) []int {
	if size < 0 {
		return nil
	}
	s := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i, _ := range s {
		s[i] = rand.Intn(COUNT)
	}
	return s
}

const COUNT int = 1000

//非比较排序算法
//稳定排序
//总的时间代价是O(k+n),当k=O(n)时,总的代价就是O(n)
func countingSort(A, result []int, max int) {
	c := make([]int, max)
	for _, v := range A {
		c[v]++
	}
	for i := 1; i < max; i++ {
		c[i] += c[i-1]
	}
	for i := len(A) - 1; i >= 0; i-- {
		result[c[A[i]]-1] = A[i]
		c[A[i]]--
	}
}

func main() {
	cnt := 10
	a := getRandomSlice(cnt)
	fmt.Println(a)
	result := make([]int, len(a))
	countingSort(a, result, COUNT)
	fmt.Println(result)
}
