package repos

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/oauth2"

	"github.com/google/go-github/v28/github"
	"github.com/urfave/cli"
)

func List(c *cli.Context) error {
	ctx := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: c.GlobalString("token")},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	if len(c.Args()) == 0 {
		// your repositories
		return listRepos(
			ctx, "", func(ctx context.Context, key string, page int) (
				[]*github.Repository, *github.Response, error,
			) {
				return client.Repositories.List(ctx, "", &github.RepositoryListOptions{
					ListOptions: github.ListOptions{Page: page},
				})
			})
	}
	arg := c.Args().First()
	if strings.HasPrefix(arg, "users/") {
		// user's repositories
		return listRepos(
			ctx, arg[len("users/"):],
			func(ctx context.Context, key string, page int) (
				[]*github.Repository, *github.Response, error,
			) {
				return client.Repositories.List(ctx, key, &github.RepositoryListOptions{
					ListOptions: github.ListOptions{Page: page},
				})
			})
	}
	if strings.HasPrefix(arg, "orgs/") {
		// org's repositories
		return listRepos(
			ctx, arg[len("orgs/"):],
			func(ctx context.Context, key string, page int) (
				[]*github.Repository, *github.Response, error,
			) {
				return client.Repositories.ListByOrg(
					ctx, key, &github.RepositoryListByOrgOptions{
						ListOptions: github.ListOptions{Page: page},
					})
			})
	}
	if arg == "all" {
		// all repositories
		return listRepos(
			ctx, arg[len("users/"):],
			func(ctx context.Context, key string, page int) (
				[]*github.Repository, *github.Response, error,
			) {
				return client.Repositories.ListAll(ctx, &github.RepositoryListAllOptions{
					Since: int64(page),
				})
			})
	}
	return errors.New("unexpected argument: " + arg)
}

func listRepos(
	ctx context.Context, key string,
	list func(ctx context.Context, key string, page int) ([]*github.Repository, *github.Response, error),
) error {
	page := 0
	for {
		repos, resp, err := list(ctx, key, page)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		if resp.StatusCode >= 400 {
			return errors.New("GitHub API's status code >= 400: " + strconv.Itoa(resp.StatusCode))
		}
		if len(repos) == 0 {
			break
		}
		names := make([]string, len(repos))
		for i, repo := range repos {
			names[i] = *repo.FullName
		}
		fmt.Println(strings.Join(names, "\n"))
		if resp.NextPage == 0 {
			break
		}
		page = resp.NextPage
	}
	return nil
}
