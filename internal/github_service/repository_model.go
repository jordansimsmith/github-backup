package github_service

type RepositoryModel struct {
	Id       uint64 `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	SSHUrl   string `json:"ssh_url"`
}
