package main

import (
	"gitlet/cmd/gitlet/internal"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:     "Gitlet",
		Version:  "v00.00.1",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{
				Name:  "Chilly Forest",
				Email: "sunhuayangak47@gmail.com",
			},
		},
		Usage:     "a small git",
		UsageText: "Gitlet command [arguments...]",
		Commands: []*cli.Command{
			{
				Name:   "init",
				Usage:  "initialize repository",
				Action: internal.Init,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
