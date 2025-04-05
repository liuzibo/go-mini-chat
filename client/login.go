package main

import (
	"encoding/json"
	"fmt"
	"go-mini-chat/common/message"
	"go-mini-chat/common/proto"
	"net"
)

func login(userName string, userPwd string) (err error) {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("Failed to connect to server")
		return err
	}
	defer conn.Close()

	var loginMes message.LoginMes
	loginMes.UserName = userName
	loginMes.UserPwd = userPwd
	mesData, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("Failed to marshal login message")
		return err
	}
	var mes message.Message
	mes.Type = message.LoginMesType
	mes.Data = string(mesData)
	// Serialize the message
	err = proto.SendMessage(conn, &mes)
	if err != nil {
		return err
	}
	// Receive the response
	var resMes *message.Message
	fmt.Println("Waiting for login response...")
	resMes, err = proto.RecvMessage(conn)
	if err != nil {
		return err
	}
	fmt.Println("Login response:", resMes.Data)

	return nil
}
