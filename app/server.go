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

	switch os.Args[1] {
	case "PING":
		fmt.Println(respString("PONG"))
	case "ping":
		fmt.Println(respString("PONG"))
	}

	var count int
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		go handleConnection(conn, &count)
	}

	// Uncomment this block to pass the first stage
	//
	// l, err := net.Listen("tcp", "0.0.0.0:6379")
	// if err != nil {
	// 	fmt.Println("Failed to bind to port 6379")
	// 	os.Exit(1)
	// }
	// _, err = l.Accept()
	// if err != nil {
	// 	fmt.Println("Error accepting connection: ", err.Error())
	// 	os.Exit(1)
	// }
}

func handleConnection(conn net.Conn, i *int) {
	fmt.Println("this is connection", *i)
	*i += 1
	conn.Write([]byte("a"))
	conn.Close()
}

func respString(s string) string {
	return "+" + s + "\r\n"
}
