package compilation

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

func GetSourceCode(dir string) map[string][]byte {
	solFiles := make(map[string][]byte)
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		filePath := filepath.Join(dir, file.Name())
		if strings.HasPrefix(filePath, "node_modules") {
			continue
		}

		if file.IsDir() {
			subFiles := GetSourceCode(filePath)
			for k, v := range subFiles {
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
