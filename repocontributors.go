package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/github"
	giturl "github.com/kubescape/go-git-url"
	"golang.org/x/oauth2"
)

// given a github URL return all the github username that are contributors
func repoContributors(project string, url string) (RepoContributors, error) {
	ctx := context.Background()

	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatal("GITHUB_TOKEN environment variable not set")
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	opt := &github.ListContributorsOptions{
		ListOptions: github.ListOptions{PerPage: 100}, // Set the page size to 100.
	}

	var contributorsList []string
	gitURL, err := giturl.NewGitURL(url)
	if err != nil {
		return RepoContributors{}, err
	}
	owner, repo := gitURL.GetOwnerName(), gitURL.GetRepoName()

	for {
		contributors, resp, err := client.Repositories.ListContributors(ctx, owner, repo, opt)
		if err != nil {
			fmt.Printf("something wrong with %v/%v\n", owner, repo)
			fmt.Println(err)
			break
		}

		for _, user := range contributors {
			contributorsList = append(contributorsList, *user.Login)
		}

		if resp.NextPage == 0 {
			break // We've reached the last page.
		}

		opt.Page = resp.NextPage // Move to the next page.
	}

	return RepoContributors{Name: project, Contributors: contributorsList}, nil
	//close(contributorchannel)
}
