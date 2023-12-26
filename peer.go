

// /Users/ir/Desktop/p2p_go/p2p_go/peer.go
package main

import (
	"io"
	"log"
	"net"
)

// Peerはp2pネットワークのピアを表します。
type Peer struct {
	Address string
	Conn    net.Conn
}

// NewPeerは新しいPeerを作成します。
func NewPeer(address string) (*Peer, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}

	peer := &Peer{
		Address: address,
		Conn:    conn,
	}

	return peer, nil
}

// Sendはピアにメッセージを送信します。
func (p *Peer) Send(data []byte) error {
	_, err := p.Conn.Write(data)
	if err != nil {
		return err
	}

	return nil
}

// Receiveはピアからメッセージを受信します。
func (p *Peer) Receive() ([]byte, error) {
	buf := make([]byte, 1024)
	n, err := p.Conn.Read(buf)
	if err != nil {
		if err != io.EOF {
			return nil, err
		}
	}

	return buf[:n], nil
}

// Closeはピアへの接続を閉じます。
func (p *Peer) Close() {
	err := p.Conn.Close()
	if err != nil {
		log.Printf("ピア%sへの接続のクローズに失敗しました: %s", p.Address, err)
	}
}
