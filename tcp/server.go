package main

import (
	"fmt"
	"io"
	"net"
)
func process(conn net.Conn){
	defer conn.Close()
	// 循环接收某个客户端发送的数据
	for {
		// 创建一个新的切片
		buf := make([]byte,1024)
		// 等待客户端通过conn发送信息，如果客户端没有Write，那么协程就会阻塞
		// fmt.Printf("等待客户端 %v 发送数据 \n",conn.RemoteAddr().String())
		n,err := conn.Read(buf)
		if err == io.EOF {
			fmt.Printf("客户端 %v 已退出\n",conn.RemoteAddr().String())
			return
		}	else if err != nil {
			fmt.Println("客户端出现错误",err)
			return
		}
		// 显示客户端发送的内容在服务端的终端
		fmt.Print(string(buf[:n]))
	}
}

func main(){
	fmt.Println("服务器开始监听...")
	listen, err := net.Listen("tcp","127.0.0.1:8888")
	if err != nil {
		fmt.Println("listen err = ",err)
		return
	}
	defer listen.Close()

	// 循环等待客户端来连接
	for {
		conn,err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=",err)
		} else {
			// conn.RemoteAddr()表示客户端的接口
			fmt.Printf("客户端 %v 已连接 \n",conn.RemoteAddr().String())
		}
		go process(conn)
	}

	// fmt.Println("listen = ",listen)
}