package main

import (
	"bufio"
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

	scanner := bufio.NewScanner(conn)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		conn.Write([]byte(respString("PONG")))
	}
	conn.Close()
}

func respString(s string) string {
	return "+" + s + "\r\n"
}
