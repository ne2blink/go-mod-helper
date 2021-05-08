package helper

type GithubCommit struct {
	SHA    string `json:"sha"`
	Commit Commit `json:"commit"`
}

type Commit struct {
	Author Author `json:"author"`
}

type Author struct {
	Date string `json:"date"`
}
