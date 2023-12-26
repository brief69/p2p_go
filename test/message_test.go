

// /Users/ir/Desktop/p2p_go/p2p_go/test/message_test.go
package message_test

import (
	"testing"

	"../message"
)

func TestNewMessage(t *testing.T) {
	msg, err := message.NewMessage("Alice", "Bob", "Hello, Bob!")
	if err != nil {
		t.Fatalf("Failed to create new message: %s", err)
	}

	if msg.From != "Alice" || msg.To != "Bob" || msg.Content != "Hello, Bob!" {
		t.Fatalf("Message fields are incorrect")
	}
}

func TestEncode(t *testing.T) {
	msg, _ := message.NewMessage("Alice", "Bob", "Hello, Bob!")
	data, err := msg.Encode()
	if err != nil {
		t.Fatalf("Failed to encode message: %s", err)
	}

	expected := `{"from":"Alice","to":"Bob","content":"Hello, Bob!"}`
	if string(data) != expected {
		t.Fatalf("Encoded message is incorrect")
	}
}

func TestDecodeMessage(t *testing.T) {
	data := []byte(`{"from":"Alice","to":"Bob","content":"Hello, Bob!"}`)
	msg, err := message.DecodeMessage(data)
	if err != nil {
		t.Fatalf("Failed to decode message: %s", err)
	}

	if msg.From != "Alice" || msg.To != "Bob" || msg.Content != "Hello, Bob!" {
		t.Fatalf("Decoded message fields are incorrect")
	}
}
```
