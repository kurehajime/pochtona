// server_test.go
package main

import (
	"html/template"
	"strings"
	"testing"
)

func TestIndex(t *testing.T) {
	temp := template.New("index")
	temp, err := temp.Parse(assets["_assets/index.html"])
	if err != nil {
		t.Error("template parse error:", err)
	}
}

func TestCmd(t *testing.T) {
	//Command OK?
	ans := "HelloWorld"
	res := cmd("echo HelloWorld")
	if strings.Trim(ans, " \r\n") != strings.Trim(res, " \r\n") {
		t.Error("cmd Error:", ans, res)
		return
	}
}
func TestCheckIP(t *testing.T) {
	tests := []struct {
		addr             string
		allowIpAddresses []string
		assert           bool
	}{
		{"192.168.1.1", []string{"192.168.1.1"}, true},
		{"192.168.1.2", []string{"192.168.1.1"}, false},
		{"192.168.1.42", []string{"192.168.1.*"}, true},
		{"192.168.42.1", []string{"192.168.1.*"}, false},
		{"192.168.42.1", []string{"192.168.*.1"}, true},
	}
	for i := range tests {
		if checkIP(tests[i].addr, tests[i].allowIpAddresses) != tests[i].assert {
			t.Fatal("failed to IP Check: %s ,%s assert :%s", tests[i].addr, tests[i].allowIpAddresses, tests[i].assert)
		}
	}

}
