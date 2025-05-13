package serverprocess

import (
	"fmt"
	"go-mini-chat/internal/common/message"
	"go-mini-chat/internal/common/utils"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (processor *Processor) serverProcess(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		// 处理登录消息
		userProcessor := UserProcess{Conn: processor.Conn}
		err = userProcessor.ServerProcessLogin(mes)
		if err != nil {
			fmt.Println("serverprocess login message failed, err:", err)
			return
		}
		fmt.Println("收到登录消息:", mes.Data)
	case message.RegisterMesType:
		// 处理注册消息
		userProcessor := UserProcess{Conn: processor.Conn}
		err = userProcessor.ServerProcessRegister(mes)
		if err != nil {
			fmt.Println("serverprocess register message failed, err:", err)
			return
		}
		fmt.Println("收到注册消息:", mes.Data)
	default:
		// 处理其他消息
		fmt.Println("收到其他消息:", mes.Type, mes.Data)
	}
	return nil

}

func (processor *Processor) Connprocess() {
	defer processor.Conn.Close() // 关闭连接
	for {
		transfer := utils.Transfer{Conn: processor.Conn}
		mes, err := transfer.RecvMessage() // 接收数据
		if err != nil {
			if err == io.EOF {
				fmt.Println("client disconnected")
			} else {
				fmt.Println("read package failed, err:", err)
			}
			return
		}
		err = processor.serverProcess(mes) // 处理数据
		if err != nil {
			fmt.Println("serverprocess message failed, err:", err)
			return
		}
	}
}
