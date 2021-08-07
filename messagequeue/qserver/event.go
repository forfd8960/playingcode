package qserver

import "sync"

type Event struct {
	qsrv *QServer

	mu       sync.Mutex
	Name     string
	channels map[string]*Channel
}
