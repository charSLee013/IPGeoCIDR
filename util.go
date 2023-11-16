// 辅助函数
// util.go
package main

import (
	"log"
	"net"
	"os"
)

// inc increments an IP address and returns a new IP address
func inc(ip net.IP) net.IP {
	// Copy the IP to a new slice
	inc := make(net.IP, len(ip))
	copy(inc, ip)
	// Increment the IP from right to left
	for j := len(inc) - 1; j >= 0; j-- {
		inc[j]++
		// If the byte is not zero, break the loop
		if inc[j] > 0 {
			break
		}
	}
	return inc
}

// checkError checks if the error is not nil and prints the message and exits if it is
func checkError(err error, message string) {
	if err != nil {
		log.Println(message)
		log.Println(err)
		os.Exit(1)
	}
}
