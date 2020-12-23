package main

import (
	"fmt"
	"net"
	"time"
)

func sendHandler(client net.Conn) {
	fmt.Println("start send handler")
	go func(c net.Conn) {
		//i := 0
		for {
			var input string
			_, err := fmt.Scanln(&input)
			fmt.Println("input: ", input)
			if err != nil {
				fmt.Println(err)
				return
			}
			//s := "Hello " + strconv.Itoa(i)

			n, writeErr := c.Write([]byte(input)) // 서버로 부터 데이터 전송
			fmt.Println("send size: ", n)
			if writeErr != nil {
				fmt.Println(writeErr)
				return
			}
			//i++
			time.Sleep(1 * time.Second)
		}
	}(client)
}
