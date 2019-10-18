package repos

import (
	"context"
	"errors"

	"github.com/google/go-github/v28/github"
	"github.com/urfave/cli"

	"github.com/suzuki-shunsuke/github-cli/pkg/util"
)

func ListOrgRepos(c *cli.Context) error {
	ctx := context.Background()
	client := util.NewGitHubClient(ctx, c.GlobalString("token"))

	baseOpt := github.RepositoryListByOrgOptions{
		Type:      c.String("type"),
		Sort:      c.String("sort"),
		Direction: c.String("direction"),
	}

	org := c.Args().First()
	if org == "" {
		return errors.New("organization name is required")
	}

	return listRepos(
		ctx, org, func(ctx context.Context, key string, page int) (
			[]*github.Repository, *github.Response, error,
		) {
			opt := baseOpt
			opt.Page = page
			return client.Repositories.ListByOrg(ctx, key, &opt)
		})
}
