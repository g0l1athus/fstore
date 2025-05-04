package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	listenAddress := ":4991"
	transport := NewTCPTransport(
		WithListenAddress(listenAddress),
	)

	assert.Equal(t, listenAddress, transport.listenAddress)

	assert.Nil(t, transport.ListenAndAccept())
}
