package api

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	// _ "github.com/mattn/go-sqlite3"
)

type PN_TPGWC_CONFIG struct {
	XMLName xml.Name `xml:"PN_TPGWC_CONFIG"`
	Node    Node     `xml:"node"`
}

type Node struct {
	XMLName   xml.Name `xml:"node"`
	IpaddrSgi string   `xml:"ipaddrSgi"`
	IntfSgi   string   `xml:"intfSgi"`
}

////////////////////////////////////////////////////////////////////////////////
// For this sqlite3 to be working we need some package
// /lib64/libc.so.6: version `GLIBC_2.28' not found (required by ./routing.bin)
////////////////////////////////////////////////////////////////////////////////

// func GetDHCPv4ServerAddress() string {
// 	// File for local testing
// 	// database, _ := sql.Open("sqlite3", "/home/sabuj/spicasys/Redline/Routing_issue/test/epcpgwconfigdb")
// 	// File for vm testing
// 	database, _ := sql.Open("sqlite3", "/opt/rdl-flexcore/db/epcpgwconfigdb")
// 	rows, _ := database.Query("select pdn_index,dhcp_ipv4_address from pgw_pdn_profile_table")
// 	var dhcp_ipv4_address string
// 	var pdn_index int
// 	for rows.Next() {
// 		rows.Scan(&pdn_index, &dhcp_ipv4_address)
// 		if pdn_index == 1 {
// 			break
// 		}
// 	}
// 	return dhcp_ipv4_address
// }

func ParsePGWConfigXML() (SgiIpAddress string, SgiInterfaceName string) {
	// Open our xmlFile
	// File for vm testing
	// xmlFile, err := os.Open("/opt/rdl-flexcore/configs/EPCConfig_PGW.xml")
	// File for local testing
	xmlFile, err := os.Open("/home/sabuj/spicasys/Redline/Routing_issue/EPCConfig_PGW.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println("Successfully Opened users.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()
	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

	// we initialize config variable
	var config PN_TPGWC_CONFIG
	// we unmarshal our byteArray which contains our xmlFiles content into 'config' which we defined above
	xml.Unmarshal(byteValue, &config)
	// fmt.Println(config)
	return config.Node.IpaddrSgi, config.Node.IntfSgi
}
