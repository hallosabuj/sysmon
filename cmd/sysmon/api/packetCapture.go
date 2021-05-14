package api

import (
	"fmt"
	"log"
	"strconv"
	"strings"
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
	// handle, err = pcap.OpenOffline("/home/sabuj/spicasys/sabuj/sysmon/bin/EPC_BB.pcap")
	// ////////////////////////////////////////////////////////////////////////////////
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	var filter string = "udp port 67"
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal(err)
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSource.Packets() {
		if len(packet.Layers()) == 4 && packet.Layers()[3].LayerType() == 118 {
			dhcp_layer := packet.Layer(layers.LayerTypeDHCPv4)
			dhcp_packet := dhcp_layer.(*layers.DHCPv4)

			var subnetMask string
			var destinationIP string
			// var srcIP string
			var isAckMessage bool = false
			for i := range dhcp_packet.Options {
				if strings.Contains(dhcp_packet.Options[i].String(), "(MessageType:Ack)") {
					// fmt.Println("Acknowledgement")
					isAckMessage = true
					// ip_layer := packet.Layer(layers.LayerTypeIPv4)
					// ip_packet := ip_layer.(*layers.IPv4)
					// destinationIP = ip_packet.DstIP.String()
					destinationIP = dhcp_packet.YourClientIP.String()
					// srcIP = ip_packet.SrcIP.String()
				}
				if strings.Contains(dhcp_packet.Options[i].String(), "(SubnetMask:") {
					subnetMask = strings.Replace(strings.Split(dhcp_packet.Options[i].String(), ":")[1], ")", "", 1)
				}
			}
			if isAckMessage {
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
		log.Fatal(err)
	}
	defer handle.Close()

	var filter string = "udp port 2123"
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal(err)
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

// func Test() {
// 	AllocatedIPs = append(AllocatedIPs, IpWithMask{IP: "120.120.120.120", SubnetMask: "255.255.255.0"})
// 	AllocatedIPs = append(AllocatedIPs, IpWithMask{IP: "220.220.220.220", SubnetMask: "255.255.255.0"})
// }

func GetAllocatedIP() []IpWithMask {
	AllocatedIPs = AllocatedIPs[:0]
	for key, value := range IPs {
		AllocatedIPs = append(AllocatedIPs, IpWithMask{IP: key, SubnetMask: value})
	}
	return AllocatedIPs
}
