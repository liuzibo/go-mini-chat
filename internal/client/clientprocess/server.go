package clientprocess

import (
	"fmt"
)

func ShowMenu() {
	fmt.Println("Welcome xxx to go-mini-chat!")
	fmt.Println("1. 显示在线用户")
	fmt.Println("2. 发送消息")
	fmt.Println("3. 信息列表")
	fmt.Println("4. 退出")
	for {
		var key int
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("Show online users")
			// ShowOnlineUsers()
		case 2:
			fmt.Println("Send message")
			// SendMessage()
		case 3:
			fmt.Println("Show message list")
			// ShowMessageList()
		case 4:
			fmt.Println("Bye!")
			return
		default:
			fmt.Println("Invalid input, please try again.")
			// ShowMenu()
		}
	}

}

// func serverProcessMes(Conn net.Conn) {
// 	defer Conn.Close()
// 	transfer := &utils.Transfer{Conn: Conn}

// 	for {
// 		mes, err := transfer.RecvMessage()
// 		if err != nil {
// 			fmt.Println("Error: ", err)
// 			return
// 		}
// 		fmt.Println("Received message: ", mes.Type, mes.Data)
// 	}
// }
