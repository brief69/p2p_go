

// /Users/ir/Desktop/p2p_go/p2p_go/main.go
package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"config"
	"server"
	"client"
)

func main() {
	// 設定を読み込む
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("設定の読み込みに失敗しました: %s", err)
	}

	// Initialize server
	srv, err := server.NewServer(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize server: %s", err)
	}

	// Initialize client
	cli, err := client.NewClient(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize client: %s", err)
	}

	// Start server and client
	go srv.Start()
	go cli.Start()

	// Wait for interrupt signal to gracefully shutdown the server and client
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	log.Println("Shutting down server and client...")

	srv.Stop()
	cli.Stop()

	log.Println("Server and client successfully shutdown")
}
