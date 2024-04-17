package p2p

import "net"

// Peer represents the remote node
type Peer interface {
	net.Conn
	Send([]byte) error
}

// Transport handles the communication between the nodes in the network
// This can be of the form (TCP, UDP, websocket)
type Transport interface {
	Dial(string) error
	ListenAndAccept() error
	Consume() <-chan RPC
	Close() error
}
