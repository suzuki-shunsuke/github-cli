package main

import (
	"log"
	"os"

	"github.com/suzuki-shunsuke/go-cliutil"
	"github.com/urfave/cli"

	"github.com/suzuki-shunsuke/github-cli/pkg/cmd/repos"
	"github.com/suzuki-shunsuke/github-cli/pkg/cmd/watching"
	"github.com/suzuki-shunsuke/github-cli/pkg/domain"
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
			Name:  "watching",
			Usage: "GitHub watching API",
			Subcommands: []cli.Command{
				{
					Name:   "set",
					Usage:  "set GitHub watching",
					Action: cliutil.WrapAction(watching.Set),
				},
			},
		},
		{
			Name:  "repos",
			Usage: "GitHub repositories API",
			Subcommands: []cli.Command{
				{
					Name:   "list",
					Usage:  "list GitHub repositories",
					Action: cliutil.WrapAction(repos.List),
					Flags: []cli.Flag{
						cli.StringFlag{
							Name: "type",
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
