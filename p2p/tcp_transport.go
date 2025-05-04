package p2p

import (
	"fmt"
	"net"
	"sync"
)

// <------------------ TCPPeer Section - START ------------------>

// TCPPeer represents a remote node connected over TCP
type TCPPeer struct {
	// connection is the underlying TCP connection with the peer
	connection net.Conn

	// outbound indicates whether the TCP connection is inbound or outbound
	// if we dialed into the peer -> outbound == true
	// if we accept a peer dialing in -> outbound == false
	outbound bool
}

// TCPPeerOption is a functional option type that allows us to configure the TCPPeer
type TCPPeerOption func(*TCPPeer)

func NewTCPPeer(options ...TCPPeerOption) *TCPPeer {
	peer := &TCPPeer{}

	for _, option := range options {
		option(peer)
	}

	return peer
}

func WithConnection(connection net.Conn) TCPPeerOption {
	return func(peer *TCPPeer) {
		peer.connection = connection
	}
}

func WithOutbound(outbound bool) TCPPeerOption {
	return func(peer *TCPPeer) {
		peer.outbound = outbound
	}
}

// <------------------ TCPPeer Section - END ------------------>

// <------------------ TCPTransport Section - START ------------------>

// TCPTransport handles communication with peers over TCP
type TCPTransport struct {
	listenAddress string
	listener      net.Listener

	mutex sync.RWMutex
	peers map[net.Addr]Peer
}

// TCPTransportOption is a functional option type that allows us to configure the TCPTransport
type TCPTransportOption func(*TCPTransport)

func NewTCPTransport(options ...TCPTransportOption) *TCPTransport {
	transport := &TCPTransport{
		peers: make(map[net.Addr]Peer),
	}

	for _, option := range options {
		option(transport)
	}

	return transport
}

func WithListenAddress(listenAddress string) TCPTransportOption {
	return func(transport *TCPTransport) {
		transport.listenAddress = listenAddress
	}
}

func (transport *TCPTransport) ListenAndAccept() error {
	var err error
	transport.listener, err = net.Listen("tcp", transport.listenAddress)
	if err != nil {
		return err
	}

	go transport.startAcceptLoop()

	return nil
}

func (transport *TCPTransport) startAcceptLoop() {
	for {
		connection, err := transport.listener.Accept()
		if err != nil {
			fmt.Printf("Error at TCP accept: %s\n", err)
		}

		go transport.handleConnection(connection)
	}
}

func (transport *TCPTransport) handleConnection(connection net.Conn) {
	peer := NewTCPPeer(
		WithConnection(connection),
		WithOutbound(false),
	)

	fmt.Printf("Handling Connection: %+v\n", peer)
}

// <------------------ TCPTransport Section - END ------------------>
