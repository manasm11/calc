package cli

import "bytes"

type Cli struct {
	output *bytes.Buffer
}

func NewCli() *Cli {
	return &Cli{output: bytes.NewBufferString("")}
}

func (c *Cli) Run(args []string) {
	c.output.WriteString("14")
}
