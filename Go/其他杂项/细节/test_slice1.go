package main

import (
	"fmt"
	"unsafe"
)

//测试0
//测试在容量>长度的切片中,使用append函数添加元素
//然后查看添加前后形参的容量、长度以及每个元素的地址
//实验结果证明在test_slice.go中最后结论中的猜测是正确的
func test(s []int) {
	fmt.Println("append函数使用前：")
	fmt.Println("地址: ", unsafe.Pointer(&s), "容量：", cap(s), "长度: ", len(s))
	fmt.Println("每个元素的地址为: ")
	for i, v := range s {
		fmt.Println(v, "--", unsafe.Pointer(&s[i]))
	}
	s = append(s, 100)
	fmt.Println("append函数使用后：")
	fmt.Println("地址: ", unsafe.Pointer(&s), "容量：", cap(s), "长度: ", len(s))
	fmt.Println("每个元素的地址为: ")
	for i, v := range s {
		fmt.Println(v, "--", unsafe.Pointer(&s[i]))
	}
}

func main() {
	//s := make([]int,3,5)
	//s[0] = 1
	//s[1] = 2
	//s[2] = 3
	//fmt.Println("调用函数前,原引用的容量和长度: ",cap(s),"-",len(s))
	//test(s)
	//fmt.Println("调用函数前,原引用的容量和长度: ",cap(s),"-",len(s))

	//测试1
	//新建一个二维数组,分别使用下表赋值和append函数添加,查看每一个元素的地址
	s := make([]int, 3)
	s1 := make([][]int, 0)
	s2 := make([][]int, 1)
	fmt.Println("测试的一维切片,每个元素及其地址: ")
	for i, v := range s {
		fmt.Println(v, unsafe.Pointer(&s[i]))
	}
	s1 = append(s1, s)
	s2[0] = s
	s[0] = 199
	fmt.Println("使用append函数添加的二维数组: ")
	for i := 0; i < len(s); i++ {
		fmt.Println(s1[0][i], unsafe.Pointer(&s1[0][i]))
	}
	fmt.Println("使用下标赋值添加的二维数组: ")
	for i := 0; i < len(s); i++ {
		fmt.Println(s2[0][i], unsafe.Pointer(&s2[0][i]))
	}
	//结论
	//无论是append方式还是下标赋值方式,都是和一维数组共享内存
	//修改一维数组都会影响二维数组,但一维数组使用append函数增加元素却不会影响二维数组
	//这一点和test.go结果一致
	//所以以后使用切片时需要注意
}
