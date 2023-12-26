

// /Users/ir/Desktop/p2p_go/p2p_go/utils.go
package main

import (
	"log"
	"net"
)

// CheckErrorはエラーが発生したかどうかを確認し、発生した場合はログに記録します。
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// GetLocalAddressはホストのローカルアドレスを返します。
func GetLocalAddress() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}

	// ネットワークインターフェースアドレスのリストから、ループバックでないIPv4アドレスを探します。
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	// 適切なアドレスが見つからない場合は空の文字列を返します。
	return ""
}
