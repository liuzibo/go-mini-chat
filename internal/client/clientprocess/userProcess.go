package clientprocess

import (
	"encoding/json"
	"fmt"
	"go-mini-chat/internal/common/message"
	"go-mini-chat/internal/common/utils"
	"net"
)

type UserProcess struct {
}

func (userProcess *UserProcess) Login(UserName string, userPwd string) (err error) {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("Failed to connect to server")
		return err
	}
	defer conn.Close()

	// 登录消息
	var loginMes message.LoginMes
	loginMes.UserName = UserName
	loginMes.UserPwd = userPwd
	mesData, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("Failed to marshal login message")
		return err
	}
	// 封装消息
	var mes message.Message
	mes.Type = message.LoginMesType
	mes.Data = string(mesData)
	// 发送消息
	transfer := utils.Transfer{Conn: conn}
	err = transfer.SendMessage(&mes)
	if err != nil {
		return err
	}
	// 接收返回
	var resMes *message.Message
	fmt.Println("Waiting for login response...")
	resMes, err = transfer.RecvMessage()
	if err != nil {
		return err
	}
	// 反序列化消息
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(resMes.Data), &loginResMes)
	if err != nil {
		fmt.Println("Failed to unmarshal login response")
		return err
	}
	if loginResMes.Code == 200 {
		// go serverProcessMes(conn)
		fmt.Println("Login success")
		ShowMenu()
	} else if loginResMes.Code == 400 {
		fmt.Println("Login failed:", loginResMes.Error)
	} else {
		fmt.Println("Unknown error:", loginResMes.Error)
	}

	return nil
}

func (userProcess *UserProcess) Register(UserName string, userPwd string) (err error) {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("Failed to connect to server")
		return err
	}
	defer conn.Close()
	// 注册消息
	var registerMes message.RegisterMes
	registerMes.UserName = UserName
	registerMes.UserPwd = userPwd
	mesData, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("Failed to marshal login message")
		return err
	}
	var mes message.Message
	mes.Type = message.RegisterMesType
	mes.Data = string(mesData)
	// 发送消息
	transfer := utils.Transfer{Conn: conn}
	err = transfer.SendMessage(&mes)
	if err != nil {
		return err
	}
	// 接收返回信息
	var resMes *message.Message
	fmt.Println("Waiting for register response...")
	resMes, err = transfer.RecvMessage()
	if err != nil {
		return err
	}
	// 反序列化返回信息
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(resMes.Data), &registerResMes)
	if err != nil {
		fmt.Println("Failed to unmarshal register response")
		return err
	}
	if registerResMes.Code == 200 {
		fmt.Println("Register success")
	} else if registerResMes.Code == 505 {
		fmt.Println("Register failed:", registerResMes.Error)
	} else {
		fmt.Println("Unknown error:", registerResMes.Error)
	}
	return nil
}
