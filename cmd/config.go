package cmd

import (
	"errors"
	"log"
	"os"
	"path/filepath"
)

var (
	ereConfigFolderName = ".ere"
	ereCheckedFileName  = "checked.txt"
	ereEventsFileName   = "events.json"
)

// Returns the ere config folder path
// Creates the folder if it doesn't exist
func ereConfigFolder() string {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	folder := filepath.Join(homedir, ereConfigFolderName)
	f, err := os.Stat(folder)
	if errors.Is(err, os.ErrNotExist) {
		// Create it it doesn't exist
		err := os.MkdirAll(folder, 0o755)
		if err != nil {
			log.Fatal(err)
		}
	} else if err != nil {
		log.Fatal(err)
	}
	if err == nil && !f.IsDir() {
		log.Fatalf("expecting a directory %v got a file", folder)
	}
	return folder
}
