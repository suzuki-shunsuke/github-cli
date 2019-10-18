package repos

import (
	"context"
	"errors"

	"github.com/google/go-github/v28/github"
	"github.com/urfave/cli"

	"github.com/suzuki-shunsuke/github-cli/pkg/util"
)

func ListUserRepos(c *cli.Context) error {
	ctx := context.Background()
	client := util.NewGitHubClient(ctx, c.GlobalString("token"))

	baseOpt := github.RepositoryListOptions{
		Type:      c.String("type"),
		Sort:      c.String("sort"),
		Direction: c.String("direction"),
	}

	user := c.Args().First()
	if user == "" {
		return errors.New("user name is required")
	}

	return listRepos(
		ctx, user, func(ctx context.Context, key string, page int) (
			[]*github.Repository, *github.Response, error,
		) {
			opt := baseOpt
			opt.Page = page
			return client.Repositories.List(ctx, key, &opt)
		})
}
