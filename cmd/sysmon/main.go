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
		signals = make(chan os.Signal, 1)
		done    = make(chan bool, 1)
	)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// go api.Test()
	go api.DhcpSnooping()
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
