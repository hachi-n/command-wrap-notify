package main

import (
	"fmt"
	"github.com/hachi-n/command-wrap-notify/internal/cmd/apply"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "command-wrap-notify",
		Usage: "my project template",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "command",
				Value:    "",
				Usage:    "command fullpath.",
				Required: true,
			},
		},
		Action: func(context *cli.Context) error {
			c := context.String("command")
			return apply.Apply(c)
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
