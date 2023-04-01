package main

import (
	"sync"

	"github.com/juliangruber/go-intersect"
)

type RepoContributors struct {
	Name         string   `json:"name"`
	Contributors []string `json:"contributors"`
}

func main() {

	channel := make(chan []Project)
	orguserchannel := make(chan []string)
	go cncfProjects(channel)
	cncfProjects := <-channel
	go orgUsers(orguserchannel)
	orgusers := <-orguserchannel
	contributorsCh := make(chan RepoContributors)

	var wg sync.WaitGroup
	for _, project := range cncfProjects {

		wg.Add(1)

		go func(project string, repo string) {
			defer wg.Done()

			if repo != "https://www.github.com/deckhouse/deckhouse" {
				contributors, err := repoContributors(project, repo)
				if err != nil {
					return
				}
				contributorsCh <- contributors
			}

		}(project.Name, project.Repo)
	}

	go func() {
		wg.Wait()
		close(contributorsCh)
	}()
	var table_input []RepoContributors
	for contributor := range contributorsCh {
		contributors_from_org := intersect.Simple(orgusers, contributor.Contributors)
		if len(contributors_from_org) > 0 {
			table_input = append(table_input, RepoContributors{contributor.Name, convertToStringSlice(contributors_from_org)})
		}
	}
	customtable(table_input)
}
