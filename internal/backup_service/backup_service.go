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

func (b *BackupService) BackupRepositories(repositoryModels []model.RepositoryModel) error {
	// create backup directory
	if _, err := os.Stat(b.BackupDirectory); os.IsNotExist(err) {
		os.MkdirAll(b.BackupDirectory, os.ModePerm)
	}
	os.Chdir(b.BackupDirectory)

	// map models to Repository structs
	repositories := make([]Repository, 0)
	for _, repositoryModel := range repositoryModels {
		repository := NewRepository(repositoryModel.Id, repositoryModel.Name, repositoryModel.SSHUrl)
		repositories = append(repositories, *repository)
	}

	// backup each repository concurrently
	results := make(chan error)
	for _, repository := range repositories {

		go func(r Repository) {
			// attempt backup
			results <- r.Backup(b.BackupDirectory)
		}(repository)
	}

	// wait for goroutines
	for range repositories {
		if err := <-results; err != nil {
			return err
		}
	}
	close(results)

	// success
	return nil
}
