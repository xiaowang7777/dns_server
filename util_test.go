package main

import (
	"fmt"
	"testing"
)

var out = `
Server:         127.0.0.53
Address:        127.0.0.53#53

Non-authoritative answer:
*** Can't find yutube.com: No answer
`

var out1 = `
Server:         127.0.0.53
Address:        127.0.0.53#53

Non-authoritative answer:
Name:   youtube.com
Address: 142.250.204.46
Name:   youtube.com
Address: 2404:6800:4012:9::200e
`

func TestHandleRemoteIp(t *testing.T) {
	fmt.Println(handleRemoteIp([]byte(out)))

	fmt.Println("#########")

	fmt.Println(handleRemoteIp([]byte(out1)))
}
