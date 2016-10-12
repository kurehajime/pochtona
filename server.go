// server.go
package main

import (
	"fmt"
	"html/template"
	"net"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"
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
	server := &http.Server{}
	l, err := net.Listen("tcp4", ":"+strconv.Itoa(c.Port))
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	err = server.Serve(l)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

//index page
func index(w http.ResponseWriter, r *http.Request) {
	if checkIP(r.RemoteAddr, conf.AllowIpAddresses) == false {
		w.WriteHeader(403)
		w.Write([]byte("403 Forbidden\n"))
		return
	}
	t := template.New("index")
	t, err := t.Parse(assets["_assets/index.html"])
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	t.Execute(w, conf)
}

//action page
func action(w http.ResponseWriter, r *http.Request, a Action) {
	if checkIP(r.RemoteAddr, conf.AllowIpAddresses) == false {
		w.WriteHeader(403)
		w.Write([]byte("403 Forbidden\n"))
		return
	}
	fmt.Fprint(w, cmd("\""+a.Path+"\""))
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

//checkIP
func checkIP(addr string, allowIpAddresses []string) bool {
	addr = strings.Split(addr, ":")[0]
	re1 := regexp.MustCompile("\\.")
	re2 := regexp.MustCompile("\\*")
	for i := range allowIpAddresses {
		rule := allowIpAddresses[i]
		rule = re1.ReplaceAllString(rule, "\\.")
		rule = re2.ReplaceAllString(rule, ".*")
		if regexp.MustCompile(rule).Match([]byte(addr)) == true {
			return true
		}
	}
	return false
}
