package github_service

import (
	"encoding/json"
	"net/http"
)

type GitHubService struct {
	Username string
	Token    string
}

func NewGitHubService(username, token string) *GitHubService {
	return &GitHubService{username, token}
}

func (g *GitHubService) FetchRepositories() ([]RepositoryModel, error) {
	// construct request
	url := "https://api.github.com/user/repos"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(g.Username, g.Token)

	// send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// parse response
	var repositories []RepositoryModel
	if err := json.NewDecoder(resp.Body).Decode(&repositories); err != nil {
		return nil, err
	}

	// success
	return repositories, nil
}
