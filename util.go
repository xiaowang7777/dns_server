package main

import (
	"fmt"
	"net"
	"regexp"
	"strings"
)

var regex, _ = regexp.Compile("Address: .*")

func handleRemoteIp(data []byte) net.IP {
	all := regex.FindAll(data, -1)
	var ips []net.IP
	for _, bytes := range all {
		ip := strings.Trim(string(bytes), "Address: ")

		fmt.Println(ip)

		parseIP := net.ParseIP(ip)
		if parseIP != nil {
			ips = append(ips, parseIP)
		}
	}

	if len(ips) > 0 {
		return ips[0]
	}
	return nil
}
