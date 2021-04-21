package api

import (
	"fmt"
	"log"
	"strconv"
	"strings"

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

var IPs map[string]string

func DhcpSnooping() {
	IPs = make(map[string]string)
	// var (
	// 	device  string = "enp0s8"
	// 	snaplen int32  = 65535
	// 	promisc bool   = false
	// 	err     error
	// 	timeout time.Duration = -1 * time.Second
	// 	handle  *pcap.Handle
	// )
	// handle, err = pcap.OpenLive(device, snaplen, promisc, timeout)
	///////////////////////////////////////////////////////////////
	var (
		err    error
		handle *pcap.Handle
	)
	handle, err = pcap.OpenOffline("/home/sabuj/spicasys/sabuj/sysmon/bin/temp.pcap")
	///////////////////////////////////////////////////////////////
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// var filter string = "src host 172.20.1.8 and icmp"
	// err = handle.SetBPFFilter(filter)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSource.Packets() {
		if len(packet.Layers()) == 4 && packet.Layers()[3].LayerType() == 118 {
			// fmt.Println(packet.String())
			fmt.Println("Layer related ::::::::::::::::::::::")
			dhcp_layer := packet.Layer(layers.LayerTypeDHCPv4)
			dhcp_packet := dhcp_layer.(*layers.DHCPv4)

			var subnetMask string
			var destinationIP string
			var srcIP string
			var isAckMessage bool = false
			for i := range dhcp_packet.Options {
				if strings.Contains(dhcp_packet.Options[i].String(), "(MessageType:Ack)") {
					fmt.Println("Acknowledgement")
					isAckMessage = true
					ip_layer := packet.Layer(layers.LayerTypeIPv4)
					ip_packet := ip_layer.(*layers.IPv4)
					destinationIP = ip_packet.DstIP.String()
					srcIP = ip_packet.SrcIP.String()
				}
				if strings.Contains(dhcp_packet.Options[i].String(), "(SubnetMask:") {
					subnetMask = strings.Replace(strings.Split(dhcp_packet.Options[i].String(), ":")[1], ")", "", 1)
				}
			}
			if isAckMessage {
				// allocatedIPs = append(allocatedIPs, IpWithMask{Index: "1", IP: destinationIP, SubnetMask: subnetMask})
				IPs[destinationIP] = subnetMask
				fmt.Println("Destination address (IP allocated) :", destinationIP)
				fmt.Println("Subnet Mask :", subnetMask)
				fmt.Println("Source IP :", srcIP)
			}
			// fmt.Println("DHCP code :", dhcp_packet.Options[0])
			fmt.Println("Layer related ::::::::::::::::::::::")
			AllocatedIPs = AllocatedIPs[:0]
			for key, value := range IPs {
				AllocatedIPs = append(AllocatedIPs, IpWithMask{IP: key, SubnetMask: value})
			}
			fmt.Println(AllocatedIPs)
		}
	}
}

func gtpParsing() {
	var (
		err    error
		handle *pcap.Handle
	)
	handle, err = pcap.OpenOffline("/root/EPC_Capture_15-Apr-2021_20-47-54.pcap")
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

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
						// fmt.Println(resp.PAA)
						var ip string
						ip = strconv.Itoa(int(resp.PAA.Payload[1]))
						for j := 2; j < len(resp.PAA.Payload); j++ {
							ip += "." + strconv.Itoa(int(resp.PAA.Payload[j]))
						}
						fmt.Println("UE IP :", ip)
					}
				}
			}
		}
	}
}

// func Test() {
// 	AllocatedIPs = append(AllocatedIPs, IpWithMask{IP: "120.120.120.120", SubnetMask: "255.255.255.0"})
// 	AllocatedIPs = append(AllocatedIPs, IpWithMask{IP: "220.220.220.220", SubnetMask: "255.255.255.0"})
// }

func GetAllocatedIP() []IpWithMask {
	return AllocatedIPs
}
