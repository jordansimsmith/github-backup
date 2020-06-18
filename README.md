# GitHub Backup

Simple tool that maintains backups of your GitHub repositories. Repositories are identified using the GitHub id, to account for changes to the repository name.

## Configuration

Configuration for GitHub Backup is managed using `~/.github-backup-rc.json`. Please configure the fields as instructed.

```json
{
  "username": "GITHUB_USERNAME",
  "token": "GITHUB_PERSONAL_ACCESS_TOKEN",
  "backup_directory": "OPTIONAL_BACKUP_DIRECTORY"
}
```

- The Personal Access Token (PAT) must have full repository priviledges if private repositories are to be backed up.
- GitHub Backup manages git repositories using SSH, please ensure SSH is configured appropriately.
- The `backup_directory` field is optional. The default location is `~/Backups/github`

## Build Instructions
### From Source

1. Clone the repository

```
git clone git@github.com:jordansimsmith/github-backup.git
```

2. Change directory into the repository

```
cd github-backup.git
```

3. Compile the binary

```
go build ./cmd/github-backup
```

### From Go

```
go get github.com/jordansimsmith/github-backup/cmd/github-backup
```

## Usage Instructions

```
github-backup
```
