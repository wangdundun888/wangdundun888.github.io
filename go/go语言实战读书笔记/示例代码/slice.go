package main

import "fmt"

func main() {
	//创建切片的方式
	//方式1:make
	s := make([]int, 3) // s := make([]type,len,cap) cap若不指明则cap=len 如指明须大于len,否则会发生编译错误
	fmt.Println(s)
	//方式2:字面量
	s1 := []int{1, 2, 3}
	fmt.Println(s1)
	s1 = []int{2: 3, 1: 2, 3: 0} //通过字面量索引方式声明, 可乱序,可省略前面的索引,len = 最大索引+1
	fmt.Println(s1)
	/*
		nil切片和空切片
		var s []int  //nil切片
		//空切片
		s := make([]int,0)
		s := []int{}
	*/
	//使用切片与赋值
	_ = s[0]
	s[1] = 25
	fmt.Println(s[1:3], s[:3], s[0:]) // 赋值与访问方式与数组无异, 注意 不能进行批量赋值  s[:] = 0 是错的
	//迭代切片
	//方式1:传统索引下标法
	for i := 0; i < len(s); i++ {
		fmt.Print(s[i], " ")
	}
	//方式2:range
	for index, value := range s {
		fmt.Print(index, " ", value, " ")
		//注意,使用range方式迭代时,index和value变量都会有一个自己的地址,且迭代过程中不会改变
	}
	//向切片增加元素
	fmt.Println("\ns.cap: ", cap(s), "s.len: ", len(s))
	s = append(s, 50)
	fmt.Println(s)
	fmt.Println("s.cap: ", cap(s), "s.len: ", len(s))
	/*
		使用append增加切片元素时,应该要注意的事:
			如果cap还有未用的容量,则直接使用,如果已经没有容量,则会重新申请一块地址,把就切片复制过去,保证地址的连续性
			新的cap是原来的两倍
	*/
	//使用切片创建切片
	newSlice := s[1:3]
	fmt.Println(newSlice)
	//使用newSlice := s[i:j],假如s的cap是k,那么newSlice的长度是j-i,容量是k-i
	//newSlice := s[i:j:k], 长度是j-i,容量是k-i,注意不要超过界限,且k-i不能大于cap(s)-i
	//使用切片创建切片,两个切片会共享内存,修改切片,都会感知到,可以设置新切片的len=cap,利用append函数的机制,让新切片与原切片分离

	//func(s []int) 函数中使用切片,也是引用传递
}
