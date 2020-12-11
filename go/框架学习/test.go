package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int64 = 65535111112223123
	str := strconv.FormatInt(i, 16)
	fmt.Println(str)
}
