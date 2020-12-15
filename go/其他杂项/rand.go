package main

import (
	"fmt"
	"math/rand"
	"time"
)

const RANGE = 100

//生成随机数例子
func main() {
	//将时间作为种子数
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		//生成一个RANGE以内的整数
		nums := rand.Intn(RANGE)
		fmt.Print(nums, " ")
	}
}
