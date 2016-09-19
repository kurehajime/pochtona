// conf.go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

type Conf struct {
	AllowIpAddresses []string `json:"allow_ip_addresses"`
	Port             int      `json:"port"`
	Actions          []Action `json:"actions"`
}
type Action struct {
	Id          string `json:"id"`
	Path        string `json:"path"`
	Description string `json:"description"`
}

//InitConf makes config folder,sample config file,sample script.
func InitConf(path string) (err error) {
	//make directory
	if err := os.MkdirAll(path, 0777); err != nil {
		return fmt.Errorf("make directory: %v", err)
	}

	//make json
	makeJson := func() error {
		Confjson, err := os.Create(filepath.Join(path, "pochtona.json"))
		if err != nil {
			return fmt.Errorf("make conf: %v", err)
		}
		defer Confjson.Close()
		if err = GetSampleConf(Confjson, path); err != nil {
			return fmt.Errorf("write conf: %v", err)
		}
		return nil
	}
	if err = makeJson(); err != nil {
		return err
	}

	//make sctript
	hello := ""
	if runtime.GOOS != "windows" {
		hello = "hello.sh"
	} else {
		hello = "hello.bat"
	}
	makeHello := func() error {
		HelloSh, err := os.Create(filepath.Join(path, hello))
		if err != nil {
			return fmt.Errorf("make hello: %v", err)
		}
		defer HelloSh.Close()
		_, err = HelloSh.WriteString("echo hello world!")
		if err != nil {
			return fmt.Errorf("write hello: %v", err)
		}
		return nil
	}
	if err = makeHello(); err != nil {
		return err
	}
	if err = os.Chmod(filepath.Join(path, hello), 0777); err != nil {
		return fmt.Errorf("chmod hello: %v", err)
	}
	return nil
}

//GetSampleConf writes sample config
func GetSampleConf(out io.Writer, path string) error {
	sampleConf := Conf{
		AllowIpAddresses: []string{"*.*.*.*"},
		Port:             8080,
		Actions: []Action{
			Action{
				Id:          "hello",
				Path:        filepath.Join(path, "hello.sh"),
				Description: "echo hello",
			},
		}}
	bytes, err := json.MarshalIndent(sampleConf, "", "    ")
	if err != nil {
		return fmt.Errorf("decode conf: %v", err)
	}
	fmt.Fprintln(out, string(bytes))
	return nil
}

//ReadConf read config
func ReadConf(path string) (Conf, error) {
	var c Conf
	fInfo, err := os.Stat(path)
	if err != nil {
		return c, fmt.Errorf("read conf: %v", err)
	}
	if fInfo.IsDir() == true {
		path += "/pochtona.json"
	}
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return c, fmt.Errorf("read conf: %v", err)
	}
	c, err = ParseConf(string(b))
	if err != nil {
		return c, fmt.Errorf("read conf: %v", err)
	}
	return c, nil
}

//ParseConf parse config
func ParseConf(str string) (Conf, error) {
	var c Conf
	if err := json.Unmarshal([]byte(str), &c); err != nil {
		return c, fmt.Errorf("parse conf: %v", err)
	}
	return c, nil
}
