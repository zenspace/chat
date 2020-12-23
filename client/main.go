package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("start client")

	// 해당 서버로 접속 시도
	client, err := net.Dial("tcp", "127.0.0.1:8888")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()

	receiveHandler(client)
	sendHandler(client)
	//var input string
	//fmt.Scanln(&input)
}
