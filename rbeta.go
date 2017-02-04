package main

import (
	"os"
	"fmt"

	"gopkg.in/alecthomas/kingpin.v2"

	initCmd "github.com/likemindnetworks/rbeta/cmd/init"
)

const Version = "0.0.1"

var (
	rbeta = kingpin.New("rbeta", "rbeta")
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Error: %s \n", r)
		}
	}()

	rbeta.Version(Version)

	initCmd.ConfigCommand(rbeta, "init")

	kingpin.MustParse(rbeta.Parse(os.Args[1:]))
}
