package main

import "fmt"

//在GO中,可以给任意类型(包括内置类型,但不包括指针类型)添加相应的方法
//例子如下:
type String string

//此方式与上一种有何不同?
//上一种属于类型重新命名
//此方式属于结构体声明
type String1 struct {
	string
}

type test interface {
}

func (s *String) length() int {
	return len(*s)
}
func (s *String1) length() int {
	return len(s.string)
}
func main() {
	s := String("hello world")
	//s1 := String1{"helloworld"}
	/*
			初始化结构体的四种方法:
				s1 := new(String1) //返回一个String1指针
				s1 := &String1{}
				s1 := &String1{"helloworld"}
				s1 := &String1{string:"helloworld"}
		值得注意的是:
			var s1 String1 和 var s1 *String1
			var s1 String1 是值类型声明,会初始化为对应类型的空值,之后直接调用s1.length()是没有问题的
			但var s1 *String 是指针类型声明,此时s1是一个空指针,直接调用会出引发异常
			同时,如果结构体内有指针类型,也要注意值的初始化,避免空指针引用

	*/
	var s1 String1
	fmt.Println(s.length())
	fmt.Println(s1.length())

	/*
			var t1 test
			接口值是一个两个字长度的数据结构
			第一个字包含一个指向内部表的指针
			第二个字是一个指向所存储值的指针

			另外在接口实现中,如果方法接收者是指针,那么只有指向这个类型的指针才能实现对应的接口
		   				  如果方法接收者是值,那么这个类型的值或者类型都能实现对应的接口
			这是因为,编译器并不是总能自动获取一个值的地址,因为有些值会被优化为常量,而常量是不能取地址的
			这里包含了一个方法集的概念,等之后看了内部内存分布,可能会更加明朗和清晰
	*/

}
