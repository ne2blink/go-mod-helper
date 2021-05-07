package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	baseURL = "https://api.github.com/repos/"
	commits = "/commits"
	version = "v0.0.0-"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no link")
		os.Exit(1)
	}
	link := removeHttp(os.Args[1])
	org, err := isGithub(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	branch := commits
	branchArray := strings.Split(org, "#")
	if len(branchArray) > 2 {
		fmt.Println("link wrong format")
		os.Exit(1)
	}
	if len(branchArray) == 2 {
		branch = commits + "/" + branchArray[1]
	}
	url := baseURL + branchArray[0] + branch
	body, err := httpGet(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	commits, err := formatJson(body, len(branchArray) == 1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(link + " " + version + commits.Date + "-" + commits.SHA[:12])
}

func removeHttp(url string) string {
	url = strings.TrimLeft(url, "http://")
	url = strings.TrimLeft(url, "https://")
	return url
}

func isGithub(url string) (string, error) {
	if strings.HasPrefix(url, "github.com/") || strings.HasPrefix(url, "http://github.com/") || strings.HasPrefix(url, "https://github.com/") {
		url = strings.TrimLeft(url, "github.com/")
		url = strings.TrimLeft(url, "http://github.com/")
		url = strings.TrimLeft(url, "https://github.com/")
	} else {
		return "", errors.New("not github link")
	}
	return url, nil
}

func httpGet(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}
	return ioutil.ReadAll(resp.Body)
}

func formatJson(data []byte, isList bool) (Commits, error) {
	commits := Commits{}
	if isList {
		commitsList := []Commits{}
		if err := json.Unmarshal(data, &commitsList); err != nil {
			return commits, err
		}
		commits = commitsList[0]
	} else {
		if err := json.Unmarshal(data, &commits); err != nil {
			return commits, err
		}
	}
	layout := "2006-01-02T15:04:05Z"
	date, err := time.Parse(layout, commits.Commit.Author.Date)
	if err != nil {
		return commits, err
	}
	commits.Date = date.Format("20060102150405")

	return commits, nil
}

type Commits struct {
	SHA    string `json:"sha"`
	Date   string
	Commit Commit `json:"commit"`
}

type Commit struct {
	Author Author `json:"author"`
}

type Author struct {
	Date string `json:"date"`
}
