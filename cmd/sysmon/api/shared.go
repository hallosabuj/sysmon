package api

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sysmon/proto/sysmonpb"
	"time"
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

// This function will check whether nodes are running of not
func CheckForNodes(nodeType string) bool {
	status, err := exec.Command("sh", "-c", "/usr/bin/flexcorecli -c \""+nodeType+" get status\"").CombinedOutput()
	if err != nil {
		// Node is not running
		return false
	} else {
		if strings.Contains(string(status), "Active") {
			// Node is running
			return true
		} else {
			// Node is not running
			return false
		}
	}
}

func SetRightNbits(n int) uint32 {
	var result uint32 = 0
	for i := 0; i < n; i++ {
		result = result + uint32(math.Pow(2, float64(i)))
	}
	return result
}

// Start DHCP snooping and add other routes inside this function
func StartRoutigAgent() {
	for !CheckForNodes("pgw") {
		time.Sleep(1000000000 * 1)
	}
	///////////////////////////////////////
	dhcpServer := GetDHCPv4ServerDetails()
	if dhcpServer.dhcp_ipv4_address == "127.0.0.1" {
		// Internal DHCP server
		fmt.Println("Code here for Internal DHCP server")

		// Calculating the mask value
		num1, _ := Ip2long(dhcpServer.dhcp_ipv4_address_range_start)
		num2, _ := Ip2long(dhcpServer.dhcp_ipv4_address_range_end)
		rangeOfIP := num1 ^ num2
		i := 0
		for rangeOfIP > 0 {
			i++
			rangeOfIP = rangeOfIP >> 1
		}
		// Calculating network address
		networkIP := num1 & (^SetRightNbits(i))
		finalDestination := Long2ip(networkIP) + "/" + strconv.Itoa(32-i)
		sgiIP, sgiInterface := ParsePGWConfigXML()
		fmt.Println("FinalDestination :", finalDestination)
		fmt.Println("Server ID :", dhcpServer.dhcp_ipv4_address)
		fmt.Println("SgiIpAddress :", sgiIP)
		fmt.Println("SgiInterfaceName :", sgiInterface)
		// request := *&sysmonpb.IPRequest{Request: &sysmonpb.Request{SourceIp: finalDestination, Destination: "default", Intermediate: sgiIP, InterfaceName: sgiInterface}}
		// AddTable(&request)

	} else {
		fmt.Println("Code here for External DHCP server")
		// External DHCP server
		channelForPacket := make(chan IpWithMask)
		// Here we are fetching SGi interface name are running DHCP snooping that interface
		_, SGiInterface := ParsePGWConfigXML()
		go DhcpSnooping(channelForPacket, SGiInterface)
		go Worker(channelForPacket)
		// This portion need to be uncommented for GTP snooping working
		// go api.GtpSnooping(channelForPacket)
	}
	///////////////////////////////////////

	///////////////////////////////////////////////////////////////////////////////
	// Now we need to add rules and routes for other interfaces also
	interfaces := Interfaces()
	for i, _ := range interfaces {
		request := sysmonpb.Request{InterfaceName: interfaces[i].Name}
		//Do the same for interfaces except lo and virbr
		if !(strings.Contains(interfaces[i].Name, "lo") || strings.Contains(interfaces[i].Name, "virbr")) {
			interfaceDetails := InterfaceDetailsByName(&request)
			for j, _ := range interfaceDetails.NormalAddress {
				if strings.Compare(interfaceDetails.NormalAddress[j].Type, "V4") == 0 {
					ipWithMask := interfaceDetails.NormalAddress[j].IP
					// request := *&sysmonpb.IPRequest{Request: &sysmonpb.Request{SourceIp: ipWithMask, Destination: "default", Intermediate: strings.Split(ipWithMask, "/")[0], InterfaceName: interfaces[i].Name}}
					// AddTable(&request)
					fmt.Println(ipWithMask, "default", strings.Split(ipWithMask, "/")[0], interfaces[i].Name)
					break
				}
			}
		}
	}
}
