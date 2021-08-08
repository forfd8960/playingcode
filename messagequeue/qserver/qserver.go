package qserver

import (
	"net"
	"sync"
)

type Options struct {
	TCPAddress string
}

type Handler interface {
	Handle(conn net.Conn)
	Close() error
}

type QServer struct {
	mu      sync.Mutex
	events  map[string]*Event // eventName -> event
	handler Handler

	listner net.Listener
	wg      sync.WaitGroup
}

func New(opts *Options) (*QServer, error) {
	var err error

	qsrv := &QServer{}
	qsrv.listner, err = net.Listen("tcp", opts.TCPAddress)
	if err != nil {
		return nil, err
	}

	return qsrv, nil
}

func (qs *QServer) Serve() error {
	for {
		conn, err := qs.listner.Accept()
		if err != nil {
			return err
		}

		qs.wg.Add(1)
		go func() {
			qs.handler.Handle(conn)
			qs.wg.Done()
		}()
	}
}

func (qs *QServer) Exit() {
	if qs.listner != nil {
		qs.listner.Close()
	}

	qs.wg.Wait()
}
