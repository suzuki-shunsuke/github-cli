package main

import (
	"log"
	"os"

	"github.com/suzuki-shunsuke/go-cliutil"
	"github.com/urfave/cli"

	"github.com/suzuki-shunsuke/github-cli/pkg/cmd/repos"
	"github.com/suzuki-shunsuke/github-cli/pkg/cmd/watch"
	"github.com/suzuki-shunsuke/github-cli/pkg/domain"
)

var (
	typeFlag = cli.StringFlag{
		Name: "type,t",
	}
	sortFlag = cli.StringFlag{
		Name: "sort",
	}
	directionFlag = cli.StringFlag{
		Name: "direction,d",
	}
)

func main() {
	app := cli.NewApp()
	app.Name = "github"
	app.Version = domain.Version
	app.Author = "suzuki-shunsuke https://github.com/suzuki-shunsuke"
	app.Usage = "GitHub API CLI"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:     "token",
			Usage:    "GitHub Token",
			EnvVar:   "GITHUB_TOKEN",
			Required: true,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "watch",
			Usage: "GitHub watching API",
			Subcommands: []cli.Command{
				{
					Name:   "set",
					Usage:  "Set a Repository Subscription",
					Action: cliutil.WrapAction(watch.Set),
					Flags: []cli.Flag{
						cli.BoolFlag{
							Name: "subscribed,s",
						},
						cli.BoolFlag{
							Name: "ignored,i",
						},
					},
				},
				{
					Name:   "delete",
					Usage:  "Delete a Repository Subscription",
					Action: cliutil.WrapAction(watch.Delete),
				},
			},
		},
		{
			Name:  "repos",
			Usage: "GitHub repositories API",
			Subcommands: []cli.Command{
				{
					Name:   "list",
					Usage:  "List your repositories",
					Action: cliutil.WrapAction(repos.ListYourRepos),
					Flags: []cli.Flag{
						cli.StringFlag{
							Name: "visibility,v",
						},
						cli.StringFlag{
							Name: "affiliation,a",
						},
						typeFlag, sortFlag, directionFlag,
					},
				},
				{
					Name:   "list-user",
					Usage:  "List user repositories",
					Action: cliutil.WrapAction(repos.ListUserRepos),
					Flags: []cli.Flag{
						typeFlag, sortFlag, directionFlag,
					},
				},
				{
					Name:   "list-org",
					Usage:  "List organization repositories",
					Action: cliutil.WrapAction(repos.ListOrgRepos),
					Flags: []cli.Flag{
						typeFlag, sortFlag, directionFlag,
					},
				},
				{
					Name:   "list-all",
					Usage:  "List all public repositories",
					Action: cliutil.WrapAction(repos.ListAllRepos),
					Flags: []cli.Flag{
						cli.Int64Flag{
							Name: "since",
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
