

// /Users/ir/Desktop/p2p_go/p2p_go/test/main_test.go
package main_test

import (
	"testing"
	"os"
	"os/signal"
	"syscall"
	"time"

	"../main"
	"../config"
	"../server"
	"../client"
)

func TestMain(t *testing.T) {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		t.Fatalf("Failed to load configuration: %s", err)
	}

	// Initialize server
	srv, err := server.NewServer(cfg)
	if err != nil {
		t.Fatalf("Failed to initialize server: %s", err)
	}

	// Initialize client
	cli, err := client.NewClient(cfg)
	if err != nil {
		t.Fatalf("Failed to initialize client: %s", err)
	}

	// Start server and client
	go srv.Start()
	go cli.Start()

	// Wait for a while to let server and client start
	time.Sleep(2 * time.Second)

	// Send interrupt signal to stop server and client
	p, err := os.FindProcess(os.Getpid())
	if err != nil {
		t.Fatalf("Failed to find process: %s", err)
	}
	p.Signal(os.Interrupt)

	// Wait for a while to let server and client stop
	time.Sleep(2 * time.Second)

	// Check if server and client have stopped
	if srv.Running || cli.Running {
		t.Fatalf("Server or client did not stop")
	}
}
```
