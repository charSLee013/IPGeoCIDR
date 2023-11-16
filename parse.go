// 负责解析CIDR并存储到一个切片中
// parse.go
package main

import (
	"bufio"
	"net"
	"os"
	"sync"
)

// parseCIDR parses the CIDR and sends IP addresses to a channel
func parseCIDR(cidr string, ipChan chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()     // Decrease the wait group counter when done
	defer close(ipChan) // Close the ip channel when done
	_, _, err := net.ParseCIDR(cidr)
	if err == nil {
		// The cidr is a valid CIDR, send it to the channel
		ipChan <- cidr
	} else {
		// The cidr is a file, open it and read it line by line
		file, err := os.Open(cidr)
		checkError(err, "Failed to open the CIDR file")
		defer file.Close()
		reader := bufio.NewScanner(file)
		reader.Split(bufio.ScanLines)
		for reader.Scan() {
			ipChan <- reader.Text()
		}
	}
}
