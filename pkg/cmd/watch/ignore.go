package watch

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/google/go-github/v28/github"
	"github.com/urfave/cli"

	"github.com/suzuki-shunsuke/github-cli/pkg/util"
)

func setSubscription(c *cli.Context, ignored bool) error {
	ctx := context.Background()
	client := util.NewGitHubClient(ctx, c.GlobalString("token"))
	f := true
	subsc := github.Subscription{}
	if ignored {
		subsc.Ignored = &f
	} else {
		subsc.Subscribed = &f
	}
	for _, arg := range c.Args() {
		a := strings.Split(arg, "/")
		if len(a) != 2 {
			return errors.New("invalid argument: " + arg)
		}

		_, resp, err := client.Activity.SetRepositorySubscription(ctx, a[0], a[1], &subsc)
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

func Ignore(c *cli.Context) error {
	return setSubscription(c, true)
}
