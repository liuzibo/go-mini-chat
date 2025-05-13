package main

import (
	"fmt"
	"go-mini-chat/internal/client/clientprocess"
)

func main() {
	// 接收用户输入
	var key int

	var userName string
	var userPwd string

	for {
		fmt.Println("---------------欢迎登录多人聊天系统---------------")
		fmt.Println("\t\t1. 登录")
		fmt.Println("\t\t2. 注册")
		fmt.Println("\t\t3. 退出")
		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("请输入账号：")
			fmt.Scanln(&userName)
			fmt.Println("请输入密码：")
			fmt.Scanln(&userPwd)
			userProcess := &clientprocess.UserProcess{}
			userProcess.Login(userName, userPwd)
		case 2:
			fmt.Println("请输入账号：")
			fmt.Scanln(&userName)
			fmt.Println("请输入密码：")
			fmt.Scanln(&userPwd)
			userProcess := &clientprocess.UserProcess{}
			userProcess.Register(userName, userPwd)
		case 3:
			return
		default:
			fmt.Println("输入有误，请重新输入")
		}
	}
}
