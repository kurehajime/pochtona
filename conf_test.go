// conf_test.go
package main

import (
	"bytes"
	"testing"
)

func ParseConf_test(t *testing.T) {
	out := new(bytes.Buffer)
	GetSampleConf(out)
	err, _ := ParseConf(string(out.Bytes()))
	if err != nil {
		t.Fatal("failed to parse conf: %s", err)
	}
}
