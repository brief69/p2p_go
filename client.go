

// /Users/ir/Desktop/p2p_go/p2p_go/client.go
package main

import (
	"log"
	"peer"
	"config"
)

// Clientはp2pネットワークのクライアントを表します。
type Client struct {
	Address string
	Peer    *peer.Peer
}

// NewClient creates a new Client.
func NewClient(cfg *config.Config) (*Client, error) {
	p, err := peer.NewPeer(cfg.PeerAddress)
	if err != nil {
		return nil, err
	}

	client := &Client{
		Address: cfg.ServerAddress,
		Peer:    p,
	}

	return client, nil
}

// Start starts the client.
func (c *Client) Start() {
	log.Printf("Client started on %s", c.Address)

	// Here you can add code to send/receive messages to/from the peer
	// For example:
	// err := c.Peer.Send([]byte("Hello, peer!"))
	// if err != nil {
	// 	log.Printf("Failed to send message to peer %s: %s", c.Peer.Address, err)
	// }
	//
	// data, err := c.Peer.Receive()
	// if err != nil {
	// 	log.Printf("Failed to receive message from peer %s: %s", c.Peer.Address, err)
	// } else {
	// 	log.Printf("Received message from peer %s: %s", c.Peer.Address, string(data))
	// }
}

// Stop stops the client.
func (c *Client) Stop() {
	log.Println("Stopping client...")

	c.Peer.Close()

	log.Println("Client stopped")
}
