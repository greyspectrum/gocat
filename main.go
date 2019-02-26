package main

import (
	"fmt"
	"os"
	"flag"
	"io"
	"net"
)

var (
	host, port string
)

func handleConnectionNoLog(conn net.Conn) {

    rAddr := conn.RemoteAddr().String()
    defer fmt.Printf("Closed connection from %v\n", rAddr)

    // This will accomplish the echo if we do not want to log
    io.Copy(conn, conn)
}

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Usage:", os.Args[0], "<host>", "<port>")
		return
	}

	host := os.Args[1]
	port := os.Args[2]

	fmt.Println(host)
	fmt.Println(port)

	flag.Parse()

	// Convert host and port to host:port
	t := net.JoinHostPort(host, port)

	// Listen for connections on BindIP:BindPort
	ln, err := net.Listen("tcp", t)
	if err != nil {

	// If we cannot bind, print the error and quit
        panic(err)
    }

	// Wait for connections
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
	}

	fmt.Printf("Received connection from %v\n", conn.RemoteAddr().String())
	go handleConnectionNoLog(conn)
	}
}
