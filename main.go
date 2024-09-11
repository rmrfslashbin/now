package main

import (
	"time"

	"github.com/alecthomas/kong"
)

type Context struct {
}

type CLI struct {
}

func (r *CLI) Run(ctx *Context) error {
	now := time.Now().UTC().Format(time.RFC3339)
	println(now)
	return nil
}

func main() {
	// Parse the command line
	var cli CLI
	ctx := kong.Parse(&cli)

	// Call the Run() method of the selected parsed command.
	err := ctx.Run(&Context{})

	// FatalIfErrorf terminates with an error message if err != nil
	ctx.FatalIfErrorf(err)
}
