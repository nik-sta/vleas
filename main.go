package main

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"strings"
	"time"

	. "github.com/logrusorgru/aurora"
)

var app = cli.NewApp()

var filePath string

func main() {
	//var groupId = "ch.viascom.groundwork"
	//var artifactId = "foxhttp"
	//searchOnMavenCentral(groupId, artifactId)

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
		{
			Name:    "search",
			Aliases: []string{"s"},
			Usage:   "Search for a dependency",
			Action: func(c *cli.Context) {
				search()
			},
		},
	}
}

func flags()  {
	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "file, f",
			Usage: "Path to dependencies file `FILE`",
			Destination: &filePath,
		},
	}
}

func update()  {
	fmt.Println("update dependencies")
}

func check()  {
	var totalDeps int
	configurations := []string{"api", "implementation", "compileOnly", "runtimeOnly", "testImplementation", "testCompileOnly", "testRuntimeOnly", "compile", "runtime"}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, configuration := range configurations {
			if strings.Contains(scanner.Text(), configuration + " ") {

				totalDeps++

				dep := scanner.Text()
				dep = strings.Replace(dep, configuration, "", 1)
				dep = strings.ReplaceAll(dep, "\"", "")
				dep = strings.ReplaceAll(dep, "'", "")
				dep = strings.TrimSpace(dep)

				dependencies := strings.Split(dep, ":")

				if len(dependencies) == 2 {
					totalDeps--
					continue
				}

				if len(dependencies) >= 1 {
					fmt.Print("group: ", Yellow(dependencies[0]), ", ")

					if len(dependencies) >= 2 {
						fmt.Print("name: ", Yellow(dependencies[1]), ", ")

						if len(dependencies) >= 3 {
							fmt.Print("version: ", Yellow(dependencies[2]),  " --> ")

							fmt.Println(Red("outdated"))
						}
					}
				}
			}
		}
	}

	fmt.Println("\nTotal dependencies found:", totalDeps)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func search()  {
	fmt.Println("search dependency")
}

//func searchOnMavenCentral(groupId, artifactId string) {
//	var url = "http://search.maven.org/solrsearch/select?q=g:%22GROUP_ID%22+AND+a:%22ARTIFACT_ID%22&rows=20&core=gav"
//	url = strings.Replace(url,"GROUP_ID", groupId, -1)
//	url = strings.Replace(url,"ARTIFACT_ID", artifactId, -1)
//
//	resp, err := http.Get (url)
//	if err != nil {
//		panic(err.Error())
//	}
//}