package api

import (
	"fmt"
	"os/exec"
	"strings"
	"sysmon/proto/sysmonpb"
)

type Rules struct {
	Priority string `json:"priority"`
	Rule     string `json:"rule"`
}

func AddIPRule(request *sysmonpb.IPRequest) string {
	var source_ip = request.Request.SourceIp
	var table_name = request.Request.TableName
	var response string
	var flag bool

	//ip rule add from source_ip lookup table_name
	//Check Rule exist or not
	//If not Exist	: Add the provided rule
	//Else			: Return

	rules, _ := exec.Command("ip", "rule", "list").Output()
	for _, line := range strings.Split(strings.TrimSuffix(string(rules), "\n"), "\n") {
		if strings.Contains(line, source_ip) && strings.Contains(line, table_name) {
			response = "Rule exist.."
			flag = true
			break
		}
	}
	if !flag {

		cmd := exec.Command("sudo", "ip", "rule", "add", "from", source_ip, "lookup", table_name)
		err := cmd.Run()

		if err != nil {
			fmt.Println(err)
			response = "Error"
		} else {
			response = "Rule added"
		}
	}
	return response
}

func DelIPRule(request *sysmonpb.IPRequest) string {
	var source_ip = request.Request.SourceIp
	var response string
	flag := false

	//ip rule del from source_ip
	//Check Rule exist or not
	//If Exist	: Delete the provided rule
	//Else		: Return

	rules, _ := exec.Command("ip", "rule", "list").Output()
	for _, line := range strings.Split(strings.TrimSuffix(string(rules), "\n"), "\n") {
		if strings.Contains(line, source_ip) {
			flag = true
			break
		}
	}
	if flag {
		cmd := exec.Command("sudo", "ip", "rule", "del", "from", source_ip)
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
			response = "Error"
		} else {
			response = "Rule deleted"
		}
	} else {
		response = "Rule not exist.."
	}
	return response
}

func IPRules() []Rules {
	rules, _ := exec.Command("ip", "rule", "list").Output()
	var result []Rules
	for _, line := range strings.Split(strings.TrimSuffix(string(rules), "\n"), "\n") {
		result = append(result, Rules{Priority: strings.Split(line, ":\t")[0], Rule: strings.Split(line, ":\t")[1]})
	}
	return result
}
