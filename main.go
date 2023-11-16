// main.go
package main

import (
	"flag"
	"fmt"
	"regexp"
	"runtime"
	"sync"

	"github.com/zu1k/nali/pkg/qqwry"
)

const (
	db_path = "./qqwry.dat"
)

var (
	DB  *qqwry.QQwry
	err error

	version string = "开发版本" // 默认值，使用 -ldflags "-X main.version=x.y.z" 设置
)

func main() {
	// Define the parameters
	cidr := flag.String("cidr", "", "The CIDR or the path of the CIDR file")
	country := flag.String("country", "", "The regular expression of the country name")
	output := flag.String("output", "geo_ips.txt", "The name of the output file")
	concurrency := flag.Int("concurrency", runtime.NumCPU(), "The number of concurrent workers to use")

	// Parse the parameters and check their validity
	flag.Parse()
	if *cidr == "" {
		checkError(fmt.Errorf(""), "Usage: go run . -cidr <CIDR or CIDR file> [-country <country regex>] [-output <output file>]")
	}
	if *country == "" {
		*country = ".*"
	}

	// 打印版本信息和用户参数
	fmt.Printf("程序版本: %s\n", version)
	fmt.Printf("用户参数: CIDR=%s, Country=%s, Output=%s\n", *cidr, *country, *output)

	countryRegex := regexp.MustCompile(*country)

	// dowload geoip database if not find
	DB, err = qqwry.NewQQwry(db_path)
	checkError(err, "cannot download qqwry database,please check your internet.")

	// Create the channels
	ipChan := make(chan string, 128)    // For passing IP addresses
	matchIPChan := make(chan string, 4) // For passing matched countries and IP addresses
	var wg sync.WaitGroup               // For waiting all goroutines to finish
	wg.Add(3)                           // There are three goroutines
	defer wg.Wait()                     // Wait for them at the end

	// Start the goroutines
	go parseCIDR(*cidr, ipChan, &wg)                                      // Parse the CIDR and send IP addresses to ipChan
	go queryCountry(ipChan, matchIPChan, countryRegex, &wg, *concurrency) // Query the country and send matched countries and IP addresses to countryChan
	go writeOutput(matchIPChan, *output, &wg)                             // Write the output content to outputChan and the output file

	// 所有goroutine完成后的打印语句
	wg.Wait() // 确保所有goroutine已完成
	fmt.Printf("\n程序执行完毕！所有匹配的IP地址已保存到: %s\n", *output)
	fmt.Println("您可以打开该文件查看结果。")
}
