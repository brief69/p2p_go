

// /Users/ir/Desktop/p2p_go/p2p_go/test/peer_test.go
package peer_test

import (
	"testing"
	"net"
	"github.com/stretchr/testify/assert"
	"peer"
)

func TestNewPeer(t *testing.T) {
	listener, _ := net.Listen("tcp", "127.0.0.1:0")
	defer listener.Close()

	go func() {
		conn, _ := listener.Accept()
		conn.Close()
	}()

	peer, err := peer.NewPeer(listener.Addr().String())
	assert.NoError(t, err)
	assert.NotNil(t, peer)
}

func TestSendAndReceive(t *testing.T) {
	listener, _ := net.Listen("tcp", "127.0.0.1:0")
	defer listener.Close()

	go func() {
		conn, _ := listener.Accept()
		conn.Write([]byte("Hello, peer!"))
		conn.Close()
	}()

	peer, _ := peer.NewPeer(listener.Addr().String())
	defer peer.Close()

	err := peer.Send([]byte("Hello, server!"))
	assert.NoError(t, err)

	data, err := peer.Receive()
	assert.NoError(t, err)
	assert.Equal(t, "Hello, peer!", string(data))
}

func TestClose(t *testing.T) {
	listener, _ := net.Listen("tcp", "127.0.0.1:0")
	defer listener.Close()

	peer, _ := peer.NewPeer(listener.Addr().String())
	peer.Close()

	_, err := peer.Receive()
	assert.Error(t, err)
}
