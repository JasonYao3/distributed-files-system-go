package p2p

// Peer represents the remote node
type Peer interface {
	Close() error
}

// Transport handles the communication between the nodes in the network
// This can be of the form (TCP, UDP, websocket)
type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
}
