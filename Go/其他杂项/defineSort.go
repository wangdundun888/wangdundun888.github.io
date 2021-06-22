package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

//自定义排序二维数组

//返回一个n行2列的数组
func getArrayOfarray(n int) [][]int {
	a := make([][]int, n)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		num1 := rand.Intn(100)
		num2 := rand.Intn(100)
		subset := []int{num1, num2}
		a[i] = subset
	}
	return a
}

//方式2:实现len,swap,less三个方法
type A struct {
	array [][]int
}

func (a *A) Len() int { return len(a.array) }
func (a *A) Swap(i, j int) {
	a.array[i], a.array[j] = a.array[j], a.array[i]
}
func (a *A) Less(i, j int) bool {
	if a.array[i][0] < a.array[j][0] {
		return true
	}
	return false
}

func main() {
	//方式1 调用sort.Slice(slice interface{}, less func(i, j int) bool)
	a := getArrayOfarray(10)
	fmt.Println(a)
	sort.Slice(a, func(i, j int) bool {
		if a[i][0] < a[j][0] {
			return true
		}
		return false
	})
	fmt.Println(a)
	//方式2:实现Sort接口
	a1 := getArrayOfarray(10)
	fmt.Println(a1)
	a2 := &A{a1}
	sort.Sort(a2)
	a1 = a2.array
	fmt.Println(a1)
}
