package flaver

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

const ghReleasesRequest = "https://api.github.com/repos/%s/releases"

type ghRelease struct {
	TagName     string     `json:"tag_name"`
	PublishDate *time.Time `json:"published_at"`
	ReleaseUrl  string     `json:"html_url"`
}

func (ghr *ghRelease) Version() string {
	return ghr.TagName
}

func (ghr *ghRelease) Date() *time.Time {
	return ghr.PublishDate
}

func (ghr *ghRelease) Url() string {
	return ghr.ReleaseUrl
}

type GitHubFlaver struct {
}

func (ghf *GitHubFlaver) Find(repo *Repository) (Release, error) {
	releases, err := ghf.FindAll(repo)
	if err == nil {
		return releases[0], err
	}
	return nil, err
}

func (ghf *GitHubFlaver) FindAll(repo *Repository) ([]Release, error) {
	client := resty.New()
	request := client.NewRequest()
	url := fmt.Sprintf(ghReleasesRequest, repo.name)

	releases := []*ghRelease{}
	response, err := request.SetResult(&releases).Get(url)
	if err != nil {
		return nil, err
	}
	return ghParseResponse(response, repo.name, releases)
}

func ghParseResponse(response *resty.Response, name string, releases []*ghRelease) ([]Release, error) {
	if response.StatusCode() == 404 {
		return nil, fmt.Errorf("%s: repository not found", name)
	}
	result := []Release{}
	for _, item := range releases {
		result = append(result, item)
	}
	return result, nil
}
