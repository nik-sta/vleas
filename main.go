package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

var app = cli.NewApp()

func main() {
	searchOnMavenCentral()

	info()
	commands()
	flags()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func info() {
	app.Name = "Vleas CLI"
	app.Usage = "Vleas is an easy to use open source CLI for maintaining dependencies."
	app.Version = "0.0.1"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Nikola Stanković",
			Email: "nikola@stankovic.xyz",
		},
	}
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "update",
			Aliases: []string{"u"},
			Usage:   "Update all dependencies to newest version",
			Action: func(c *cli.Context) {
				update()
			},
		},
		{
			Name:    "check",
			Aliases: []string{"c"},
			Usage:   "Check for new dependency updates",
			Action: func(c *cli.Context) {
				check()
			},
		},
	}
}

func flags()  {
	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "file, f",
			Usage: "Path to dependencies file `FILE`",
			Required: true,
		},
	}
}

func update()  {
	fmt.Println("update dependencies")
}

func check()  {
	fmt.Println("check dependencies")
}

func searchOnMavenCentral() {
	var groupId = "ch.viascom.groundwork"
	var artifactId = "foxhttp"

	qp := "q=g:\"" + groupId + `"+AND+a:"` + artifactId + `"` +
		"&rows=20" +
		"&core=gav"

	u := &url.URL{
		Scheme:   "http",
		Host:     "search.maven.org",
		Path:     "/solrsearch/select",
		RawQuery: qp,
	}

	resp, err := http.Get (u.String())
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resp.Request.URL)
	log.Println(resp)
}
