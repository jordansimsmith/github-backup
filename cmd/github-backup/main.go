package main

import (
	"fmt"
	"log"

	"github.com/jordansimsmith/github-backup/internal/config"
)

func main() {
	fmt.Println("github-backup is starting")

	// read in config
	config, err := config.ParseConfig()
	if err != nil {
		log.Fatal("Could not parse config file: ", err)
	}
	fmt.Println("parsed config file")

	fmt.Println(config)
}
