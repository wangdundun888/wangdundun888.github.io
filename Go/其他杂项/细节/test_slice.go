package main

import (
	"fmt"
	"unsafe"
)

//主要测试切片传参到函数之后的修改问题

//测试0
//测试传参修改已有元素是否会影响原切片
//测试结果:会影响
func test(s []int) {
	fmt.Println("在test中s的地址: ", unsafe.Pointer(&s))
	s[2] = 4
}

//测试1
//测试使用append函数添加元素是否会影响原切片
//测试结果,不会影响
func test1(s []int) {
	fmt.Println("在test1中s的地址: ", unsafe.Pointer(&s))
	s = append(s, 5)
}

//测试2
//在测试1的基础上,把参数的形式改变
func test2(s *[]int) {
	fmt.Println("在test2中s的地址: ", unsafe.Pointer(s))
	*s = append(*s, 5)
	(*s) = (*s)[:len(*s)-2]
	fmt.Println("在test2中s的地址: ", unsafe.Pointer(s))
}

func main() {

	s := make([]int, 4, 5)
	fmt.Println(s)
	test(s)
	fmt.Println(s)

	test1(s)
	fmt.Println(s, "地址:", unsafe.Pointer(&s))

	test2(&s)
	fmt.Println(s)
}

/*
	结论:
		通过测试1可以看到,当参数进入函数时,引用参数和原来的引用不是同一个引用,
		所以使用append函数不会影响到原来的引用,
		但是两个引用却是指向同一块内存,修改内存会影响到原来的内存.
		那么现在问题来了,若原切片的容量比长度长,那么在test函数中使用append函数会影响到原引用吗?
			猜测,每个引用都有自己的长度和容量,test函数中使用append函数,会占用剩余容量,但对原引用不可见.
			在同目录下的test_slice1.go进行最后的两个测试.
*/
