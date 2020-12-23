package main

import (
	"fmt"
	"net"
	"time"
)

func receiveHandler(client net.Conn) {
	fmt.Println("start receiver handler")
	go func(c net.Conn) {
		data := make([]byte, 4096)

		for {
			n, err := c.Read(data)
			fmt.Println("read size: ", n)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(string(data[:n]))
			time.Sleep(1 * time.Second)
		}
	}(client)
}
