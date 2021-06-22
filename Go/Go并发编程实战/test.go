package main

import (
	"fmt"
	"time"
)

func main() {
	t, _ := time.Parse("2006-01-02 15:04:05", "2020-05-06 12:32:25")
	fmt.Println(t)
}
