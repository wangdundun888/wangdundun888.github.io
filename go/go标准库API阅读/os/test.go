package main

import (
	"fmt"
	"os"
	user "os/user"
	"strings"
)

func main() {
	hostName, err := os.Hostname()
	if err == nil {
		fmt.Println(hostName)
	}
	env := os.Environ()
	for _, v := range env {
		if strings.HasPrefix(v, "GOPATH") {
			fmt.Println(v)
			fmt.Println(os.Getenv("GOPATH"))
		}
	}
	rootPath, _ := os.Getwd()
	fmt.Println(rootPath)

	user1, _ := user.Current()
	fmt.Println(user1.HomeDir)
}
