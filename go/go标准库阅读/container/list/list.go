/*
	官方list.go源码:https://github.com/golang/go/blob/master/src/container/list/list.go
	学习list.go是因为在学习的过程中发现没有直接的stack可用,然后发现list可以临时充当一个stack来用,也
	可以封装为一个stack来用.
	本次主要学习Element和List两个数据结构,以及它们的方法,包括可导出以及不可导出方法
*/

package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	/*
		list包有一个公开函数New()和两个公开属性Element和List
		New():返回一个初始化的list,返回一个List指针
		Element:
				type Element struct {
					//Element相当于以前学习链表时的Node
					next,prev  *Element
					//双向链表,next域和prev域,不可导出,所以在包外不能通过Element.next方式操作
					Value interface{}
					//可导出值域,任何实现了空接口的类型都可以存储
					l *List
					//不可导出的list域,指向该Element所属list,目前可看的作用是在O(1)时间内移除链表中的任意一个Element
			}
		List:
				type List struct {
					root Element //root.next指向list第一个元素,root.prev指向list最后一个元素
					len  int     //链表长度
				}
	*/

	l.Init()
	/*
		New创建list时,会调用Init方法初始化,该方法会让list.root的next和prev指向root,同时令len=0
		list上已经有元素的前提下调用Init会导致清空list,即需要清空list时可调用Init方法
	*/

	l.PushFront("hello world")
	l.PushFront([]int{1, 2, 3}) //插入一个切片
	/*
		func (l *List) PushFront(v interface{}) *Element
		插入一个值v,v将会封装在一个Element里面,即使这个v本身是一个Element,也会进行封装放到一个新的Element的Value里,并返回这个Element
		PushFront会调用lazyInit方法,该方法先判断list.root.next是否为nil,如是nil调用Init方法(ps:我暂时不知道这有啥用~)
		然后调用insertValue方法将值插入list
		insertValue:
			func (l *List) insertValue(v interface{}, at *Element) *Element
			该方法会调用insert方法,意指在at后面插入Value为v的Element
		insert:
			func (l *List) insert(e, at *Element) *Element
			在at后插入e,插入后将len++,过程略
		思考:为什么要将insert过程从insertValue中拆分?
			insertValue是插入值,insert是插入一个Element,行为不同
			insertValue需要将Value封装为一个Element再调用insert
	*/

	l.PushBack(1123) //在list尾插入元素v,和PushFront异曲同工

	l1 := list.New()
	l1.PushFrontList(l)
	/*
		func (l *List) PushFrontList(other *List)
		在list头插入一个otherList,该方法先调用lazyInit方法,然后从otherList尾开始遍历,
		调用insertValue,参数是otherList上的值,相当于复制了一份,与被插入链表无关,从尾开始
		遍历是为了保证插入顺序与原来相同
	*/
	l1.PushBackList(l) //在list尾插入另一个list,与PushFrontList异曲同工

	fmt.Println("l的长度为: ", l.Len()) //返回list.len,即该list的长度

	e := l.Front()
	/*
		func (l *List) Front() *Element
		返回链表的第一个元素,如果len=0,则返回nil,所有使用前,应该进行判空
	*/
	if e != nil {
		fmt.Println("l不为空,第一个元素值为:", e.Value)
	}

	e = l.Back() //获取尾Element,内部结构与Front基本相同
	if e != nil {
		fmt.Println("l不为空,最后一个元素值为:", e.Value)
	}

	e = l.InsertAfter("last one element", e)
	/*
		func (l *List) InsertAfter(v interface{}, mark *Element) *Element
		在标记点mark后插入一个Value为v的Element,并返回这个Element
		先判断mask是否在list上,即mark.list == l,不是则返回nil
		然后调用insertValue插入
	*/
	e1 := l.Back()
	if e != nil && e1 != nil {
		fmt.Println("测试InsertAfter方法:", e.Value, "-->", e1.Value)
	}

	e = l.InsertBefore(4.5, e) //在mark元素前插入另一个元素,内部实现参考InsertAfter

	v := l.Remove(e)
	/*
		func (l *List) Remove(e *Element) interface{}
		移除list上的e, 并返回e.Value
		先判断参数e是否在list上,即e.list == l,如在,则调用remove方法
		注意,不管在不在list上,最后都是返回e.Value
		remove:
			func (l *List) remove(e *Element) *Element
			调整前后指针指向,长度减1,然后返回e
			主要要使e的next prev list都为nil.
			个人理解:不然没有指针指向e,但又有e又指向其它,倒是无法自动回收e,导致内存泄漏(等看go的gc机制才能彻底了解)
	*/
	fmt.Println(v)
	/*
		剩下四个方法:
			MoveToFront(e *Element)
				移动e到list头,先判断e是否在list上以及e是否本来就在list头
				如果都不,调用move方法
			MoveToBack(e *Element)
				参考MoveToFront
			MoveBefore(e, mark *Element)
				移动e到mark前,判断e和mark是否为同一个Element,以及e和mark是否都在list上
				如果都不,调用move方法
			MoveAfter(e, mark *Element)
				参考MoveBefore
		move:
			move(e, at *Element) *Element
				移动e到at后,先判断e==at,如果不,开始移动,然后返回e
	*/
	fmt.Println("list模仿stack开始...")
	l.Init()
	/*
		分析:一个简单的stack有isEmpty,Pop,Push,Top四个操作
		isEmpty对应list的len长度是否为0
		假定就在list存储元素
		Top对应Front
		Push对应PushFront
		Pop对应先Top一个Element,然后Romove这个Element
	*/
	fmt.Print("入栈顺序: ")
	for i := 1; i < 10; i++ {
		l.PushFront(i)
		fmt.Print(" ", i)
	}
	fmt.Println("\n栈的大小为: ", l.Len())
	e = l.Front()
	if e != nil {
		fmt.Println("栈顶元素为: ", e.Value)
	}
	fmt.Print("出栈顺序: ")
	for i := l.Len(); i > 0; i-- {
		e = l.Front()
		if e != nil {
			fmt.Print(" ", e.Value)
			l.Remove(e)
		}
	}
	fmt.Println("\n栈的大小为: ", l.Len())
	fmt.Println("list模仿stack结束...")
	/*
		终于阅读完了list.go源码,代码不多,却拖了挺久
		list能模仿stack,也可以模仿queue,在back入队,在front出队即可
		最后讲一个"气球-绳子":
			假定一个气球可以被绳子绑或者不绑,一根绳子只能绑一个气球,气球若是没有则会飞向高空,再也找不到
			现在有一个气球被一根绳子绑着,需要更换另外一根绳子,步骤是先绑上另外一根绳子,再解开原来的绳子.
			气球对应内存,绳子对应指针.
	*/
}
