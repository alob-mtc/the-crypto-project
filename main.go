package main

import (
	"os"

	"github.com/workspace/the-crypto-project/cli"
)


func main() {
	defer os.Exit(0)
	cli := &cli.CommandLine{}
	cli.Run()
}
