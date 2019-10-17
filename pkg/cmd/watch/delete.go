package watch

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/urfave/cli"

	"github.com/suzuki-shunsuke/github-cli/pkg/util"
)

func Delete(c *cli.Context) error {
	ctx := context.Background()

	client := util.NewGitHubClient(ctx, c.GlobalString("token"))

	for _, arg := range c.Args() {
		a := strings.Split(arg, "/")
		if len(a) != 2 {
			return errors.New("invalid argument: " + arg)
		}
		resp, err := client.Activity.DeleteRepositorySubscription(ctx, a[0], a[1])
		if err != nil {
			return err
		}
		resp.Body.Close()
		if resp.StatusCode >= 400 {
			return errors.New("GitHub API's status code >= 400: " + strconv.Itoa(resp.StatusCode))
		}
	}

	return nil
}
