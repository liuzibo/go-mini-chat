package utils

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go-mini-chat/internal/common/message"
	"net"
)

type Transfer struct {
	Conn net.Conn
}

func (transfer *Transfer) SendMessage(msg *message.Message) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	dataLength := int32(len(data))
	pkg := new(bytes.Buffer)
	err = binary.Write(pkg, binary.LittleEndian, dataLength)
	if err != nil {
		return err
	}

	err = binary.Write(pkg, binary.LittleEndian, data)
	if err != nil {
		return err
	}

	_, err = transfer.Conn.Write(pkg.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func (transfer *Transfer) RecvMessage() (*message.Message, error) {
	// 前四个字节为数据包长度
	reader := bufio.NewReader(transfer.Conn)
	length := make([]byte, 4)
	_, err := reader.Read(length)
	if err != nil {
		return nil, err
	}
	dataLength := binary.LittleEndian.Uint32(length)

	if reader.Buffered() < int(dataLength) {
		return nil, err
	}
	data := make([]byte, int(dataLength))
	_, err = reader.Read(data)
	if err != nil {
		return nil, err
	}
	msg := new(message.Message)
	err = json.Unmarshal(data, msg)
	if err != nil {
		fmt.Println("json unmarshal error:", err)
		return nil, err
	}
	return msg, nil
}
