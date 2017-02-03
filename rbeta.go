package main

import (
	"os"
	"fmt"
	// "reflect"

	"gopkg.in/alecthomas/kingpin.v2"

	initCmd "github.com/likemindnetworks/rbeta/cmd/init"
)

const Version = "0.0.1"

var (
	rbeta = kingpin.New("rbeta", "rbeta")
	cmdInit = initCmd.Command(rbeta)
)

func main() {
	rbeta.Version(Version)

	switch kingpin.MustParse(rbeta.Parse(os.Args[1:])) {
	case cmdInit.FullCommand():
		fmt.Printf("!")
	}
}
