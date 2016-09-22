// server.go
package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

var conf Conf

//start http server
func Start(c Conf) {
	conf = c
	//start server
	url := "http://localhost" + ":" + strconv.Itoa(c.Port)
	fmt.Println("Stop: Ctrl+C")
	//Handle
	http.HandleFunc("/", index)
	fmt.Println("Top page:\n\t" + url)
	for _, a := range c.Actions {
		fmt.Println(a.Id + ":\n\t" + url + "/" + a.Id)
		http.HandleFunc("/"+a.Id, func(w http.ResponseWriter, r *http.Request) {
			action(w, r, a)
		})
	}
	//listen
	err := http.ListenAndServe(":"+strconv.Itoa(c.Port), nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

//index page
func index(w http.ResponseWriter, r *http.Request) {
	t := template.New("index")
	t, err := t.Parse(assets["_assets/index.html"])
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	t.Execute(w, conf)
}

//action page
func action(w http.ResponseWriter, r *http.Request, a Action) {
	fmt.Fprint(w, cmd(a.Path))
}

//exec command
func cmd(commandString string) string {
	var command string
	var op string
	if runtime.GOOS == "windows" {
		command = "cmd"
		op = "/c"
	} else {
		command = "/bin/sh"
		op = "-c"
	}
	out, err := exec.Command(command, op, commandString).Output()
	if err != nil {
		return string(err.Error())
	}
	return string(out)
}
