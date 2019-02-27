package main

import (
	"os"
	"github.com/greyspectrum/gocat/tcp"
)

func main() {
	os.Exit(command.Run(os.Args[1:]))
}
