package main

import (
	"fmt"
)

func main() {
	// 接收用户输入
	var key int
	// 循环控制
	var loop = true
	var userName string
	var userPwd string

	for loop {
		fmt.Println("---------------欢迎登录多人聊天系统---------------")
		fmt.Println("\t\t1. 登录")
		fmt.Println("\t\t2. 注册")
		fmt.Println("\t\t3. 退出")
		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("请输入用户名：")
			fmt.Scanln(&userName)
			fmt.Println("请输入密码：")
			fmt.Scanln(&userPwd)
			login(userName, userPwd)
			loop = false
		case 2:
			fmt.Println("注册功能暂未开放")
			loop = false
		case 3:
			loop = false
		default:
			fmt.Println("输入有误，请重新输入")
		}
	}
}
