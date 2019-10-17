package util

import (
	"context"

	"golang.org/x/oauth2"

	"github.com/google/go-github/v28/github"
)

func NewGitHubClient(ctx context.Context, token string) *github.Client {
	return github.NewClient(oauth2.NewClient(
		ctx, oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})))
}
