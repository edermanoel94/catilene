package main

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

const (
	parentFolder = "Parent"
)

var (
	extensionCounterStats = make(map[string]int)
	extensionsMap         = make(map[string][]string)
)

// call api virus total

func main() {
	// passing a file in args

	filename := os.Args[1]

	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	fileStat, err := file.Stat()

	if err != nil {
		log.Fatal(err)
	}

	if fileStat.IsDir() {
		filepath.Walk(filename, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				ext := extensionCleaned(path)
				extensionsMap[ext] = append(extensionsMap[ext], path)
				extensionCounterStats[ext] += 1
			}
			return nil
		})
	} else {
		ext := filepath.Ext(file.Name())

		if ext == "zip" {

		}
	}

	batchTransfer(extensionsMap)

	report(extensionCounterStats)
}
