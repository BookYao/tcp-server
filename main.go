package tcp_server

import (
	"fmt"
	"net"
	"os"
	"runtime"
)

func main() {
	fmt.Println("TCP Server Start...")
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s Port", os.Args[0])
		os.Exit(0)
	}

	service := ":7090"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkerr(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkerr(err)

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("Accept TCP connect failed.")
		}
		go handleConn(conn)
	}
}

func getFileNameLine() string {
	funcName, file, line, ok := runtime.Caller(0)
	if ok {
		return fmt.Sprintf("%s:%s:%d", file, runtime.FuncForPC(funcName).Name(), line)
	}
	return ""
}

func checkerr(err error) {
	if err != nil {
		fmt.Printf("%s. strerror:%s", getFileNameLine(), err.Error())
		os.Exit(0)
	}
}