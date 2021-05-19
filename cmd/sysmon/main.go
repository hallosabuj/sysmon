package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"sysmon/cmd/sysmon/api"
	"sysmon/cmd/sysmon/server"
)

func main() {
	var (
		port = flag.Int("port", 8080, "http server port")
	)
	flag.Parse()

	var (
		signals          = make(chan os.Signal, 1)
		done             = make(chan bool, 1)
		channelForPacket = make(chan api.IpWithMask)
	)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// Here we are fetching all the interface name are running DHCP snooping on each interface
	interfaces := api.Interfaces()
	for i := 0; i < len(interfaces); i++ {
		go api.DhcpSnooping(channelForPacket, interfaces[i].Name)
		go api.Worker(channelForPacket)
	}
	// This portion need to be uncommented for GTP snooping working
	// go api.GtpSnooping(channelForPacket)
	go func() {
		<-signals
		done <- true
	}()

	go func() {
		err := server.Start(*port)
		if nil != err {
			fmt.Println("Failed starting server, error: ", err)
			done <- true
		}
	}()

	<-done
}
