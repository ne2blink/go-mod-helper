package mod

import (
	"errors"
	"net/url"
)

// Resolve resolves a uri to go mod dependency format.
func Resolve(uri string) (string, error) {
	url, err := url.Parse(uri)
	if err != nil {
		return "", err
	}
	if url.Host == "github.com" {
		return resolveGitHub(url.Path, url.Fragment)
	}
	return "", errors.New("only GitHub repositories are supported")
}
