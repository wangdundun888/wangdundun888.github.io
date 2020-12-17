package main

import (
	"fmt"
	"net"
)

func main() {
	var ip net.IP
	ip = net.ParseIP("12.1222.23.43")
	fmt.Println(ip)
}
