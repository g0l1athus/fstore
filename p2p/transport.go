package p2p

// Peer represents a remote node.
type Peer interface{}

// Transport represents a handler for a communication channel
// between peers. This can be of type TCP, UDP, Websockets, gRPC, ...
type Transport interface {
	ListenAndAccept() error
}
