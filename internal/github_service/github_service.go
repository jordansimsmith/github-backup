package github_service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jordansimsmith/github-backup/internal/model"
)

type GitHubService struct {
	Username string
	Token    string
}

func NewGitHubService(username, token string) *GitHubService {
	return &GitHubService{username, token}
}

func (g *GitHubService) FetchRepositories() ([]model.RepositoryModel, error) {
	repositories := make([]model.RepositoryModel, 0)

	// traverse paginated resources
	hasNext := true
	page := 1
	for hasNext {
		paginatedRepositories, err := g.fetchRepositoriesPage(page)
		if err != nil {
			return nil, err
		}

		// no more results
		if len(paginatedRepositories) == 0 {
			hasNext = false
		}

		repositories = append(repositories, paginatedRepositories...)
		page++
	}

	// success
	return repositories, nil
}

func (g *GitHubService) fetchRepositoriesPage(page int) ([]model.RepositoryModel, error) {
	// construct request
	url := "https://api.github.com/user/repos"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(g.Username, g.Token)
	query := req.URL.Query()
	query.Add("type", "owner")
	query.Add("page", fmt.Sprint(page))
	req.URL.RawQuery = query.Encode()

	// send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// parse response
	var repositories []model.RepositoryModel
	if err := json.NewDecoder(resp.Body).Decode(&repositories); err != nil {
		return nil, err
	}

	// success
	return repositories, nil
}
