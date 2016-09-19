// server_test.go
package main

import (
	"html/template"
	"strings"
	"testing"
)

func TestIndex(t *testing.T) {
	temp := template.New("index")
	temp, err := temp.Parse(templete_index)
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
