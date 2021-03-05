package flaver

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
)

type Repository struct {
	name string
	url  string
}

type Release interface {
	Version() string
	Date() *time.Time
	Url() string
}

type Flaver interface {
	Find(repo *Repository) (Release, error)
	FindAll(repo *Repository) ([]Release, error)
}

func IsValidRepositoryName(name string) bool {
	values := strings.Split(name, "/")
	return len(values) == 2
}

func IsValidUrl(urlString string) bool {
	if urlObject, err := url.Parse(urlString); err == nil {
		return urlObject.Host != "" && urlObject.Scheme != "" && govalidator.IsURL(urlString)
	}
	return false
}

func NewFlaver(repo *Repository) (Flaver, error) {
	switch {
	case strings.HasPrefix(repo.url, "https://github.com/"):
		return &GitHubFlaver{}, nil
	case strings.HasPrefix(repo.url, "https://gitlab.com/"):
		return &GitLabFlaver{}, nil
	}
	return nil, fmt.Errorf("%s: suitable flaver not found", repo.url)
}

func BuildRepository(name string) (*Repository, error) {
	switch {
	case strings.HasPrefix(name, "https://github.com/"):
		return NewRepository(strings.TrimPrefix(name, "https://github.com/"), name)
	case strings.HasPrefix(name, "github.com/"):
		repoName := strings.TrimPrefix(name, "github.com/")
		return NewRepository(repoName, fmt.Sprintf("https://%s", name))
	case strings.HasPrefix(name, "gitlab.com/"):
		repoName := strings.TrimPrefix(name, "gitlab.com/")
		return NewRepository(repoName, fmt.Sprintf("https://%s", name))
	default:
		return NewGitHubRepository(name)
	}
}

func NewRepository(name, url string) (*Repository, error) {
	if !IsValidUrl(url) {
		return nil, fmt.Errorf("%s: invalid url", url)
	}
	return &Repository{name: name, url: url}, nil
}

func NewGitHubRepository(name string) (*Repository, error) {
	if !IsValidRepositoryName(name) {
		return nil, fmt.Errorf("%s: invalid github repository name", name)
	}
	return NewRepository(name, fmt.Sprintf("https://github.com/%s", name))
}
