package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"
)

var dependencies = make([][]string, 0)

func main() {
	app := cli.NewApp()
	app.Name = "vleas"
	app.Usage = "be always up to date, extremely fast ;)"
	app.Version = "0.0.1"
	app.Author = "Nikola Stanković"
	app.Email = "nikola@stankovic.xyz"
	app.Description = "Vleas is an easy to use open source CLI for maintaining deps."

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "file, f",
			Usage: "Load deps from `FILE`",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "check",
			Aliases: []string{"c"},
			Usage:   "check for new deps",
			Action: func(c *cli.Context) error {
				check(c.GlobalString("file"))
				return nil
			},
		},
		{
			Name:    "update",
			Aliases: []string{"u"},
			Usage:   "update all deps to latest version",
			Action: func(c *cli.Context) error {
				update(c.GlobalString("file"))
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

func check(file string) {
	contentBytes, _ := ioutil.ReadFile(file)
	content := string(contentBytes)

	regex := regexp.MustCompile("(?P<group>[^\"$,;\\(\\)\\[\\]\\{\\}']+):(?P<name>[^\"$,;\\(\\)\\[\\]\\{\\}']+):(?P<version>[^\"$,;\\(\\)\\[\\]\\{\\}']+)")
	deps := regex.FindAllStringSubmatch(content, -1)

	for i := range deps {
		group := deps[i][1]
		name := deps[i][2]
		latestVersion := fetchLatestDeps(group, name)

		deps[i] = append(deps[i], latestVersion)

		if validDep(deps[i]) {
			fmt.Printf("group: %s name: %s version: %s --> %s\n", deps[i][1], deps[i][2], deps[i][3], deps[i][4])
			dependencies = append(dependencies, deps[i])
		}
	}
}

func validDep(dep []string) bool {
	currentVersion := dep[3]
	latestVersion := dep[4]

	if strings.EqualFold(currentVersion, latestVersion) {
		return false
	}

	if latestVersion == "" {
		return false
	}

	return true
}

func update(file string) {
	fmt.Println("update deps from file: " + file)
}

func fetchLatestDeps(group, name string) string {
	url := "http://search.maven.org/solrsearch/select?q=g:%22#GROUP%22+AND+a:%22#NAME%22&"
	url = strings.Replace(url, "#GROUP", group, 1)
	url = strings.Replace(url, "#NAME", name, 1)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return gjson.Get(string(body), "response.docs.0.latestVersion").String()
}