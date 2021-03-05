package flaver

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

const glReleasesRequest = "https://gitlab.com/api/v4/projects/%s/releases"

type glRelease struct {
	TagName     string     `json:"tag_name"`
	PublishDate *time.Time `json:"released_at"`
	Links       struct {
		Self string `json:"self"`
	} `json:"_links"`
}

func (glr *glRelease) Version() string {
	return glr.TagName
}

func (glr *glRelease) Date() *time.Time {
	return glr.PublishDate
}

func (glr *glRelease) Url() string {
	return glr.Links.Self
}

type GitLabFlaver struct {
}

func (glf *GitLabFlaver) Find(repo *Repository) (Release, error) {
	releases, err := glf.FindAll(repo)
	if err == nil {
		return releases[0], err
	}
	return nil, err
}

func (glf *GitLabFlaver) FindAll(repo *Repository) ([]Release, error) {
	client := resty.New()
	request := client.NewRequest()
	url := fmt.Sprintf(glReleasesRequest, strings.ReplaceAll(repo.name, "/", "%2F"))

	releases := []*glRelease{}
	response, err := request.SetResult(&releases).Get(url)
	if err != nil {
		return nil, err
	}
	return glParseResponse(response, repo.name, releases)
}

func glParseResponse(response *resty.Response, name string, releases []*glRelease) ([]Release, error) {
	if response.StatusCode() == 404 {
		return nil, fmt.Errorf("%s: repository not found", name)
	}
	result := []Release{}
	for _, item := range releases {
		result = append(result, item)
	}
	return result, nil
}
