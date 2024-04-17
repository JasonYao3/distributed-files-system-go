package p2p

import "net"

// Peer represents the remote node
type Peer interface {
	Send([]byte) error
	RemoteAddr() net.Addr
	Close() error
}

// Transport handles the communication between the nodes in the network
// This can be of the form (TCP, UDP, websocket)
type Transport interface {
	Dial(string) error
	ListenAndAccept() error
	Consume() <-chan RPC
	Close() error
}
