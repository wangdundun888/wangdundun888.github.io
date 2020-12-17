package main

import (
	"fmt"
	"math"
	"time"
)

func getTime() {

	time.Sleep(time.Duration(100000))

}

func main() {
	fmt.Println(time.Now())
	getTime()
	for i := 1; i < math.MaxInt32; i++ {

	}
	fmt.Println(time.Now())
}
