// 负责将IP地址写入到输出文件中
// write.go
package main

import (
	"os"
	"sync"
)

// writeOutput writes the output content to a channel and the output file
func writeOutput(outputChan <-chan string, output string, wg *sync.WaitGroup) {
	defer wg.Done()  // Decrease the wait group counter when done
	// Open the output file and defer its closing
	out, err := os.OpenFile(output, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) // Open an existing file and truncate it, or create a new file
	checkError(err, "Failed to open the output file")
	defer out.Close()
	defer out.Sync() // Force flush to disk when done
	// Iterate over the output channel and write the IP addresses to the file
	for ip := range outputChan {
		// Write the IP address to the file
		out.WriteString(ip + "\n")
	}
}
