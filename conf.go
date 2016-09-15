// conf.go
package main

import (
	"encoding/json"
	"fmt"
	"io"
)

type Conf struct {
	Actions []Action `json:"actions"`
}
type Action struct {
	Title       string `json:"title"`
	Path        string `json:"path"`
	Description string `json:"description"`
}

func GetSampleConf(out io.Writer) error {
	sampleConf := Conf{Actions: []Action{
		Action{
			Title:       "hello",
			Path:        "./sample/hello.sh",
			Description: "echo hello",
		},
	}}
	bytes, err := json.MarshalIndent(sampleConf, "", "    ")
	if err != nil {
		return err
	}
	fmt.Fprintln(out, string(bytes))
	return nil
}
func ParseConf(str string) (error, Conf) {
	var c Conf
	if err := json.Unmarshal([]byte(str), &c); err != nil {
		return err, c
	}
	return nil, c
}
