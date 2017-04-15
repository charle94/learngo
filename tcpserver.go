package learn

import (
	"fmt"
	"net"
)

func Server() {
	l, e := net.Listen("tcp", ":8090")
	if e != nil {
		fmt.Println(e)
	}
	for {
		conn, e := l.Accept()
		if e != nil {
			fmt.Println("接收数据异常", e.Error())
			continue
		}
		defer conn.Close()
		go func() {
			data := make([]byte, 128)
			for {
				i, e := conn.Read(data)
				fmt.Println("数据为", string(data[0:i]))
				if e != nil {
					fmt.Println("读取数据错误", e.Error())
					break
				}
			}
			conn.Write([]byte{'f', 'i', 'n', 'i', 's', 'h'})
		}()

	}

}
func Client() {
	conn, err := net.Dial("tcp", ":8090")
	if err != nil {
		fmt.Println("连接错误", err.Error())
	}
	ms := []byte("yes we are")
	defer conn.Close()
	conn.Write(ms)

}
