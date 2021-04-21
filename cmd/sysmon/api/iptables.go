package api

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"strings"
	"sysmon/proto/sysmonpb"
)

// Retruns array of table names
func Tables() []string {
	tables, err := ioutil.ReadFile("/etc/iproute2/rt_tables")
	if err != nil {
		fmt.Println("Error")
	}
	var result []string
	for _, line := range strings.Split(strings.TrimSuffix(string(tables), "\n"), "\n") {
		if !strings.Contains(line, "#") {
			result = append(result, line)
		}
	}
	return result
}

func Ip2long(ipAddr string) (uint32, error) {
	ip := net.ParseIP(ipAddr)
	if ip == nil {
		return 0, errors.New("wrong ipAddr format")
	}
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip), nil
}

func Long2ip(ipLong uint32) string {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, ipLong)
	ip := net.IP(ipByte)
	return ip.String()
}

func int32ToString(n uint32) string {
	buf := [11]byte{}
	pos := len(buf)
	i := int64(n)
	signed := i < 0
	if signed {
		i = -i
	}
	for {
		pos--
		buf[pos], i = '0'+byte(i%10), i/10
		if i == 0 {
			if signed {
				pos--
				buf[pos] = '-'
			}
			return string(buf[pos:])
		}
	}
}

func FindInArray(array []int, element int) bool {
	for i := range array {
		if array[i] == element {
			return true
		}
	}
	return false
}

func AddTable(request *sysmonpb.IPRequest) string {
	var interfaceName string = request.Request.InterfaceName
	var ip string = strings.Split(request.Request.SourceIp, "/")[0]
	//Create a number w.r.t. the ip
	ipNumber, err := Ip2long(ip)
	var response string
	if err != nil {
		response = err.Error()
		return response
	}
	ip = int32ToString(ipNumber)
	//Check whether table exist or not
	tables := Tables()
	var tableNumber []int
	var tableExists bool = false
	for i := range tables {
		temp, _ := strconv.Atoi(strings.Fields(tables[i])[0])
		tableNumber = append(tableNumber, temp)
		if strings.Contains(strings.Fields(tables[i])[1], interfaceName+"_"+ip) {
			tableExists = true
			break
		}
	}

	// If table exist then add the rule and route directly
	if tableExists {
		response = "Table exist"
		request.Request.TableName = interfaceName + "_" + ip
		response = response + "\n" + AddIPRule(request)
		response = response + "\n" + strings.TrimSuffix(AddIPRoute(request), "\n")
		return response
	} else {
		newTableNumber := 200
		newTableName := interfaceName + "_" + ip
		for true {
			if FindInArray(tableNumber, newTableNumber) {
				newTableNumber = newTableNumber + 1
			} else {
				break
			}
		}
		lineToAdd := strconv.Itoa(newTableNumber) + "\t" + newTableName + "\n"
		//Now write to rt_tables
		file, err := os.OpenFile("/etc/iproute2/rt_tables", os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			// fmt.Println(err)
			response = err.Error()
			return response
		}
		defer file.Close()
		if _, err := file.WriteString(lineToAdd); err != nil {
			// fmt.Println(err)
			response = err.Error()
			return response
		}
		response = "Table added..."
		request.Request.TableName = newTableName
		response = response + "\n" + AddIPRule(request)
		response = response + "\n" + strings.TrimSuffix(AddIPRoute(request), "\n")
		return response
	}
}
