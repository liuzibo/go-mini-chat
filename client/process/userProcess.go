package process

import (
	"encoding/json"
	"fmt"
	"go-mini-chat/client/message"
	"go-mini-chat/client/utils"
	"net"
)

type UserProcess struct {
}

func (userProcess *UserProcess) Login(userName string, userPwd string) (err error) {
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
	transfer := utils.Transfer{Conn: conn}
	err = transfer.SendMessage(&mes)
	if err != nil {
		return err
	}
	// Receive the response
	var resMes *message.Message
	fmt.Println("Waiting for login response...")
	resMes, err = transfer.RecvMessage()
	if err != nil {
		return err
	}
	// Deserialize the response
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(resMes.Data), &loginResMes)
	if err != nil {
		fmt.Println("Failed to unmarshal login response")
		return err
	}
	if loginResMes.Code == 200 {
		go serverProcessMes(conn)
		fmt.Println("Login success")
		for {
			ShowMenu()
		}
	} else if loginResMes.Code == 400 {
		fmt.Println("Login failed:", loginResMes.Error)
	} else {
		fmt.Println("Unknown error:", loginResMes.Error)
	}
	// fmt.Println("Login response:", resMes.Data)

	return nil
}
