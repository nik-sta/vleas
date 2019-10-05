package gradle

import (
	"github.com/rhea-0b1/vleas/model"
	"github.com/rhea-0b1/vleas/util"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

var ResolvedDependencies = make([]model.Dependency, 0)
var UnresolvedDependencies = make([]model.Dependency, 0)

func Handle(file string) {
	contentBytes, _ := ioutil.ReadFile(file)
	content := string(contentBytes)

	regex := regexp.MustCompile("(?P<group>[^\"$,;()\\[\\]{}']+):(?P<name>[^\"$,;()\\[\\]{}']+):(?P<version>[^\"$,;()\\[\\]{}']+)")
	deps := regex.FindAllStringSubmatch(content, -1)

	for i := range deps {
		group := deps[i][1]
		name := deps[i][2]
		latestVersion := fetchLatestDeps(group, name)

		deps[i] = append(deps[i], latestVersion)

		if validDep(deps[i]) {
			ResolvedDependencies = append(ResolvedDependencies, model.Dependency{
				Group:          deps[i][1],
				Name:           deps[i][2],
				CurrentVersion: deps[i][3],
				LatestVersion:  deps[i][4],
			})
		} else {
			if strings.EqualFold(deps[i][3], deps[i][4]) == false {
				UnresolvedDependencies = append(UnresolvedDependencies, model.Dependency{
					Group:          deps[i][1],
					Name:           deps[i][2],
					CurrentVersion: deps[i][3],
					LatestVersion:  deps[i][4],
				})
			}
		}
	}
	ResolvedDependencies = util.RemoveDuplicates(ResolvedDependencies)
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