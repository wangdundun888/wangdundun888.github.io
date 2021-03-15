package main

import (
	"fmt"
	"unsafe"
)

func dfs(s []int, t int) {
	if t == 5 {
		return
	} else {
		fmt.Println("append之前")
		fmt.Println(s, unsafe.Pointer(&s), "容量及长度:", cap(s), " ", len(s))
		for i, v := range s {
			fmt.Println(v, unsafe.Pointer(&s[i]))
		}
		s = append(s, t)
		fmt.Println("append之后")
		fmt.Println(s, unsafe.Pointer(&s), "容量及长度:", cap(s), " ", len(s))
		for i, v := range s {
			fmt.Println(v, unsafe.Pointer(&s[i]))
		}
		dfs(s, t+1)
	}
}

//结论
//与slice的扩容机制有关
//当slice容量与长度一致时使用append函数,则会申请心得内存存放元素,容量扩增
//若容量>长度时使用append则会使用原来的地址,修改原来长度的元素会影响原引用

func main() {
	s := []int{99, 98, 97}
	fmt.Println("dfs之前:", s, unsafe.Pointer(&s))
	for i, v := range s {
		fmt.Println(v, unsafe.Pointer(&s[i]))
	}
	dfs(s, 0)
	fmt.Println("dfs之后:", s, unsafe.Pointer(&s))
	for i, v := range s {
		fmt.Println(v, unsafe.Pointer(&s[i]))
	}
}
