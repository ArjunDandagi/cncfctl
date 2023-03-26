package main

import (
	"fmt"

	intersect "github.com/juliangruber/go-intersect"
)

func main() {
	cncfProjects := cncfProjects()
	orgusers := orgUsers()
	for _, p := range cncfProjects {
		name, url := p.Name, p.Repo
		var repo_contributors []string
		if url != "" {
			repo_contributors = repoContributors(url)
		}
		contributors_from_org := intersect.Simple(orgusers, repo_contributors)
		if len(contributors_from_org) > 0 {
			fmt.Printf("%v,%v\n", name, contributors_from_org)
		}

	}

}
