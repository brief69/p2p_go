

// /Users/ir/Desktop/p2p_go/p2p_go/test/utils_test.go
package utils_test

import (
	"testing"
	"net"
	"../utils"
)

func TestCheckError(t *testing.T) {
	err := net.UnknownNetworkError("unknown network error")
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	utils.CheckError(err)
}

func TestGetLocalAddress(t *testing.T) {
	addr := utils.GetLocalAddress()
	if addr == "" {
		t.Errorf("Failed to get local address")
	}
}