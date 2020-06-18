package backup_service

import "fmt"

type Repository struct {
	SSHUrl string
}

func NewRepository(sshUrl string) *Repository {
	return &Repository{sshUrl}
}

func (r *Repository) Backup() error {
	// TODO: implement backup
	fmt.Println("backing up ", r.SSHUrl)

	// success
	return nil
}
