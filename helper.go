package helper

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	Version = "v0.0.0-"
	Layout  = "2006-01-02T15:04:05Z"
)

type Helper struct {
	url  string
	SHA  string
	Date time.Time
}

// New a Helper
func New(url string) *Helper {
	return &Helper{
		url: url,
	}
}

// Get
func (h *Helper) Get() error {
	if strings.HasPrefix(h.url, "github.com/") || strings.HasPrefix(h.url, "http://github.com/") || strings.HasPrefix(h.url, "https://github.com/") {
		return h.github()
	} else {
		return errors.New("not support url")
	}
}

func (h Helper) Print() {
	version := h.clearURL() + " " + Version + h.Date.Format("20060102150405") + "-" + h.SHA[:12]
	fmt.Println(version)
}

func (h Helper) clearURL() string {
	h.url = strings.TrimPrefix(h.url, "http://")
	h.url = strings.TrimPrefix(h.url, "https://")
	return strings.Split(h.url, "#")[0]
}
