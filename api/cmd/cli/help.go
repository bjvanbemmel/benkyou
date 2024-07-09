package main

import (
	"fmt"
	"io"

	_ "embed"
)

type CommandHelp struct{}

//go:embed help_content.txt
var content string

func (c CommandHelp) Execute(w io.ReadWriter, args ...string) error {
	fmt.Fprint(w, content)

	return nil
}
