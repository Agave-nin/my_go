//tcp服务器搭建
package main

import (
	"fmt"
	"log"
	"net"
)

func handle_conn(conn net.Conn) {
	defer conn.Close()
	fmt.Println("New connect", conn.RemoteAddr())

	buf := make([]byte, 256)
	for {
		//从网络中读
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Failed to read", err)
			break
		}
		fmt.Println(n, string(buf))
		//写回网络--回射服务器
		n, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println("Failed to write", err)
			break
		}
	}
}

func main() {
	//绑定IP端口，设置监听
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		log.Panic("Failed to listen", err)
	}
	//延迟关闭
	defer listener.Close()
	//循环等待新连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Failed to accept", err)
		}
		//与新连接通信 goroutine
		go handle_conn(conn)
	}
}
