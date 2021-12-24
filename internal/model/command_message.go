package model

import (
	"encoding/json"
	"fmt"
	"strings"
)

type CommandMessage struct {
	Command  string
	ExitCode int
	Error    string
}

func (m *CommandMessage) PrettyJson() []byte {
	b, err := json.MarshalIndent(m, "", strings.Repeat(" ", 4))
	if err != nil {
		fmt.Println(err)
	}
	return b
}

func NewCommandMessage(command string, exitCode int, err error) *CommandMessage {
	return &CommandMessage{Command: command, ExitCode: exitCode, Error: err.Error()}
}
