package main

import(
	"fmt"
	"net"
)

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
			fmt.Println("Accept() conn=",conn,conn.RemoteAddr().String())
		}

	}

	// fmt.Println("listen = ",listen)
}