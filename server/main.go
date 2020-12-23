package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("server running 8888 port")
	// create port by 8888
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	for {
		// wait, when incoming user
		conn, err := listener.Accept()
		if err != nil {
			fmt.Print(err)
			continue
		}
		defer conn.Close()

		go receiveHandler(conn)
	}
}

func receiveHandler(conn net.Conn) {
	data := make([]byte, 4096) // fixed length by packet

	for {
		n, err := conn.Read(data)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(data[:n]))

		_, err = conn.Write(data[:n])
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
