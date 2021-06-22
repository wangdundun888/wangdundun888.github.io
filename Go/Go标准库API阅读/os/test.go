package main

import (
	"fmt"
	"os"
	"os/user"
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

	//获得进程id,获得进程的父进程id
	pid := os.Getpid()
	ppid := os.Getppid()
	fmt.Println(pid, "  ", ppid)
}
