// server.go
package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type ServerConf struct {
	Port   int
	Config Conf
}

func Start(sc ServerConf) {
	//start server
	url := "http://localhost" + ":" + strconv.Itoa(sc.Port)
	fmt.Println(url)
	fmt.Println("Stop: Ctrl+C")
	//Handle
	http.HandleFunc("/", index)
	for _, a := range sc.Config.Actions {
		http.HandleFunc("/"+a.Id, func(w http.ResponseWriter, r *http.Request) {
			action(w, r, a)
		})
	}
	//listen
	err := http.ListenAndServe(":"+strconv.Itoa(sc.Port), nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
func index(w http.ResponseWriter, r *http.Request) {

}

func action(w http.ResponseWriter, r *http.Request, a Action) {

}
