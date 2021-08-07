package qserver

type Message struct {
	ID       string
	Body     []byte
	clientID int64
}
