package api

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
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

// This portion of code will be used for checking PGW status

type UsernameToken struct {
	Username  string
	CreatedAt string
	Nonce     string
	Password  string
}

func makeTwoDigitRepresentation(num int) string {
	if num < 0 {
		num = -num
	}
	if num < 10 {
		return fmt.Sprintf("0%v", num)
	} else {
		return fmt.Sprintf("%v", num)
	}
}
func GenerateUsernameToken() (usernameToken UsernameToken) {
	createdAt := ""
	presentTime := time.Now()
	_, offset := presentTime.Zone()
	if (offset/60)/60 > 0 {
		createdAt = fmt.Sprintf("%v-%v-%vT%v:%v:%v+%v:%v", presentTime.Year(), int(presentTime.Month()), presentTime.Day(), presentTime.Hour(), presentTime.Minute(), presentTime.Second(), makeTwoDigitRepresentation((offset/60)/60), makeTwoDigitRepresentation((offset/60)%60))
	} else {
		createdAt = fmt.Sprintf("%v-%v-%vT%v:%v:%v-%v:%v", presentTime.Year(), int(presentTime.Month()), presentTime.Day(), presentTime.Hour(), presentTime.Minute(), presentTime.Second(), makeTwoDigitRepresentation((offset/60)/60), makeTwoDigitRepresentation((offset/60)%60))
	}
	simple_nonce := 620452692
	encoded_nonce := base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(simple_nonce)))
	h := sha1.New()
	inputForSHA1 := strconv.Itoa(simple_nonce) + createdAt + "rdl"
	h.Write([]byte(inputForSHA1))
	bs := h.Sum(nil)

	usernameToken.CreatedAt = createdAt
	usernameToken.Nonce = encoded_nonce
	usernameToken.Username = "Redline"
	usernameToken.Password = base64.StdEncoding.EncodeToString(bs)

	return usernameToken
}

func GetPgwPort() string {
	PGW_MANAGEMENT_PORT := ""
	// environmentVariables, err := exec.Command("cat", "env.txt").Output()
	environmentVariables, err := exec.Command("systemctl", "show-environment").Output()
	if err != nil {
		fmt.Println("Error occured :", err)
	}

	for _, line := range strings.Split(strings.TrimSuffix(string(environmentVariables), "\n"), "\n") {
		if strings.Contains(line, "PGW_MANAGEMENT_PORT") {
			PGW_MANAGEMENT_PORT = strings.Split(line, "=")[1]
		}
	}
	return PGW_MANAGEMENT_PORT
}

func soapCall(usernameToken UsernameToken, PGW_MANAGEMENT_PORT string) string {
	username := usernameToken.Username
	password := usernameToken.Password
	nonce := usernameToken.Nonce
	createdAt := usernameToken.CreatedAt
	xmlbody := `<?xml version="1.0" encoding="UTF-8"?>
				<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns1="http://tempuri.org/EPC-PGW.xsd" xmlns:ns2="http://tempuri.org/wsse.xsd" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:SOAP-ENC="http://schemas.xmlsoap.org/soap/encoding/" SOAP-ENV:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
				<SOAP-ENV:Header><wsse:Security xmlns:wsse="http://tempuri.org/wsse.xsd">
				<wsse:UsernameToken>
					<wsse:Username>` + username + `</wsse:Username>
					<wsse:Password Type="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordDigest">` + password + `</wsse:Password>
					<wsse:Nonce EncodingType="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-soap-message-security-1.0#Base64Binary">` + nonce + `</wsse:Nonce>
					<wsu:Created xmlns:wsu="http://tempuri.org/wsu.xsd">` + createdAt + `</wsu:Created>
				</wsse:UsernameToken>
				</wsse:Security>
				</SOAP-ENV:Header>
				<SOAP-ENV:Body>
					<ns1:GetNetEPCStatus/>
				</SOAP-ENV:Body>
				</SOAP-ENV:Envelope>`

	resp, err := http.Post("http://127.0.0.1:"+PGW_MANAGEMENT_PORT, "text/xml", strings.NewReader(xmlbody))
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body) // byte body
	return string(body)
}

// This function will check whether PWG are running of not
func GetPgwStatus() bool {
	usernameToken := GenerateUsernameToken()
	stringValueBody := soapCall(usernameToken, GetPgwPort())
	if len(strings.Split(stringValueBody, "<item>")) > 1 {
		status := strings.Split(strings.Split(strings.Split(stringValueBody, "<item>")[1], "</itemp>")[0], "|")[0]
		if status == "Active" {
			return true
		} else {
			return false
		}

	} else {
		return false
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
	for !GetPgwStatus() {
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

// // This function will check whether nodes are running of not
// func CheckForNodes(nodeType string) bool {
// 	status, err := exec.Command("sh", "-c", "/usr/bin/flexcorecli -c \""+nodeType+" get status\"").CombinedOutput()
// 	if err != nil {
// 		// Node is not running
// 		return false
// 	} else {
// 		if strings.Contains(string(status), "Active") {
// 			// Node is running
// 			return true
// 		} else {
// 			// Node is not running
// 			return false
// 		}
// 	}
// }
