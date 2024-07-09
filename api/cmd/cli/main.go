package main

import (
	"fmt"
	"io"
	"os"
)

type Command interface {
	Execute(io.ReadWriter, ...string) error
}

func main() {
	if len(os.Args[1:]) < 1 {
		err := CommandHelp{}.Execute(os.Stdout)
		if err != nil {
			panic(err)
		}

		return
	}

	command := os.Args[1]
	switch command {
	default:
		fallthrough
	case "help":
		err := CommandHelp{}.Execute(os.Stdout)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	case "migrate":
		err := CommandMigrate{}.Execute(os.Stdout, os.Args[2:]...)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
