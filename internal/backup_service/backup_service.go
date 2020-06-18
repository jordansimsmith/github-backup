package backup_service

import (
	"os"
	"path"

	"github.com/jordansimsmith/github-backup/internal/model"
)

type BackupService struct {
	BackupDirectory string
}

func NewBackupService(backupDirectory string) (*BackupService, error) {
	// get default backup location
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	defaultFileLocation := path.Join(home, "Backups/github")

	backupService := &BackupService{defaultFileLocation}

	// override backup directory
	if len(backupDirectory) > 0 {
		backupService.BackupDirectory = backupDirectory
	}

	return backupService, nil
}

func (b *BackupService) BackupRepositories(repositories []model.RepositoryModel) error {
	// TODO: backup repositories

	// success
	return nil
}
