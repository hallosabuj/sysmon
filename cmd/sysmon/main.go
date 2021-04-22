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
	go api.DhcpSnooping(channelForPacket)
	go api.GtpSnooping(channelForPacket)
	go api.Worker(channelForPacket)
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
