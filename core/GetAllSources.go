package core

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

func GetAllSources(dir string) map[string][]byte {
	solFiles := make(map[string][]byte)
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		filePath := filepath.Join(dir, file.Name())

		if file.IsDir() {
			subFiles := GetAllSources(filePath)
			for k, v := range subFiles {
				if strings.HasPrefix(k, ".deps/npm/") {
					k = strings.TrimPrefix(k, ".deps/npm/")
				}

				if strings.HasPrefix(k, "node_modules/") {
					k = strings.TrimPrefix(k, "node_modules/")
				}

				solFiles[k] = v
			}
		} else if filepath.Ext(file.Name()) == ".sol" {
			content, err := ioutil.ReadFile(filePath)
			if err != nil {
				log.Printf("Failed to read file %s: %v", filePath, err)
				continue
			}

			solFiles[filePath] = content
		}

	}

	return solFiles
}
