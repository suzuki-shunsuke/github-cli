package repos

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/google/go-github/v28/github"
)

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
		resp.Body.Close()
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
