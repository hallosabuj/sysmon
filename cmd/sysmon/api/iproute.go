package api

import (
	"fmt"
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
		if strings.Contains(line, destination) {
			response = "Route exist.."
			flag = true
			break
		}
	}

	if !flag {
		cmd := exec.Command("sudo", "ip", "route", "add", destination, "via", intermediate, "dev", net_interface, "table", table_name)
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error here : ", err)
			response = "Error"
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
		cmd := exec.Command("sudo", "ip", "route", "del", destination, "table", table_name)
		err := cmd.Run()

		if err != nil {
			fmt.Println("Error here : ", err)
			response = "Error"
		} else {
			response = "Route deleted"
		}
	} else {
		response = "Route does not exist"
	}
	return response
}

func IPRoutes() map[string]string {
	rules, _ := exec.Command("ip", "route", "list").Output()
	result := make(map[string]string)
	var i = 1
	for _, line := range strings.Split(strings.TrimSuffix(string(rules), "\n"), "\n") {
		result[strconv.Itoa(i)] = line
		i = i + 1
	}
	return result
}
