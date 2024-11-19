package main

import (
	"fmt"
	"os"

	"github.com/manasm11/calc/cli"
)

func main() {
	c := cli.NewCli()
	c.Run(os.Args)
	fmt.Print(c.Output())
}
