package api

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
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

func ParsePGWConfigXML() (SgiIpAddress string, SgiInterfaceName string) {
	// Open our xmlFile
	xmlFile, err := os.Open("/opt/rdl-flexcore/configs/EPCConfig_PGW.xml")
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
