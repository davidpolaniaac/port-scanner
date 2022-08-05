package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

var host = flag.String("host", "scanme.nmap.org", "host to scan")

var startPort = flag.Int("start", 22, "start of port range")

var endPort = flag.Int("end", 80, "end of port range")

type Scan struct{}

func (Scan) scanner(host string, startPort int, endPort int) {
	var wg sync.WaitGroup
	for i := startPort; i <= endPort; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("Port %d is open\n", port)
		}(i)
	}
	wg.Wait()
}

func main() {
	flag.Parse()
	var scan Scan
	startTime := time.Now()
	fmt.Println("scanning ports.....")
	scan.scanner(*host, *startPort, *endPort)
	fmt.Printf("total scan time: %s\n", time.Since(startTime))
}
