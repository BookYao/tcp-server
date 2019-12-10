package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: %s serverIp serverPort", os.Args[0])
		os.Exit(0)
	}

	serverAddr := os.Args[1] + ":" + os.Args[2]
	fmt.Println("serverAddr:", serverAddr)
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("TCP Connect Failed!")
		os.Exit(0)
	}
	defer conn.Close()

	pingMsg := pingMsgBuild()
	conn.Write(pingMsg)
}
