package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
)

type Config struct {
	Username        string `json:"username"`
	Token           string `json:"token"`
	BackupDirectory string `json:"backup_directory"`
}

func ParseConfig() (*Config, error) {
	// get config file location
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	configFileLocation := path.Join(home, ".github-backup-rc.json")

	// open config file
	f, err := os.Open(configFileLocation)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// read file contents
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	// parse json
	var config Config
	if err := json.Unmarshal(buf, &config); err != nil {
		return nil, err
	}

	// success
	return &config, nil
}
