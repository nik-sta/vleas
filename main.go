package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

var app = cli.NewApp()

func main() {
	info()
	commands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func info() {
	app.Name = "Vleas CLI"
	app.Usage = "Vleas is a open source CLI for updating Gradle and Maven dependencies."
	app.Author = "Nikola Stanković"
	app.Version = "0.0.1"
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "update",
			Aliases: []string{"u"},
			Usage:   "Update dependencies",
			Action: func(c *cli.Context) {
				fmt.Println("update")
			},
		},
	}
}
