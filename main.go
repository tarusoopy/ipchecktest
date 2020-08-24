package main

import (
	"fmt"
	"net"
)

func main() {
	var remoteaddr = "192.168.1.100:80"
	var configaddr = "192.168.1.1/24"

	_, cnet, _ := net.ParseCIDR(configaddr)
	fmt.Printf("cnet: %s\n", cnet.Network())
	rip := net.ParseIP(remoteaddr)
	fmt.Printf("raddr: %s\n", rip.String())
	if cnet.Contains(rip) {
		fmt.Printf("rip contained cnet.\n")
		return
	}
	fmt.Printf("faild contained check.\n")
	return
}
