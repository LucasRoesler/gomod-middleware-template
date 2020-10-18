package pkg

import (
	"context"
	"log"
	"os"

	"github.com/google/go-github/github"
)

func GetReposByOrg(org string) ([]string, error) {
	urls := []string{}
	log.Printf("Version %s", os.Getenv("VERSION"))

	client := github.NewClient(nil)

	// list public repositories for org "github"
	opt := &github.RepositoryListByOrgOptions{Type: "public"}
	repos, _, err := client.Repositories.ListByOrg(context.Background(), org, opt)

	if err != nil {
		return urls, err
	}

	for _, r := range repos {
		log.Printf("Found repo: %s", r.GetName())
		urls = append(urls, r.GetCloneURL())
	}

	return urls, nil
}
