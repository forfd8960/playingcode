package qserver

import (
	"net"
	"sync"
)

type Handler interface {
	Handle(conn net.Conn)
}

type QServer struct {
	mu      sync.Mutex
	events  map[string]*Event // eventName -> event
	handler Handler
}
