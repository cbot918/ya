package cli

import (
	"fmt"
)

type Help struct {}

func NewHelp() *Help{
	h := new(Help)
	return h
}

func (h *Help) Get() string{
	return fmt.Sprintf("command:\nlist\ninstall\ndelete\ninstalled")
}