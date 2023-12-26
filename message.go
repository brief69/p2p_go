

// /Users/ir/Desktop/p2p_go/p2p_go/message.go
package main

import (
	"encoding/json"
	"errors"
)

// Messageはp2pネットワークのメッセージを表します。
type Message struct {
	From    string `json:"from"`    // 送信元のアドレス
	To      string `json:"to"`      // 宛先のアドレス
	Content string `json:"content"` // メッセージの内容
}

// NewMessageは新しいMessageを作成します。
func NewMessage(from, to, content string) (*Message, error) {
	if from == "" || to == "" || content == "" {
		return nil, errors.New("無効なメッセージパラメータ")
	}

	msg := &Message{
		From:    from,
		To:      to,
		Content: content,
	}

	return msg, nil
}

// EncodeはメッセージをJSON文字列にエンコードします。
func (m *Message) Encode() ([]byte, error) {
	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// DecodeMessageはJSON文字列をMessageにデコードします。
func DecodeMessage(data []byte) (*Message, error) {
	var msg Message
	err := json.Unmarshal(data, &msg)
	if err != nil {
		return nil, err
	}

	return &msg, nil
}