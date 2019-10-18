package repos

import (
	"context"

	"github.com/google/go-github/v28/github"
	"github.com/urfave/cli"

	"github.com/suzuki-shunsuke/github-cli/pkg/util"
)

func ListYourRepos(c *cli.Context) error {
	ctx := context.Background()
	client := util.NewGitHubClient(ctx, c.GlobalString("token"))

	baseOpt := github.RepositoryListOptions{
		Visibility:  c.String("visibility"),
		Affiliation: c.String("affiliation"),
		Type:        c.String("type"),
		Sort:        c.String("sort"),
		Direction:   c.String("direction"),
	}

	return listRepos(
		ctx, "", func(ctx context.Context, _ string, page int) (
			[]*github.Repository, *github.Response, error,
		) {
			opt := baseOpt
			opt.Page = page
			return client.Repositories.List(ctx, "", &opt)
		})
}
