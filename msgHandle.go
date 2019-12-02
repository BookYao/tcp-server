package tcp_server

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

const (
	AACP_HEAD_LENGHT = 24
)

const (
	AACP_SET_LINKAGE_WATCHER = 330
)

type AACPHead struct  {
	Length int
	TermID int
	CommendID int
	Type int
	Seqno int
	Mark int
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	msgHeadBuf := make([]byte, AACP_HEAD_LENGHT)
	for {
		length, err := conn.Read(msgHeadBuf)
		if err != nil  || length != AACP_HEAD_LENGHT {
			fmt.Println("Read AACP MsgHead Failed.", err.Error())
			continue
		}

		aacpHead := AACPHead{}
		bytesBuf := bytes.NewBuffer(msgHeadBuf)
		binary.Read(bytesBuf, binary.LittleEndian, &aacpHead)
		fmt.Printf("THis AACP Msg Total Length: %d\n", aacpHead.Length)

		aacpMsgHandle(conn, &aacpHead)
	}
}

func aacpMsgHandle(conn net.Conn, aacpHead *AACPHead)  {
	msgBodyLen := aacpHead.Length - AACP_HEAD_LENGHT
	switch aacpHead.Type {
	case AACP_SET_LINKAGE_WATCHER:
		fmt.Println("Recv Set Linkage Watcher MSG, bodylen:", msgBodyLen)
	default:
		fmt.Println("Unknown AACP Msg type.")
	}
}