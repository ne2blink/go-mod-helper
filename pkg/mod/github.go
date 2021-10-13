package mod

import (
	"fmt"
	"strings"

	"github.com/ne2blink/go-mod-helper/pkg/github"
)

func resolveGitHub(path, branch string) (string, error) {
	repo := strings.TrimSuffix(path, ".git")
	repo = strings.TrimPrefix(repo, "/")

	commit, err := github.GetHeadCommit(repo, branch)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("github.com/%s v0.0.0-%s-%s",
		repo,
		commit.Commit.Committer.Date.UTC().Format("20060102150405"),
		commit.SHA[:12],
	), nil
}
