// pochtona.go
package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	current, err := filepath.Abs("./_test")
	if err != nil {
		print(err)
	}
	err = InitConf(current)
	if err != nil {
		fmt.Println(err)
	}
}
