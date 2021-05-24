package api

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"sysmon/proto/sysmonpb"
)

type Response struct {
	Msg string `json:"msg"`
}

func MakeSudo() {
	user := os.Getenv("USER")
	password, _ := ioutil.ReadFile("/home/" + user + "/password")
	c1 := exec.Command("echo", string(password))
	c2 := exec.Command("sudo", "-S", "ls")

	var stderr bytes.Buffer
	read_end, write_end := io.Pipe()
	c1.Stdout = write_end
	c2.Stdin = read_end
	c2.Stderr = &stderr
	c1.Start()

	go func() {
		defer write_end.Close()
		c1.Wait()
	}()

	err := c2.Run()

	if err != nil {
		fmt.Println(stderr)
	}
}

// This function will check whether all nodes are running of not
func CheckForNodes() bool {
	// pass
	return true
}

// Start DHCP snooping and add other routes inside this function
func StartRoutigAgent() {
	channelForPacket := make(chan IpWithMask)
	// Here we are fetching SGi interface name are running DHCP snooping that interface
	_, SGiInterface := ParsePGWConfigXML()
	go DhcpSnooping(channelForPacket, SGiInterface)
	go Worker(channelForPacket)
	// This portion need to be uncommented for GTP snooping working
	// go api.GtpSnooping(channelForPacket)

	///////////////////////////////////////////////////////////////////////////////
	// Now we need to add rules and routes for other interfaces also
	interfaces := Interfaces()
	for i, _ := range interfaces {
		request := sysmonpb.Request{InterfaceName: interfaces[i].Name}
		interfaceDetails := InterfaceDetailsByName(&request)
		for j, _ := range interfaceDetails.NormalAddress {
			if strings.Compare(interfaceDetails.NormalAddress[j].Type, "V4") == 0 {
				ipWithMask := interfaceDetails.NormalAddress[j].IP
				// fmt.Println(ipWithMask, ":", interfaces[i].Name)
				request := *&sysmonpb.IPRequest{Request: &sysmonpb.Request{SourceIp: ipWithMask, Destination: "default", Intermediate: strings.Split(ipWithMask, "/")[0], InterfaceName: interfaces[i].Name}}
				AddTable(&request)
				fmt.Println(ipWithMask, "default", strings.Split(ipWithMask, "/")[0], interfaces[i].Name)
				break
			}
		}
	}
}
