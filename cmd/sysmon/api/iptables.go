package api

import (
	"fmt"
	"io/ioutil"
	"strings"
)

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
