package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	// http로 통신하기 위한 웹서버 띄우는 부분
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

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

		user := NewUser(conn)
		addUser(user)

		go receiveHandler(user)
	}
}

func receiveHandler(user *UserConn) {
	data := make([]byte, 4096) // fixed length by packet

	for {
		readSize, err := user.conn.Read(data)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(data[:readSize]))

		//_, err = user.conn.Write(data[:n])
		//if err != nil {
		//	fmt.Println(err)
		//	return
		//}
		//sendOtherUser(user, data, readSize)
		sendAll(data, readSize)
	}
}

func sendAll(data []byte, readSize int) {
	// key: conn, value: user
	for _, value := range users {
		_, err := value.conn.Write(data[:readSize])
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func sendOtherUser(fromUser *UserConn, data []byte, readSize int) {
	// key: conn, value: user
	for key, value := range users {
		if key != fromUser.conn {
			_, err := value.conn.Write(data[:readSize])
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

//func receiveHandler(conn net.Conn) {
//	data := make([]byte, 4096) // fixed length by packet
//
//	for {
//		n, err := conn.Read(data)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		fmt.Println(string(data[:n]))
//
//		_, err = conn.Write(data[:n])
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//	}
//}
