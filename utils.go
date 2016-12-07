package main

import (
	"log"
	"os"
)

// safeOpen just exits if it fails to open the file.
func safeOpen(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Unable to open %s: %s", path, err)
	}
	return file
}
