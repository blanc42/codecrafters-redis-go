package main

import (
	"fmt"
	"net"
	"os"
	// Uncomment this block to pass the first stage
	// "net"
	// "os"
)

func main() {

	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {

	buf := make([]byte, 1028)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("no data from client", err.Error())
		os.Exit(1)
	}
	conn.Write([]byte(respString("PONG")))
	conn.Close()
}

func respString(s string) string {
	return "+" + s + "\r\n"
}
