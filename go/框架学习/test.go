package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	flag.Parse()
	args := flag.Args()
	url := args[0]
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	fmt.Println(resp.Body)
}
