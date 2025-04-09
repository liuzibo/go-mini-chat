package main

import (
	"fmt"
	"net"
)

func mainprocess(conn net.Conn) {
	defer conn.Close() // 关闭连接
	processor := Processor{
		Conn: conn,
	}
	processor.connprocess()
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go mainprocess(conn) // 启动一个goroutine处理连接
	}
}
