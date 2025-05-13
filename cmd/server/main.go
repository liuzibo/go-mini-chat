package main

import (
	"fmt"
	"go-mini-chat/internal/server/model"
	"go-mini-chat/internal/server/serverprocess"
	"net"
)

func mainprocess(conn net.Conn) {
	defer conn.Close() // 关闭连接
	processor := serverprocess.Processor{
		Conn: conn,
	}
	processor.Connprocess()
}

func main() {
	rdb := model.InitRedisPool("localhost:6379", "", 0, 10, 5)
	model.MyUserDao = model.NewUserDao(rdb) // 初始化用户Dao
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
