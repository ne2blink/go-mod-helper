package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type Commit struct {
	SHA    string `json:"sha"`
	Commit struct {
		Author struct {
			Date time.Time `json:"date"`
		} `json:"author"`
	} `json:"commit"`
}

func GetHeadCommit(repo, branch string) (*Commit, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/commits", repo)
	if branch != "" {
		url += "/" + branch
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	decoder := json.NewDecoder(resp.Body)

	if branch == "" {
		var result []Commit
		if err := decoder.Decode(&result); err != nil {
			return nil, err
		}
		return &result[0], nil
	} else {
		var result Commit
		if err := decoder.Decode(&result); err != nil {
			return nil, err
		}
		return &result, nil
	}
}
