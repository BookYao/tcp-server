package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"
)

const (
	MSG_HEAD_LENGHT = 24
)

const (
	MSG_PING                = 1
	MSG_SET_LINKAGE_WATCHER = 330
)

type MSGHead struct {
	Length    int
	TermID    int
	CommendID int
	Type      int
	Seqno     int
	Mark      int
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	msgHeadBuf := make([]byte, MSG_HEAD_LENGHT)
	for {
		length, err := conn.Read(msgHeadBuf)
		if err != nil || length != MSG_HEAD_LENGHT {
			fmt.Println("Read MSG MsgHead Failed.", err.Error())
			continue
		}

		MSGHead := MSGHead{}
		bytesBuf := bytes.NewBuffer(msgHeadBuf)
		binary.Read(bytesBuf, binary.LittleEndian, &MSGHead)
		fmt.Printf("THis MSG Msg Total Length: %d\n", MSGHead.Length)

		MSGMsgHandle(conn, &MSGHead)
	}
}

func saveLinkageWatcher(msg []byte) bool {
	LINKAGE_WATCHER_FILE := "/var/tmp/linkage.conf"
	fp, err := os.Open(LINKAGE_WATCHER_FILE)
	if err != nil {
		fmt.Println("Open Linkage Watcher File Failed.", err)
		return false
	}

	_, err = fp.Write(msg)
	if err != nil {
		fp.Close()
		fmt.Println("Write Linkage Watcher File Failed.", err)
		return false
	}
	fp.Close()
	return true
}

func MSGMsgHandle(conn net.Conn, MSGHead *MSGHead) bool {
	var status bool = true
	msgBodyLen := MSGHead.Length - MSG_HEAD_LENGHT
	switch MSGHead.Type {
	case MSG_PING:
		buf := make([]byte, msgBodyLen)
		len, err := conn.Read(buf)
		if err != nil || len == 0 {
			fmt.Println("Set Linkage Watcher Read failed.", err)
			status = false
		}
		_, err = conn.Write(buf)
		if err != nil {
			fmt.Println("MSG Ping Write Failed.", err)
			status = false
		}
	case MSG_SET_LINKAGE_WATCHER:
		buf := make([]byte, msgBodyLen)
		len, err := conn.Read(buf)
		if err != nil || len == 0 {
			fmt.Println("Set Linkage Watcher Read failed.", err)
			status = false
		}
		fmt.Printf("Recv Set Linkage Watcher MSG, bodylen:%d-recvlen:%d\n", msgBodyLen, len)
		if saveLinkageWatcher(buf) != true {
			fmt.Println("Save Linkage Watcher Failed.")
			status = false
		}
	default:
		fmt.Println("Unknown MSG Msg type.")
	}

	return status
}
