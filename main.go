package main

import (
	// "errors"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "domain-tool",
		Usage: "ascertain relevant sysadmin info about a domain name",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "recursive",
				Usage:   "If set to true, send a recursive query (ANY). This is blocked by some nameservers.",
				Aliases: []string{"r"},
			},
			&cli.StringFlag{
				Name:     "domain",
				Usage:    "The fully qualified domain name",
				Aliases:  []string{"d"},
				Required: true,
			},
		},
		Action: func(ctx *cli.Context) error {
			domain := ctx.String("domain")
			isRecursive := ctx.Bool("recursive")

			fmt.Println(domain, isRecursive)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
