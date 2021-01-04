package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//	对go结构体的一些思考

func main() {

	s := &ListNode{}

	s1 := s
	//fmt.Printf("%p------%p\n",s,&s)
	//fmt.Println(s," ",&s," ",s.Val," ",s.Next)
	//
	//var s1 ListNode
	//
	//fmt.Printf("%v------%p\n",s1,&s1)
	//fmt.Println(s1," ",&s1," ",s1.Val," ",s1.Next)

	s.Next = &ListNode{5, nil}

	for s1 != nil {
		fmt.Println(s1.Val)
		s1 = s1.Next
	}

}
