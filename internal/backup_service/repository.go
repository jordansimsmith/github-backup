package backup_service

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

type Repository struct {
	Id     uint64
	Name   string
	SSHUrl string
}

func NewRepository(id uint64, name, sshUrl string) *Repository {
	return &Repository{id, name, sshUrl}
}

func (r *Repository) Backup(backupDirectory string) error {

	id := fmt.Sprint(r.Id)

	// check if the repository has already been cloned
	repositoryDirectory := path.Join(backupDirectory, id)
	if _, err := os.Stat(repositoryDirectory); os.IsNotExist(err) {
		// clone new repository
		cmd := exec.Command("git", "clone", r.SSHUrl, id)
		if err := cmd.Run(); err != nil {
			return err
		}
		fmt.Printf("cloned %s into %d\n", r.SSHUrl, r.Id)
	} else {
		// pull existing repository
		cmd := exec.Command("git", "pull")
		cmd.Dir = id
		if err := cmd.Run(); err != nil {
			return err
		}
		fmt.Printf("pulled %s in %d\n", r.SSHUrl, r.Id)
	}

	// success
	return nil
}
