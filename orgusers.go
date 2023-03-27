package main

import (
	"context"
	"log"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func orgUsers(orguserchannel chan []string) {
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

	opt := &github.ListMembersOptions{
		ListOptions: github.ListOptions{PerPage: 100}, // Set the page size to 100.
	}
	org := os.Getenv("GITHUB_ORG")
	if org == "" {
		log.Fatal("GITHUB_ORG env var is not set")
	}
	var userslice []string
	for {
		users, resp, err := client.Organizations.ListMembers(ctx, org, opt)
		if err != nil {
			log.Fatal(err)
		}

		for _, user := range users {
			userslice = append(userslice, *user.Login)
		}

		if resp.NextPage == 0 {
			break // We've reached the last page.
		}

		opt.Page = resp.NextPage // Move to the next page.
	}

	orguserchannel <- userslice
	close(orguserchannel)
}
