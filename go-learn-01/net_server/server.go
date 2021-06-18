package main

import (
	"bufio"
	"fmt"
	"net"
)

// tcp 练习
func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed", err)
		return
	}
	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn) // 启动一个goroutine处理连接
	}
}

// 处理函数
func process(conn net.Conn) {
	defer conn.Close() // defer 关闭连接
	for {
		var buf [128]byte
		reader := bufio.NewReader(conn)
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read from client failed err : ", err)
			break
		}
		receStr := string(buf[:n])
		fmt.Println("from client data", receStr)
		conn.Write([]byte(receStr))
	}

}
