package server

import "net"

type CmdType int

const (
	REGISTER CmdType = iota
	JOIN
	LEAVE
	MSG
	CHANNELS
	USERS
)

// Room
type Room struct {
}

type Client struct {
	conn net.Conn
}

type Channel struct {
}

type Command struct {
	CType CmdType
}
