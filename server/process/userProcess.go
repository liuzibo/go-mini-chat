package process

import (
	"encoding/json"
	"fmt"
	"go-mini-chat/server/message"
	"go-mini-chat/server/utils"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}

func (userProcess *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
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
	transfer := &utils.Transfer{
		Conn: userProcess.Conn,
	}
	err = transfer.SendMessage(&resMes)
	if err != nil {
		fmt.Println("send login res message failed, err:", err)
		return
	}
	// fmt.Println(resMes)

	return nil
}
