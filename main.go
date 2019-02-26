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

}
