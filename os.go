package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	host, _ := os.Hostname()
	fmt.Println(host)

	path, err := exec.LookPath("gocod")
	if err != nil {
		log.Fatal("wrong enev")
	}
	fmt.Println(path)
}
