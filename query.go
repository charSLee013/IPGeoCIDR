// 负责查询每个IP地址所属的国家
// query.go
package main

import (
	"net"
	"regexp"
	"sync"

	"fmt"
	"sync/atomic"
)

// 添加两个原子计数器，用于追踪IP地址的数量
var (
	totalIPs   int64 = 0
	matchedIPs int64 = 0
)

// queryCountry queries the country for each IP address and sends matched countries and IP addresses to a channel
func queryCountry(ipChan <-chan string, countryChan chan<- string, countryRegex *regexp.Regexp, wg *sync.WaitGroup, workerCount int) {
	defer wg.Done()          // Decrease the wait group counter when done
	defer close(countryChan) // Close the country channel when done

	// full usage
	var workerWG sync.WaitGroup
	for i := 0; i < workerCount; i++ {
		workerWG.Add(1)
		go func() {
			defer workerWG.Done()
			// Iterate over the ip channel and query the country for each CIDR
			for c := range ipChan {
				// Parse the CIDR and get the IP range
				ip, ipnet, err := net.ParseCIDR(c)
				if err != nil {
					// Invalid CIDR, skip it
					fmt.Printf("查询错误: %s\n", err)
					continue
				}
				// Get the first and last IP in the range
				firstIP := ip.Mask(ipnet.Mask)
				lastIP := net.IP(make([]byte, len(firstIP)))
				for i := range firstIP {
					lastIP[i] = firstIP[i] | ^ipnet.Mask[i]
				}
				// Iterate over the IP range and query the country for each IP
				for ip := firstIP; ipnet.Contains(ip); ip = inc(ip) {
					atomic.AddInt64(&totalIPs, 1) // 增加遍历过的IP地址数量
					// Query the country using geoio database
					result, err := DB.Find(ip.String())
					if err != nil {
						// Error in querying, skip it
						continue
					}
					// Check if the country matches the regex
					if countryRegex.MatchString(result.String()) {
						atomic.AddInt64(&matchedIPs, 1) // 增加匹配成功的IP地址数量
						// Send the country and IP to the channel
						countryChan <- ip.String()
					}
					// 每遍历1000打印信息
					if totalIPs%1000 == 0 {
						fmt.Printf("\r已遍历IP地址数量: %d, 匹配成功的IP地址数量: %d", totalIPs, matchedIPs)
					}
				}
			}

		}()
	}

	workerWG.Wait()
	fmt.Printf("\r已遍历IP地址数量: %d, 匹配成功的IP地址数量: %d", totalIPs, matchedIPs)
}
