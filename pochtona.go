// pochtona.go
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mitchellh/cli"
)

func main() {
	c := cli.NewCLI("pochtona", "1.0.0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"init": func() (cli.Command, error) {
			return &Init{}, nil
		},
		"run": func() (cli.Command, error) {
			return &Run{}, nil
		},
	}
	exitStatus, err := c.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	os.Exit(exitStatus)
}

type Run struct{}

func (t *Run) Help() string {
	return "Run Server \n\t $pochtona run [Config file]"
}

func (t *Run) Run(args []string) int {
	var path string
	var err error
	if len(args) < 1 {
		_, err = ioutil.ReadFile("./pochtona.json")
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			fmt.Fprintln(os.Stderr, "Config file not found.")
			fmt.Fprintln(os.Stderr, "Please run `pochtona init`")
			return 1
		}
		path = "./pochtona.json"
	} else {
		path = args[0]
	}
	c, err := ReadConf(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return 1
	}
	Start(c)
	return 0
}

func (t *Run) Synopsis() string {
	return "Run server."
}

type Init struct{}

func (t *Init) Help() string {
	return "Make sample config file and Sample script. \n\t $pochtona init [path]"
}

func (t *Init) Run(args []string) int {
	var path string
	if len(args) < 1 {
		path = "."
	} else {
		path = args[0]
	}
	current, err := filepath.Abs(path)
	if err != nil {
		fmt.Println(err)
		return 1
	}
	err = InitConf(current)
	if err != nil {
		fmt.Println(err)
		return 1
	}
	return 0
}

func (t *Init) Synopsis() string {
	return "Make sample config file and Sample script."
}
