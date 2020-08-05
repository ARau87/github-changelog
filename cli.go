package main

import (
	"github.com/andreasrau-earlynode/repo-changelog/actions"
	"github.com/urfave/cli/v2"
)

func SetupApp() *cli.App {

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "create",
				Aliases: []string{"c"},
				Usage:   "Create changelog",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "out",
						Aliases:  []string{"o"},
						Usage:    "Path to output file",
						Required: true,
					},
					&cli.StringFlag{
						Name:    "tag",
						Aliases: []string{"t"},
						Value:   "Changelog:",
						Usage:   "Tag in the issue body indicating infos for the changelog",
					},
					&cli.StringFlag{
						Name:     "owner",
						Aliases:  []string{"O"},
						Required: true,
						Usage:    "The owner of the repository",
					},
					&cli.StringFlag{
						Name:    "version",
						Aliases: []string{"v"},
						Usage:   "Version that should appear in the changelog",
					},
					&cli.StringFlag{
						Name:     "repo",
						Aliases:  []string{"r"},
						Required: true,
						Usage:    "The repository name",
					},
					&cli.StringFlag{
						Name:     "oauth",
						Aliases:  []string{"X"},
						Usage:    "Personal access token used to perform actions",
						Required: true,
						EnvVars:  []string{"OAUTH_TOKEN"},
					},
					&cli.Float64Flag{
						Name:    "sprint",
						Aliases: []string{"s"},
						Usage:   "Length of the sprint in days e.g. '-s 7' means last 7 days",
						Value:   7.0,
					},
					&cli.TimestampFlag{
						Name:    "since",
						Aliases: []string{"S"},
						Usage:   "Alternative to the sprint flag. With this flag you can specify the start date of the sprint in the format 2006-01-02T15:04:05",
						Layout:  "2006-01-02T15:04:05",
					},
				},
				Action: actions.GetIssues,
			},
		},
	}

	return app
}
