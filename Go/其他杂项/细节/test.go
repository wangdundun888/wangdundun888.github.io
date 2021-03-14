package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//1.使用append函数会复制一个切片
	//修改被传入的切片并不会影响二维切片
	//ss := make([][]int,0)
	//s := make([]int,1)
	//ss = append(ss,s)
	//fmt.Println(ss)
	//s = append(s,1)
	//fmt.Println(ss)

	//2.直接使用下标赋值
	//和append函数的结果一样
	ss := make([][]int, 0)
	s := make([]int, 0)
	fmt.Println(ss, "--", s)
	s = append(s, 1)
	fmt.Println(ss, "--", s)

	//p := new([]int)
	// 思考 *[]int 与[]*int 的区别
	// make 与 new 的区别

	//3.append中使用指针
	a := 3
	var p *int = &a
	fmt.Println(a, "--", *p)
	a = 4
	fmt.Println(a, "--", *p)
	*p = 5
	fmt.Println(a, "--", *p)

	fmt.Println(unsafe.Pointer(&s))
	s = append(s, a)
	fmt.Println(unsafe.Pointer(&s))
	s = append(s, *p)
	fmt.Println(unsafe.Pointer(&s))
	fmt.Println(s)
	a = 4
	fmt.Println(s)
	*p = 6
	fmt.Println(s)

	fmt.Println(ss)
	fmt.Println(unsafe.Pointer(&ss), "--", unsafe.Pointer(&s))
	Test(&ss, &s)
	fmt.Println(unsafe.Pointer(&ss), "--", unsafe.Pointer(&s))
	fmt.Println(ss)

	fmt.Println(s)
	Test1(&s)
	fmt.Println(s)

}

func Test(ss *[][]int, s *[]int) {
	fmt.Println(unsafe.Pointer(ss), "--", unsafe.Pointer(s))
	(*ss) = append(*ss, *s)
	fmt.Println(unsafe.Pointer(ss), "--", unsafe.Pointer(s))
}

func Test1(s *[]int) {
	fmt.Println(*s)
	(*s) = append(*s, 100)
	fmt.Println(*s, unsafe.Pointer(s))
	(*s) = (*s)[:len(*s)-1]
	fmt.Println(*s, unsafe.Pointer(s))
	fmt.Println(*s)
}
