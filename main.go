package main

import (
	"fmt"
	"time"

	"github.com/alecthomas/kong"
)

var version = "1.0.0" // This should be set during the build process

type CLI struct {
	Version kong.VersionFlag `help:"Print version information and exit"`
}

func (cli *CLI) Run() error {
	now := time.Now().UTC().Format(time.RFC3339)
	fmt.Println(now)
	return nil
}

func main() {
	cli := CLI{}

	ctx := kong.Parse(&cli,
		kong.Description("A simple CLI that prints the current UTC time in RFC3339 format."),
		kong.UsageOnError(),
		kong.Vars{"version": version},
	)

	err := cli.Run()
	ctx.FatalIfErrorf(err)
}
