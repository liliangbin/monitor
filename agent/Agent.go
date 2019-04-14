package agent

import (
	"os/exec"
	"strings"
)

type Command struct {
	Cmd  string
	Args string
}

func (cmd *Command) Execute() ([]byte, error) {

	var argArray []string
	if cmd.Args != "" {
		argArray = strings.Split(cmd.Cmd+" "+cmd.Args, " ")
	} else {
		argArray = make([]string, 0)
	}
	_cmd := exec.Command("/bin/sh", argArray...)
	return _cmd.Output()
}
