package main

import (
	"fmt"
	"sync"

	intersect "github.com/juliangruber/go-intersect"
)

func main() {

	channel := make(chan []Project)
	orguserchannel := make(chan []string)
	go cncfProjects(channel)
	cncfProjects := <-channel
	go orgUsers(orguserchannel)
	orgusers := <-orguserchannel
	contributorsCh := make(chan []string)

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

	for contributor := range contributorsCh {
		contributors_from_org := intersect.Simple(orgusers, contributor)
		if len(contributors_from_org) > 0 {
			fmt.Printf("%v\n", contributors_from_org)
		}
	}

}
