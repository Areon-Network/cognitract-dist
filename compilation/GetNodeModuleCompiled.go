package compilation

import (
	"io/ioutil"
	"log"
)

// This compiles node.js modules that are inside
// the argument array, not the whole module directory
func GetNodeModuleCompiled(dir []string) map[string][]byte {
	solFiles := make(map[string][]byte)
	for _, i := range dir {
		content, err := ioutil.ReadFile("node_modules/" + i)
		if err != nil {
			log.Printf("Failed to read file")
		}

		solFiles[i] = content
	}

	return solFiles

}
