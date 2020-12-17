package main

import "math"

//最大优先队列demo
//假设有集合S

//返回S集合中最大关键字的元素,时间O(1)
func maximum(A []int) int {
	return A[1]
}

//去掉并返回S集合中最大关键字的元素,时间为O(lgn)
func extractMax(A []int) int {
	if heapsize < 1 {
		panic("heap underflow")
	}
	max := A[1]
	A[1] = A[heapsize]
	heapsize -= 1
	//维护最大堆性质
	maxHeapify(A, 1)
	return max
}

//将堆中i的关键字增加到key,假定key大于i
func increaseKey(A []int, i, key int) {
	if A[i] < key {
		panic("new key is smaller than current key")
	}
	A[i] = key
	//提升了关键字key后,不断向上浮,找到自己合适的位置,时间为O(lgn)
	//上浮的过程中可证并不会破坏最大堆的性质,所以不需要维护最大堆
	for i > 1 && A[i] > A[parent(i)] {
		A[i], A[parent(i)] = A[parent(i)], A[i]
		i = parent(i)
	}
}

//向堆中插入一个元素x,时间为O(lgn)
func insert(A []int, x int) {
	heapsize += 1
	//插入一个最小值,使之依旧为最大堆,然后复用increaseKey
	A[heapsize] = math.MinInt32
	increaseKey(A, heapsize, x)

	//如果不直接复用increaseKey也可以,如下
	//heapsize += 1
	//A[heapsize] = x
	//i := heapsize
	//for i > 1 && A[i] > A[parent(i)] {
	//	A[i], A[parent(i)] = A[parent(i)], A[i]
	//	i = parent(i)
	//}
	//同样的道理
}
