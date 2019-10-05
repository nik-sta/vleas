package main

import (
	"fmt"
	"github.com/rhea-0b1/vleas/gradle"
	"github.com/rhea-0b1/vleas/model"
	"github.com/urfave/cli"
	"log"
	"os"
	"sort"
	"strings"
)

var resolvedDependencies = make([]model.Dependency, 0)
var unresolvedDependencies = make([]model.Dependency, 0)

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
	buildTool := detectBuildTool(file)

	switch buildTool {
	case "Gradle":
		gradle.Handle(file)
		printResult()
	case "Maven":
		println("Sorry :( Maven will be supported soon.")
	default:
		println("Sorry :( We are not supporting this type of dependency file ...")
	}
}

func detectBuildTool(file string) string {
	var buildTool = ""

	if strings.Contains(file, ".gradle") {
		buildTool = "Gradle"
	}

	if strings.Contains(file, "pom.xml") {
		buildTool = "Maven"
	}

	return buildTool
}

func printResult() {
	resolvedDependencies = gradle.ResolvedDependencies
	unresolvedDependencies = gradle.UnresolvedDependencies

	if len(resolvedDependencies) > 0 {
		fmt.Printf("\nVleas found %d dependency update(s):\n\n", len(resolvedDependencies))
	} else {
		fmt.Printf("\nGreat! Your project is up to date :)")
	}

	for _, dep := range resolvedDependencies {
		fmt.Printf("group: %s name: %s version: %s --> %s\n", dep.Group, dep.Name, dep.CurrentVersion, dep.LatestVersion)
	}

	if len(unresolvedDependencies) > 0 {
		fmt.Printf("\nThe following dependencies have not been able to check:\n\n")
		for _, dep := range unresolvedDependencies {
			fmt.Printf("group: %s name: %s version: %s\n", dep.Group, dep.Name, dep.CurrentVersion)
		}
	}
}

func update(file string) {
	fmt.Println("update deps from file: " + file)
}
