package mod

import (
	"fmt"
	"strings"

	"github.com/ne2blink/go-mod-helper/pkg/github"
)

func resolveGitHub(path string) (string, error) {
	parts := strings.SplitN(path, "#", 2)
	repo := strings.TrimSuffix(parts[0], ".git")
	repo = strings.TrimPrefix(repo, "/")
	branch := ""
	if len(parts) == 2 {
		branch = parts[1]
	}

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
