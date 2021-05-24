package api

import (
	"fmt"
	"net"
	"os/exec"
	"strings"
	"sysmon/proto/sysmonpb"
)

type Gateway struct {
	Destination string `json:"destination"`
	Gateway     string `json:"gateway"`
}

type IpAddr struct {
	IP   string `json:"ip"`
	Type string `json:"type"`
}

type Interface struct {
	Index int    `json:"index"`
	Name  string `json:"name"`
}

type NetInterfaceDetails struct {
	Name             string    `json:"name"`
	Gateway          []Gateway `json:"gatewayList"`
	NormalAddress    []IpAddr  `json:"normalAddress"`
	MulticastAddress []IpAddr  `json:"multicastAddress"`
}

// Returns all the allocated IP in the machine
func InterfaceAddresses() []IpAddr {
	addr, _ := net.InterfaceAddrs()
	var result []IpAddr
	for i := range addr {
		temp := IpAddr{IP: addr[i].String(), Type: CheckV4orV6(addr[i].String())}
		result = append(result, temp)
	}
	return result
}

// Check an IP whether it IPv4 or IPv6
// For invalid IP returns nil
func CheckV4orV6(ip string) string {
	ip = strings.Split(ip, "/")[0]
	parsedIP := net.ParseIP(ip)
	isV4 := net.IP.To4(parsedIP)
	if isV4 != nil {
		return "V4"
	}
	isV6 := net.IP.To16(parsedIP)
	if isV6 != nil {
		return "V6"
	}
	return "nil"
}

// Returns interface details for a given interface name
func InterfaceDetailsByName(request *sysmonpb.Request) NetInterfaceDetails {
	interface_name := request.InterfaceName
	interfaceObj, _ := net.InterfaceByName(interface_name)

	var result NetInterfaceDetails
	result.Name = interface_name
	ipAddr, _ := interfaceObj.Addrs()

	//fmt.Println("Normal addresses...")
	var tempAddress []IpAddr
	for i := range ipAddr {
		tempAddress = append(tempAddress, IpAddr{IP: ipAddr[i].String(), Type: CheckV4orV6(ipAddr[i].String())})
	}

	result.NormalAddress = tempAddress

	multicastAddr, _ := interfaceObj.MulticastAddrs()
	//fmt.Println("Multicast addresses...")
	tempAddress = nil
	for i := range multicastAddr {
		tempAddress = append(tempAddress, IpAddr{IP: multicastAddr[i].String(), Type: CheckV4orV6(multicastAddr[i].String())})
	}
	result.MulticastAddress = tempAddress

	//Fetching gateway details
	cmd := exec.Command("sh", "-c", "netstat -rn | grep "+interface_name+" | tr -s ' ' | cut -d ' ' -f 1-2")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error : ", err)
	}

	var gatewayList []Gateway
	if len(output) > 0 {
		for _, line := range strings.Split(strings.TrimSuffix(string(output), "\n"), "\n") {
			temp := Gateway{Destination: strings.Split(line, " ")[0], Gateway: strings.Split(line, " ")[1]}
			gatewayList = append(gatewayList, temp)
		}
	}
	result.Gateway = gatewayList
	return result
}

// Returns name of all the available interfaces in the machine
func Interfaces() []Interface {
	interfaces, _ := net.Interfaces()
	var result []Interface
	for i := range interfaces {
		temp := Interface{Index: interfaces[i].Index, Name: interfaces[i].Name}
		result = append(result, temp)
	}
	return result
}
