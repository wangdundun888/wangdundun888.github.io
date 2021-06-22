package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func spy(c1, c2 <-chan int) {
	for {
		time.Sleep(time.Second * 2)
		select {
		case <-c1:
			fmt.Println("通道1受到信息,值为: ", <-c1)
		case <-c2:
			fmt.Println("通道2受到信息,值为: ", <-c2)
		default:
			fmt.Println("timeout")
		}
	}
}

func sub(c chan<- int) {
	for {
		time.Sleep(time.Second * 2)
		rand.Seed(time.Now().UnixNano())
		nums := rand.Intn(100)
		c <- nums
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	c1, c2 := make(chan int), make(chan int)
	go spy(c1, c2)
	go sub(c1)
	go sub(c2)
	wg.Wait()
}
