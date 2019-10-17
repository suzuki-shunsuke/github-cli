package repos

import (
	"context"

	"github.com/google/go-github/v28/github"
	"github.com/urfave/cli"

	"github.com/suzuki-shunsuke/github-cli/pkg/util"
)

func ListAllRepos(c *cli.Context) error {
	ctx := context.Background()
	client := util.NewGitHubClient(ctx, c.GlobalString("token"))

	baseOpt := github.RepositoryListAllOptions{
		Since: c.Int64("since"),
	}

	return listRepos(
		ctx, "", func(ctx context.Context, key string, page int) (
			[]*github.Repository, *github.Response, error,
		) {
			opt := baseOpt
			opt.Since = int64(page)
			return client.Repositories.ListAll(ctx, &opt)
		})
}
