package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	var remoteaddr = "192.168.1.100:80"
	var configaddr = "192.168.1.1/24"

	fmt.Printf("remoteaddr: %s\n", remoteaddr)
	fmt.Printf("configaddr: %s\n", configaddr)

	addrslice := strings.Split(remoteaddr, ":")
	fmt.Printf("addr: %s\n", addrslice[0])
	cip, cnet, err := net.ParseCIDR(configaddr)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	fmt.Printf("cip: %s\n", cip.String())
	fmt.Printf("cnet: %s\n", cnet.String())
	rip := net.ParseIP(addrslice[0])
	fmt.Printf("raddr: %s\n", rip.String())
	if cnet.Contains(rip) {
		fmt.Printf("rip contained cnet.\n")
		return
	}
	fmt.Printf("faild contained check.\n")
	return
}
