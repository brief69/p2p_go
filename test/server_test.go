

// /Users/ir/Desktop/p2p_go/p2p_go/test/server_test.go
package server_test

import (
	"testing"
	"net"
	"time"
	"./server"
	"./peer"
)

func TestNewServer(t *testing.T) {
	srv, err := server.NewServer("localhost:8080")
	if err != nil {
		t.Fatalf("Failed to create server: %s", err)
	}

	if srv.Address != "localhost:8080" {
		t.Errorf("Expected server address to be 'localhost:8080', got '%s'", srv.Address)
	}

	if srv.Listener == nil {
		t.Error("Expected server listener to be initialized, got nil")
	}

	if len(srv.Peers) != 0 {
		t.Errorf("Expected server to have 0 peers, got %d", len(srv.Peers))
	}
}

func TestServerStartAndStop(t *testing.T) {
	srv, err := server.NewServer("localhost:8080")
	if err != nil {
		t.Fatalf("Failed to create server: %s", err)
	}

	go srv.Start()
	time.Sleep(time.Second) // give server time to start

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		t.Fatalf("Failed to connect to server: %s", err)
	}

	peer := &peer.Peer{
		Address: conn.RemoteAddr().String(),
		Conn:    conn,
	}

	srv.Peers = append(srv.Peers, peer)

	if len(srv.Peers) != 1 {
		t.Errorf("Expected server to have 1 peer, got %d", len(srv.Peers))
	}

	srv.Stop()
	time.Sleep(time.Second) // give server time to stop

	if len(srv.Peers) != 0 {
		t.Errorf("Expected server to have 0 peers after stop, got %d", len(srv.Peers))
	}
}