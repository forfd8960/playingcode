package qserver

import "sync"

type Consumer interface {
}

type Channel struct {
	qsrv *QServer

	mu        sync.Mutex
	eventName string
	name      string
	msgChan   chan Message
	consumers map[int64]Consumer
}
