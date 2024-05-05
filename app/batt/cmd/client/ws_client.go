package main

import (
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: "127.0.0.1:8001", Path: "/test"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	// 创建数据
	data := []byte(`{"head": "test"}`)
	packetLength := uint32(len(data) + 4) // 数据长度 + 4字节的包长度字段

	// 构建消息包
	packet := make([]byte, 4+len(data))
	packet[0] = byte(0x01)                        // 设置类型（假设为0x01）
	packet[1] = byte(packetLength & 0xFF)         // 设置包长度低位字节
	packet[2] = byte((packetLength >> 8) & 0xFF)  // 设置包长度中位字节
	packet[3] = byte((packetLength >> 16) & 0xFF) // 设置包长度高位字节
	copy(packet[4:], data)                        // 设置数据

	// 发送消息
	err = c.WriteMessage(websocket.BinaryMessage, packet)
	if err != nil {
		log.Println("write:", err)
		return
	}

	// 等待接收服务器的响应
	go func() {
		defer c.Close()
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("received: %s", message)
		}
	}()

	// 等待中断信号来关闭连接
	select {
	case <-interrupt:
		log.Println("interrupt")
		// 发送关闭消息
		err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		if err != nil {
			log.Println("write close:", err)
			return
		}
		select {
		case <-time.After(time.Second):
		}
		return
	}
}
