

// /Users/ir/Desktop/p2p_go/p2p_go/test/client_test.go
package client_test

import (
	"testing"
	"../client"
	"../config"
	"../utils"
)

func TestNewClient(t *testing.T) {
	cfg, err := config.LoadConfig()
	utils.CheckError(err)

	c, err := client.NewClient(cfg)
	utils.CheckError(err)

	if c.Address != cfg.ServerAddress {
		t.Errorf("Expected client address to be %s, but got %s", cfg.ServerAddress, c.Address)
	}

	if c.Peer.Address != cfg.PeerAddress {
		t.Errorf("Expected peer address to be %s, but got %s", cfg.PeerAddress, c.Peer.Address)
	}
}

func TestStartAndStopClient(t *testing.T) {
	cfg, err := config.LoadConfig()
	utils.CheckError(err)

	c, err := client.NewClient(cfg)
	utils.CheckError(err)

	c.Start()
	c.Stop()

	// Here you can add more tests to check if the client is correctly started and stopped
	// For example, you can check if the client is able to send/receive messages to/from the peer
	// You can also check if the client is correctly stopped (e.g., the connection to the peer is closed)
}