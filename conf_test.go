// conf_test.go
package main

import (
	"bytes"
	"path/filepath"
	"testing"
)

func TestParseConf(t *testing.T) {
	current, err := filepath.Abs("./_test")
	if err != nil {
		t.Fatal("failed to parse conf: %s", err)
	}
	out := new(bytes.Buffer)
	GetSampleConf(out, current)
	_, err = ParseConf(string(out.Bytes()))
	if err != nil {
		t.Fatal("failed to parse conf: %s", err)
	}

}
func TestInitConf(t *testing.T) {
	current, err := filepath.Abs("./_test")
	if err != nil {
		t.Fatal("failed to init: %s", err)
	}
	err = InitConf(current)
	if err != nil {
		t.Fatal("failed to init: %s", err)
	}
}

func TestReadConf(t *testing.T) {
	current, err := filepath.Abs("./_test")
	if err != nil {
		t.Fatal("failed to read conf: %s", err)
	}
	c1, err := ReadConf(filepath.Join(current, "pochtona.json"))
	if err != nil {
		t.Fatal("failed to read conf: %s", err)
	}

	out := new(bytes.Buffer)
	GetSampleConf(out, current)
	c2, err := ParseConf(string(out.Bytes()))

	if c1.Actions[0].Title != c2.Actions[0].Title {
		t.Fatal("failed to read conf: %s!=%s2", c1.Actions[0].Title, c2.Actions[0].Title)
	}

}
