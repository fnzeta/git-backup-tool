package main

import (
	"log"
	"os"

	"github.com/xanzy/go-gitlab"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

func main() {

	// Check if GITHUBTOKEN is present in env
	token := os.Getenv("TOKEN")
	if token == "" {
		log.Fatalf("No github token found")
	}

	group := os.Getenv("GROUP")
	if group == "" {
		log.Fatalf("No group found")
	}

	lab := gitlab.NewClient(nil, token)

	listProjectsOptions := &gitlab.ListProjectsOptions{
		// Owned: gitlab.Bool(true),
		Membership: gitlab.Bool(true),
	}
	projects, _, err := lab.Projects.ListProjects(listProjectsOptions)
	if err != nil {
		log.Fatal(err)
	}

	for _, project := range projects {
		dir, errx := os.Getwd()

		//check if you need to panic, fallback or report
		if errx != nil {
			log.Fatal(errx)
		}
		_, err := git.PlainClone(dir+"/backup/"+project.Name, false, &git.CloneOptions{
			Auth: &http.BasicAuth{
				Username: "abc123", // yes, this can be anything except an empty string
				Password: token,
			},
			URL:      project.HTTPURLToRepo,
			Progress: os.Stdout,
		})
		if err != nil {
			log.Fatal(err)
		}
	}

	listGroupProjectsOptions := &gitlab.ListGroupProjectsOptions{}
	listGroupProjects, _, err := lab.Groups.ListGroupProjects(group, listGroupProjectsOptions)
	if err != nil {
		log.Fatal(err)
	}
	for _, project := range listGroupProjects {

		dir, errx := os.Getwd()

		//check if you need to panic, fallback or report
		if errx != nil {
			log.Fatal(errx)
		}
		_, err := git.PlainClone(dir+"/backup/"+project.Name, false, &git.CloneOptions{
			Auth: &http.BasicAuth{
				Username: "abc123", // yes, this can be anything except an empty string
				Password: token,
			},
			URL:      project.HTTPURLToRepo,
			Progress: os.Stdout,
		})
		if err != nil {
			log.Fatal(err)
		}
	}

}
