package main

import (
	"fmt"
	"log"

	"github.com/jordansimsmith/github-backup/internal/backup_service"
	"github.com/jordansimsmith/github-backup/internal/config"
	"github.com/jordansimsmith/github-backup/internal/github_service"
)

func main() {
	fmt.Println("github-backup is starting")

	// read in config
	config, err := config.ParseConfig()
	if err != nil {
		log.Fatal("could not parse config file: ", err)
	}
	fmt.Println("parsed config file")
	fmt.Println(config)

	// read repositories
	gitHubService := github_service.NewGitHubService(config.Username, config.Token)
	repositories, err := gitHubService.FetchRepositories()
	if err != nil {
		log.Fatal("could not fetch repositories", err)
	}
	fmt.Println("fetched repositories")
	fmt.Println(repositories)

	// backup repositories
	backupService, err := backup_service.NewBackupService(config.BackupDirectory)
	if err != nil {
		log.Fatal("could not initialise backup service", err)
	}
	if err := backupService.BackupRepositories(repositories); err != nil {
		log.Fatal("could not backup one or more repositories", err)
	}

	// success
	fmt.Println("github backup successful")
}
