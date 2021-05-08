package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	BaseURL = "https://api.github.com/repos/"
	Commits = "/commits"
)

func (h *Helper) github() error {
	url, err := h.githubApiURL()
	fmt.Println(url)
	if err != nil {
		return err
	}
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = h.githubJsonDecode(body)
	if err != nil {
		return err
	}
	return nil
}

func (h Helper) githubApiURL() (string, error) {
	h.url = strings.TrimPrefix(h.url, "github.com/")
	h.url = strings.TrimPrefix(h.url, "http://github.com/repos")
	h.url = strings.TrimPrefix(h.url, "https://github.com/")

	branch := Commits
	branchSlice := strings.Split(h.url, "#")
	if len(branchSlice) > 2 {
		return "", errors.New("url wrong format")
	}
	if len(branchSlice) == 2 {
		branch = branch + "/" + branchSlice[1]
	}
	return BaseURL + branchSlice[0] + branch, nil
}

func (h *Helper) githubJsonDecode(data []byte) error {
	var dateString string
	if len(strings.Split(h.url, "#")) == 2 {
		var commit GithubCommit
		if err := json.Unmarshal(data, &commit); err != nil {
			return err
		}
		h.SHA = commit.SHA
		dateString = commit.Commit.Author.Date
	} else {
		var commits []GithubCommit
		if err := json.Unmarshal(data, &commits); err != nil {
			return err
		}
		h.SHA = commits[0].SHA
		dateString = commits[0].Commit.Author.Date
	}
	date, err := time.Parse(Layout, dateString)
	if err != nil {
		return err
	}
	h.Date = date
	return nil
}
