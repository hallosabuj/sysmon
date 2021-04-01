package api

import (
	"encoding/json"
	"net"
	"net/http"
	"strings"
)

type Gateway struct {
	Destination string `json:"destination"`
	Gateway     string `json:"gateway"`
}

type IpAddr struct {
	IP   string `json:"ip"`
	Type string `json:"type"`
}

type NetInterface struct {
	Index int16  `json:"index"`
	Name  string `json:"name"`
}

type NetInterfaceDetails struct {
	Name             string    `json:"name"`
	Gateway          []Gateway `json:"gatewayList"`
	NormalAddress    []IpAddr  `json:"normalAddress"`
	MulticastAddress []IpAddr  `json:"multicastAddress"`
}

func ListInterfaceAddresses(w http.ResponseWriter, r *http.Request) {
	addr, _ := net.InterfaceAddrs()
	var result []IpAddr

	for i := range addr {
		temp := IpAddr{IP: addr[i].String(), Type: CheckV4orV6(addr[i].String())}
		result = append(result, temp)
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(result)
}

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

// func GetInterfaceInfoByName(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	interface_name := vars["name"]
// 	interfaceObj, _ := net.InterfaceByName(interface_name)

// 	var result NetInterfaceDetails
// 	result.Name = interface_name
// 	ipAddr, _ := interfaceObj.Addrs()

// 	//fmt.Println("Normal addresses...")
// 	var tempAddress []IpAddr
// 	for i := range ipAddr {
// 		tempAddress = append(tempAddress, IpAddr{IP: ipAddr[i].String(), Type: CheckV4orV6(ipAddr[i].String())})
// 	}

// 	result.NormalAddress = tempAddress

// 	multicastAddr, _ := interfaceObj.MulticastAddrs()
// 	//fmt.Println("Multicast addresses...")
// 	tempAddress = nil
// 	for i := range multicastAddr {
// 		tempAddress = append(tempAddress, IpAddr{IP: multicastAddr[i].String(), Type: CheckV4orV6(multicastAddr[i].String())})
// 	}
// 	result.MulticastAddress = tempAddress

// 	//Fetching gateway details
// 	cmd := exec.Command("sh", "-c", "netstat -rn | grep "+interface_name+" | tr -s ' ' | cut -d ' ' -f 1-2")
// 	output, err := cmd.CombinedOutput()
// 	if err != nil {
// 		fmt.Println("Error : ", err)
// 	}

// 	var gatewayList []Gateway
// 	for _, line := range strings.Split(strings.TrimSuffix(string(output), "\n"), "\n") {
// 		temp := Gateway{Destination: strings.Split(line, " ")[0], Gateway: strings.Split(line, " ")[1]}
// 		gatewayList = append(gatewayList, temp)
// 	}
// 	result.Gateway = gatewayList

// 	w.Header().Set("content-type", "application/json")
// 	json.NewEncoder(w).Encode(result)
// }

func ListInterfaces(w http.ResponseWriter, r *http.Request) {
	interfaces, _ := net.Interfaces()
	var result []NetInterface

	for i := range interfaces {
		temp := NetInterface{Index: int16(interfaces[i].Index), Name: interfaces[i].Name}
		result = append(result, temp)
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(result)
}
