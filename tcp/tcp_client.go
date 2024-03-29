package tcp

import (
	"fmt"
	"os"
	"flag"
	"io"
	"net"
)

var (
	Host, Port string
)

func handleConnectionNoLog(conn net.Conn) {

    rAddr := conn.RemoteAddr().String()
    defer fmt.Printf("Closed connection from %v\n", rAddr)

    // This will accomplish the echo if we do not want to log
    io.Copy(conn, conn)
}

func tcpClient() {

	if len(os.Args) != 3 {
		fmt.Println("Usage:", os.Args[0], "<host>", "<port>")
		return
	}

	Host := os.Args[1]
	Port := os.Args[2]

	flag.Parse()

	// Convert host and port to host:port
	t := net.JoinHostPort(Host, Port)

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
