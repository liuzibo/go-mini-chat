package serverprocess

import (
	"encoding/json"
	"fmt"
	"go-mini-chat/internal/common/message"
	"go-mini-chat/internal/common/utils"
	"go-mini-chat/internal/server/model"
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
	user, err := model.MyUserDao.Login(loginMes.UserName, loginMes.UserPwd)
	if err != nil {
		fmt.Println("login failed, err:", err)
		return
	}
	if loginMes.UserName == user.Username && loginMes.UserPwd == user.Password {
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

func (userProcess *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	// 读取注册消息
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("unmarshal register data failed, err:", err)
		return
	}

	// 注册用户 封装并发送注册结果
	var registerResMes message.RegisterResMes
	err = model.MyUserDao.Register(registerMes.UserName, registerMes.UserPwd)
	if err != nil {
		if err == model.ErrUserExists {
			registerResMes.Code = 505
			registerResMes.Error = "用户名已存在"
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册失败, 未知错误"
		}
		fmt.Println("register failed, err:", err)
	} else {
		registerResMes.Code = 200
		registerResMes.Error = ""
	}
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("marshal register res data failed, err:", err)
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
		fmt.Println("send register res message failed, err:", err)
		return
	}
	// fmt.Println(resMes)
	return nil
}
