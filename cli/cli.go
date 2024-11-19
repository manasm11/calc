package cli

import (
	"bytes"
	"strconv"

	"github.com/manasm11/calc/calc"
)

const usage = `
Usage:
	calc add <n1> <n2>...
	calc mul <n1> <n2>...
	calc sub <n1> <n2>
`

type Cli struct {
	output *bytes.Buffer
}

func NewCli() *Cli {
	return &Cli{output: bytes.NewBufferString("")}
}

func (c *Cli) Run(args []string) {
	nums := []int{}
	var isErr bool
	for _, arg := range args[2:] {
		n, err := strconv.Atoi(arg)
		if err != nil {
			if isErr {
				c.output.WriteString("\n")
			}
			isErr = true
			c.output.WriteString("invalid integer '" + arg + "'")
			continue
		}
		nums = append(nums, n)
	}
	if isErr {
		c.output.WriteString(usage)
		return
	}
	res := calc.Add(nums...)
	c.output.WriteString(strconv.Itoa(res))
}
