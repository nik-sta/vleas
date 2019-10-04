package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
	"sort"
)

func main() {
	app := cli.NewApp()
	app.Name = "vleas"
	app.Usage = "be always up to date, extremely fast ;)"
	app.Version = "0.0.1"
	app.Author = "Nikola Stanković"
	app.Email = "nikola@stankovic.xyz"
	app.Description = "Vleas is an easy to use open source CLI for maintaining dependencies."

	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "file, f",
			Usage: "Load dependencies from `FILE`",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "check",
			Aliases: []string{"c"},
			Usage:   "check for new dependencies",
			Action:  func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:    "update",
			Aliases: []string{"u"},
			Usage:   "update all dependencies to latest version",
			Action:  func(c *cli.Context) error {
				return nil
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
