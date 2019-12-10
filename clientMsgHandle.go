package main

import (
	"bytes"
	"encoding/binary"
)

const (
	MSG_PING                = 1
	MSG_SET_LINKAGE_WATCHER = 330
)

type ClientMsgHead struct {
	Length    int
	TermID    int
	CommendID int
	Type      int
	Seqno     int
	Mark      int
}

func pingMsgBuild() []byte {
	data := ClientMsgHead{}
	data.CommendID = MSG_PING
	data.Length = 24
	data.Mark = 1
	data.Seqno = 0
	data.Type = 0
	data.TermID = 0

	buf := bytes.Buffer{}
	binary.Write(&buf, binary.BigEndian, uint32(data.CommendID))
	binary.Write(&buf, binary.BigEndian, uint32(data.Length))
	binary.Write(&buf, binary.BigEndian, uint32(data.Mark))
	binary.Write(&buf, binary.BigEndian, uint32(data.Seqno))
	binary.Write(&buf, binary.BigEndian, uint32(data.Type))
	binary.Write(&buf, binary.BigEndian, uint32(data.TermID))

	return buf.Bytes()
}
