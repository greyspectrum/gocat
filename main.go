package main

import (
	"fmt"
	"os"
)

var (
	host, port string
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Usage:", os.Args[0], "host", "port")
		return
	}

	host := os.Args[1]
	port := os.Args[2]

	fmt.Println(host)
	fmt.Println(port)
}
