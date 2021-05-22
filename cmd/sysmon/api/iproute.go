package api

import (
	"bytes"
	"os/exec"
	"strconv"
	"strings"
	"sysmon/proto/sysmonpb"
)

func AddIPRoute(request *sysmonpb.IPRequest) string {
	var response string
	var flag bool
	destination := request.Request.Destination
	intermediate := request.Request.Intermediate
	net_interface := request.Request.InterfaceName
	table_name := request.Request.TableName

	//sudo ip route add default via 192.168.57.1 dev enp0s3 table custom1
	//Check Route exist or not
	//If Not Exist	:  Add the provided route
	//Else  		:  Return

	routes, _ := exec.Command("ip", "route", "list", "table", table_name).Output()

	for _, line := range strings.Split(strings.TrimSuffix(string(routes), "\n"), "\n") {
		if strings.Contains(line, destination) && strings.Contains(line, intermediate) && strings.Contains(line, net_interface) && strings.Contains(line, table_name) {
			response = "Route exist.."
			flag = true
			break
		}
	}

	if !flag {
		cmd := exec.Command("ip", "route", "add", destination, "via", intermediate, "dev", net_interface, "table", table_name)
		var stderr bytes.Buffer
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err != nil {
			response = stderr.String()
		} else {
			response = "Route added"
		}
	}
	return response
}

func DelIPRoute(request *sysmonpb.IPRequest) string {
	var response string
	var flag bool
	destination := request.Request.Destination
	table_name := request.Request.TableName

	//sudo ip route del default table custom1
	//Check Route exist or not
	//If Not Exist	:  Do nothing
	//Else  		:  Delete the route

	routes, _ := exec.Command("ip", "route", "list", "table", table_name).Output()

	for _, line := range strings.Split(strings.TrimSuffix(string(routes), "\n"), "\n") {
		if strings.Contains(line, destination) {
			flag = true
			break
		}
	}

	if flag {
		cmd := exec.Command("ip", "route", "del", destination, "table", table_name)
		var stderr bytes.Buffer
		cmd.Stderr = &stderr
		err := cmd.Run()

		if err != nil {
			response = stderr.String()
		} else {
			response = "Route deleted"
		}
	} else {
		response = "Route does not exist"
	}
	return response
}

type Routes struct {
	Index     string
	Route     string
	TableName string
}

// This function provides all the routes present in the machine
func IPRoutes() []Routes {
	tables := Tables()
	var result []Routes
	index := 1
	for j := range tables {
		temp := IPRoutesByTableName(&sysmonpb.Request{TableName: strings.Fields(tables[j])[1]})
		for i := range temp {
			temp[i].Index = strconv.Itoa(index)
			result = append(result, temp[i])
			index += 1
		}
	}
	return result
}

// This function provides all the routes for a given routing table name
func IPRoutesByTableName(request *sysmonpb.Request) []Routes {
	rules, _ := exec.Command("ip", "route", "list", "table", request.TableName).Output()
	var result []Routes
	var i = 1
	for _, line := range strings.Split(strings.TrimSuffix(string(rules), "\n"), "\n") {
		if line != "" {
			result = append(result, Routes{Index: strconv.Itoa(i), Route: line, TableName: request.TableName})
			i = i + 1
		}
	}
	return result
}
