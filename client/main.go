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
	//go func(c net.Conn) {
	//	data := make([]byte, 4096)
	//
	//	for {
	//		n, err := c.Read(data)
	//		if err != nil {
	//			fmt.Println(err)
	//			return
	//		}
	//
	//		fmt.Println(string(data[:n]))
	//		time.Sleep(1 * time.Second)
	//	}
	//}(client)

	//go func(c net.Conn) {
	//	i := 0
	//	for {
	//		s := "Hello " + strconv.Itoa(i)
	//
	//		_, err := c.Write([]byte(s)) // 서버로 부터 데이터 전송
	//		if err != nil {
	//			fmt.Println(err)
	//			return
	//		}
	//		i++
	//		time.Sleep(1 * time.Second)
	//	}
	//}(client)
	//var input string
	//fmt.Scanln(&input)
}
