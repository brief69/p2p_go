

// /Users/ir/Desktop/p2p_go/p2p_go/server.go
package main

import (
	"log"
	"net"
	"peer"
)

// Serverはp2pネットワークのサーバーを表します。
type Server struct {
	Address  string
	Listener net.Listener
	Peers    []*peer.Peer
}

// NewServerは新しいServerを作成します。
func NewServer(address string) (*Server, error) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	server := &Server{
		Address:  address,
		Listener: listener,
		Peers:    []*peer.Peer{},
	}

	return server, nil
}

// Startはサーバーを開始します。
func (s *Server) Start() {
	log.Printf("サーバーが開始しました %s", s.Address)

	for {
		conn, err := s.Listener.Accept()
		if err != nil {
			log.Printf("接続の受け入れに失敗しました: %s", err)
			continue
		}

		peer := &peer.Peer{
			Address: conn.RemoteAddr().String(),
			Conn:    conn,
		}
		s.Peers = append(s.Peers, peer)

		go s.handlePeer(peer)
	}
}

// handlePeerはピアとの通信を処理します。
func (s *Server) handlePeer(p *peer.Peer) {
	for {
		data, err := p.Receive()
		if err != nil {
			log.Printf("ピアからのデータの受信に失敗しました %s: %s", p.Address, err)
			p.Close()
			s.removePeer(p)
			break
		}

		log.Printf("ピアからデータを受信しました %s: %s", p.Address, string(data))
	}
}

// removePeerはピアをサーバーから削除します。
func (s *Server) removePeer(p *peer.Peer) {
	for i, peer := range s.Peers {
		if peer == p {
			s.Peers = append(s.Peers[:i], s.Peers[i+1:]...)
			break
		}
	}
}

// Stopはサーバーを停止します。
func (s *Server) Stop() {
	log.Println("サーバーを停止中...")

	for _, peer := range s.Peers {
		peer.Close()
	}

	err := s.Listener.Close()
	if err != nil {
		log.Printf("サーバーのクローズに失敗しました: %s", err)
	}

	log.Println("サーバーが停止しました")
}
