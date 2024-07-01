package utils

import (
	"context"

	"github.com/aliceh/alertops/pkg/config"
	"github.com/google/go-github/v50/github"
	"golang.org/x/oauth2"
)

func GetGHReadme(owner, repo, path string) (string, error) {
	cfg, err := config.LoadConfig(config.PathOsdctl)
	if err != nil {
		return "", err
	}
	// Use Backgound Context
	ctx := context.Background()

	// Generate Token Source and Token Client
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: cfg.AccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	// Create GitHub Client
	client := github.NewClient(tc)
	options := github.RepositoryContentGetOptions{}

	// Get Contents Accordingly
	content, _, _, err := client.Repositories.GetContents(ctx, owner, repo, path, &options)
	if err != nil {
		return "", err
	}
	decodedContent, err := content.GetContent()
	if err != nil {
		return "", err
	}
	return decodedContent, nil
}
