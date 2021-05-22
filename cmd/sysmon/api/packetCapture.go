package api

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"sysmon/proto/sysmonpb"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/wmnsk/go-gtp/gtpv2/message"
)

type IpWithMask struct {
	IP         string
	SubnetMask string
}

var AllocatedIPs []IpWithMask

var IPs map[string]string = make(map[string]string)

func Worker(ch chan IpWithMask) {
	for {
		select {
		case data := <-ch:
			IPs[data.IP] = data.SubnetMask
		}
	}
}

func DhcpSnooping(ch chan IpWithMask, interfaceName string) {
	var (
		device  string = interfaceName
		snaplen int32  = 1000000000
		promisc bool   = false
		err     error
		timeout time.Duration = -1 * time.Second
		handle  *pcap.Handle
	)
	handle, err = pcap.OpenLive(device, snaplen, promisc, timeout)

	// ////////////////////////////////////////////////////////////////////////////////
	// // This portion need to be deleted and abobe portion needs to be uncommented
	// var (
	// 	err    error
	// 	handle *pcap.Handle
	// )
	// handle, err = pcap.OpenOffline("/home/sabuj/spicasys/sabuj/sysmon/bin/ens192.pcap")
	// ////////////////////////////////////////////////////////////////////////////////
	if err != nil {
		log.Println("Error here", err)
		return
	}
	defer handle.Close()

	var filter string = "udp port 67"
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Println(err)
		return
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSource.Packets() {
		// DHCPv4 packet have 4 layers, and LayerType in DHCPv4 and integer representaion is 118
		if len(packet.Layers()) == 4 && packet.Layers()[3].LayerType() == 118 {
			dhcp_layer := packet.Layer(layers.LayerTypeDHCPv4)
			dhcp_packet := dhcp_layer.(*layers.DHCPv4)

			var subnetMask string
			var destinationIP string
			var isAckMessage bool = false // This variable is made true if packet is an acknowledgement packet
			var dhcp_ipv4_address string = ""
			for i := range dhcp_packet.Options {
				if strings.Contains(dhcp_packet.Options[i].String(), "(MessageType:Ack)") {
					isAckMessage = true
					destinationIP = dhcp_packet.YourClientIP.String()
				}
				if strings.Contains(dhcp_packet.Options[i].String(), "(SubnetMask:") {
					subnetMask = strings.Replace(strings.Split(dhcp_packet.Options[i].String(), ":")[1], ")", "", 1)
				}
				if strings.Contains(dhcp_packet.Options[i].String(), "ServerID") {
					dhcp_ipv4_address = strings.Replace(strings.Split(dhcp_packet.Options[i].String(), ":")[1], ")", "", 1)
				}
			}
			// { SGi interface name : ens192, Gateway of SGi interface : 10.250.0.1, SGi interface address : 10.250.0.152 }
			// Fetching SGi interface name and SGi interface address from the XML file
			SgiIpAddress, SgiInterfaceName := ParsePGWConfigXML()

			// Here we are checking 3 conditions that are
			// 1. The packet is an acknowledgement packet
			// 2. SGi intrface ip is same as relay agent IP mentioned in DHCP packet
			//    Which signify that the packet is destined for this SGi only
			// 3. DHCP server address in DHCP packet is same as mentioned in the database
			//    Which signify that the DHCP packet coming fro the correct DHCP server
			if isAckMessage && SgiIpAddress == dhcp_packet.RelayAgentIP.String() && dhcp_ipv4_address == GetDHCPv4ServerAddress() {

				// Here destination assress is modified and Subnet Mask is added and finalDestination IP is generated
				dst := net.ParseIP(destinationIP)
				mask := net.IPMask(net.ParseIP(subnetMask).To4())
				prefixSize, _ := mask.Size()
				finalDestination := dst.Mask(mask).String() + "/" + strconv.Itoa(prefixSize) // Allocated UE range
				fmt.Println("FinalDestination :", finalDestination)
				fmt.Println("Interface Name :", interfaceName)
				fmt.Println("Server ID :", dhcp_ipv4_address)
				fmt.Println("SgiIpAddress :", SgiIpAddress)
				fmt.Println("SgiInterfaceName :", SgiInterfaceName)
				request := *&sysmonpb.IPRequest{Request: &sysmonpb.Request{SourceIp: finalDestination, Destination: "default", Intermediate: SgiIpAddress, InterfaceName: SgiInterfaceName}}

				// Now call the AddTable() of routing agent
				// Which will add the table, rule and the route in the VM.
				AddTable(&request)
				isAckMessage = false
				ch <- IpWithMask{IP: destinationIP, SubnetMask: subnetMask}
			}
		}
	}
}

func GtpSnooping(ch chan IpWithMask) {
	var (
		device  string = "lo"
		snaplen int32  = 1000000000
		promisc bool   = false
		err     error
		timeout time.Duration = -1 * time.Second
		handle  *pcap.Handle
	)
	handle, err = pcap.OpenLive(device, snaplen, promisc, timeout)

	//////////////////////////////////////////////////////////////////////////////////
	// var (
	// 	err    error
	// 	handle *pcap.Handle
	// )
	// handle, err = pcap.OpenOffline("/home/sabuj/spicasys/sabuj/sysmon/bin/gtp.pcap")
	//////////////////////////////////////////////////////////////////////////////////
	if err != nil {
		log.Println(err)
		return
	}
	defer handle.Close()

	var filter string = "udp port 2123"
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Println(err)
		return
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSource.Packets() {
		if len(packet.Layers()) > 2 && packet.Layers()[2].LayerType() == 45 { //Checking layer 2 is UDP or not
			msg, err := message.Parse(packet.Layers()[2].LayerPayload())
			if err != nil {
				//Unable to decode means it is not a GTPv2 packet
				continue
			} else {
				// fmt.Println(msg.MessageTypeName())
				if msg.MessageType() == 33 { //33 means "Create Session Response"
					resp, err := message.ParseCreateSessionResponse(packet.Layers()[2].LayerPayload())
					if err != nil {
						fmt.Println(err)
					} else {
						if resp.PAA != nil {
							// fmt.Println(resp.PAA)
							var ip string
							ip = strconv.Itoa(int(resp.PAA.Payload[1]))
							for j := 2; j < len(resp.PAA.Payload); j++ {
								ip += "." + strconv.Itoa(int(resp.PAA.Payload[j]))
							}
							// fmt.Println("UE IP :", ip)
							ch <- IpWithMask{IP: ip}
							// AllocatedIPs = append(AllocatedIPs, IpWithMask{IP: ip})
							// fmt.Println("Allocated IP:", AllocatedIPs)
						}
					}
				}
			}
		}
	}
}

func GetAllocatedIP() []IpWithMask {
	AllocatedIPs = AllocatedIPs[:0]
	for key, value := range IPs {
		AllocatedIPs = append(AllocatedIPs, IpWithMask{IP: key, SubnetMask: value})
	}
	return AllocatedIPs
}
