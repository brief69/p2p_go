// /Users/ir/Desktop/p2p_go/p2p_go/config.go
package main

import (
	"encoding/json"
	"os"
)

// Configは設定情報を表します。
type Config struct {
	ServerAddress string `json:"server_address"` // サーバーのアドレス
	PeerAddress   string `json:"peer_address"`   // ピアのアドレス
}

// LoadConfigはJSONファイルから設定を読み込みます。
func LoadConfig() (*Config, error) {
	file, err := os.Open("config.json") // config.jsonファイルを開く
	if err != nil {
		return nil, err
	}
	defer file.Close() // 関数が終了するときにファイルを閉じる

	decoder := json.NewDecoder(file) // JSONデコーダを作成
	config := &Config{}              // 新しいConfigを作成
	err = decoder.Decode(config)     // JSONをConfigにデコード
	if err != nil {
		return nil, err
	}

	return config, nil // 読み込んだ設定を返す
}
