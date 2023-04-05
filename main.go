package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var ctx = context.Background()
var ts = oauth2.StaticTokenSource(
	&oauth2.Token{AccessToken: os.Getenv("GH_TOKEN")},
)
var tc = oauth2.NewClient(ctx, ts)
var client = github.NewClient(tc)

func main() {
	args := os.Args
	args_len := len(args)

	listUserOrganization()
	if args_len > 1 {
		repoState := convertStringToBool(args[2])
		createRepository(args[1], repoState)
	}

}

func listUserOrganization() {

	orgs, _, err := client.Organizations.List(ctx, os.Getenv("GH_USER"), nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(orgs)

}

func createRepository(repName string, repVisibility bool) {
	repo := &github.Repository{
		Name:    github.String(repName),
		Private: github.Bool(repVisibility),
	}

	createdRepo, _, err := client.Repositories.Create(ctx, "", repo)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*createdRepo.HTMLURL)
}

func convertStringToBool(str string) bool {
	boolValue, err := strconv.ParseBool(str)
	if err != nil {
		fmt.Println(err)
	}
	return boolValue
}
