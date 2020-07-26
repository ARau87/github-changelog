package actions

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
	"time"

	"github.com/google/go-github/github"
	"github.com/urfave/cli/v2"
	"golang.org/x/oauth2"
)

type ChangeLogData struct {
	Version        string
	ChangeLogTexts []string
}

func GetIssues(ctx *cli.Context) error {

	// Setup Guthub Client
	con := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ctx.String("oauth")},
	)
	tc := oauth2.NewClient(con, ts)

	client := github.NewClient(tc)

	// Find Pull Requests
	opt := &github.PullRequestListOptions{State: "closed"}
	pullrequests, _, err := client.PullRequests.List(con, ctx.String("owner"), ctx.String("repo"), opt)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	// Create file and setup the template file
	file, err := createFile(ctx.String("out"))
	if err != nil {
		log.Fatalln(err)
		return err
	}
	defer file.Close()
	tmpl := setupTemplate()

	changeLogData := &ChangeLogData{
		Version: ctx.String("version"),
	}

	for _, pullrequest := range pullrequests {
		if pullrequest.MergedAt != nil && pullrequest.MergedAt.After(getSprintStartDate(ctx)) {
			parsePullRequestBody(ctx.String("tag"), pullrequest, changeLogData)
		}
	}

	tmpl.Execute(file, changeLogData)

	return nil
}

func parsePullRequestBody(tag string, pullrequest *github.PullRequest, changeLogData *ChangeLogData) {

	regExpression := regexp.MustCompile(fmt.Sprintf(`\b%s+(.*)`, tag))
	changeLogs := regExpression.FindAll([]byte(pullrequest.GetBody()), 10)
	for _, changeLog := range changeLogs {
		changeLogString := strings.Replace(string(changeLog), fmt.Sprintf(`%s `, tag), "", 1)
		changeLogData.ChangeLogTexts = append(changeLogData.ChangeLogTexts, changeLogString)
	}
}

func setupTemplate() *template.Template {

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	templateFile := path.Join(dir, "template.tmpl")
	tmpl := template.Must(template.ParseFiles(templateFile))
	return tmpl

}

func getSprintStartDate(ctx *cli.Context) time.Time {

	if ctx.Timestamp("since") != nil {
		return *ctx.Timestamp("since")
	}
	return time.Now().Add(time.Duration(-24.0*ctx.Float64("sprint")) * time.Hour)
}

func createFile(out string) (*os.File, error) {

	currentDir, err := os.Getwd()
	file, err := os.Create(path.Join(currentDir, out))
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return file, nil

}
