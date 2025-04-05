package main

import (
	"encoding/json"
	"fmt"
	"go-mini-chat/common/message"
	"go-mini-chat/common/proto"
	"io"
	"net"
)

func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("unmarshal login data failed, err:", err)
		return
	}

	var LoginResMes message.LoginResMes
	if loginMes.UserName == "admin" && loginMes.UserPwd == "123456" {
		// 登录成功
		LoginResMes.Code = 200
		LoginResMes.Error = ""
		fmt.Println("登录成功")
	} else {
		// 登录失败
		LoginResMes.Code = 400
		LoginResMes.Error = "用户名或密码错误"
		fmt.Println("登录失败")
	}
	data, err := json.Marshal(LoginResMes)
	if err != nil {
		fmt.Print("marshal login res data failed, err:", err)
		return
	}

	var resMes message.Message
	resMes.Type = message.LoginResMesType
	resMes.Data = string(data)
	err = proto.SendMessage(conn, &resMes)
	if err != nil {
		fmt.Println("send login res message failed, err:", err)
		return
	}
	fmt.Println(resMes)

	return nil
}

func serverProcess(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		// 处理登录消息
		err = serverProcessLogin(conn, mes)
		if err != nil {
			fmt.Println("process login message failed, err:", err)
			return
		}
		fmt.Println("收到登录消息:", mes.Data)
	case message.RegisterMesType:
		// 处理注册消息
		fmt.Println("收到注册消息:", mes.Data)
	default:
		// 处理其他消息
		fmt.Println("收到其他消息:", mes.Type, mes.Data)
	}
	return nil

}

func process(conn net.Conn) {
	defer conn.Close() // 关闭连接
	for {
		mes, err := proto.RecvMessage(conn) // 接收数据
		if err != nil {
			if err == io.EOF {
				fmt.Println("client disconnected")
			} else {
				fmt.Println("read package failed, err:", err)
			}
			return
		}
		err = serverProcess(conn, mes) // 处理数据
		if err != nil {
			fmt.Println("process message failed, err:", err)
			return
		}
	}
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
		go process(conn) // 启动一个goroutine处理连接
	}
}
