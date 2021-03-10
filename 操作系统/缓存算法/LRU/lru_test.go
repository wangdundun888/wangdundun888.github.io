package LRU

import (
	"fmt"
	"strconv"
	"testing"
)

var mlc *myLruCache = NewMLC(3)

func TestMyLruCache_Set(t *testing.T) {
	for i := 1; i < 5; i++ {
		mlc.Set(strconv.Itoa(i), strconv.Itoa(i))
		fmt.Print("now lruCache : ")
		printList()
	}
}

func TestMyLruCache_Get(t *testing.T) {
	for i := 1; i < 6; i++ {
		v := mlc.Get(strconv.Itoa(i))
		fmt.Println("get lruCache,k-v", i, "-", v)
		fmt.Print("now lruCache : ")
		printList()
	}
}

func printList() {
	for e := mlc.List.Front(); e != nil; e = e.Next() {
		for k, e1 := range mlc.Map {
			if e1 == e {
				fmt.Print("(", k, "-", e.Value.(string), ")")
			}
		}
	}
	fmt.Println()
}
