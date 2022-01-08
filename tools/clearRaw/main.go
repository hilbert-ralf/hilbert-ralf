package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// build for windows: env GOOS=windows GOARCH=amd64 go build
// see for details: https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04
func main() {

	rawPathFileInfoMap := make(map[string]os.FileInfo)
	var jpegs []string

	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			//fmt.Println(path)

			if strings.HasSuffix(path, ".ARW") {
				rawPathFileInfoMap[path] = info
			}

			if strings.HasSuffix(path, ".JPG") {
				jpgWithoutEnding := strings.TrimSuffix(info.Name(), ".JPG")
				jpegs = append(jpegs, jpgWithoutEnding)
			}

			return nil
		})

	if err != nil {
		log.Println(err)
	}

	var pathsToDelete []string
	for path, info := range rawPathFileInfoMap {
		//fmt.Println(path, info.Name())
		rawWithoutEnding := strings.TrimSuffix(info.Name(), ".ARW")
		if contains(jpegs, rawWithoutEnding) {
			pathsToDelete = append(pathsToDelete, path)
		}
	}

	fmt.Println("Sollen folgende Dateien gelöscht werden?")
	fmt.Println(pathsToDelete)

	fmt.Println("Bitte mit 'y' oder 'n' antworten:")

	var response string
	_, err2 := fmt.Scanln(&response)
	if err2 != nil {
		log.Fatal(err2)
	}

	if response == "y" {
		fmt.Println("Dateien werden gelöscht!")

		for _, rawToDelete := range pathsToDelete {
			fmt.Println("delete", rawToDelete)
			os.Remove(rawToDelete)
		}

	} else {
		fmt.Println("Nichts wurde gelöscht!")
	}
}

func contains(given []string, searched string) bool {
	for _, a := range given {
		if a == searched {
			return true
		}
	}
	return false
}
