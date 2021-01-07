package main

import "testing"

func TestSendRtx(t *testing.T) {
	var typ = 100
	var receiver = []string{"WayneWang"}
	var title = "rtx测试"
	var content = "test"

	SendRtx(typ, receiver, title, content)
}
