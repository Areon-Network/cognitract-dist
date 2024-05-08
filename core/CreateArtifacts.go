package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/areon-network/cognitract/gsolc"
)

func CreateArtifacts(out *gsolc.Output) {
	fmt.Println("- Creating artifacts in your project directory...")
	artifactData, err := json.Marshal(out)
	if err != nil {
		fmt.Println(err)
	}

	if _, err := os.Stat("./.cognitract"); os.IsNotExist(err) {
		os.Mkdir("./.cognitract", os.ModeDir|0755)
	}

	errWrite := ioutil.WriteFile("./.cognitract/build-info.json", artifactData, 0644)
	if errWrite != nil {
		fmt.Println(errWrite)
	}

	fmt.Println("- Artifact files are created successfully.")
}

func CreateAbiFolder(abi []byte) {
	fmt.Println("- Creating main contract ABI in your project directory...")
	if _, err := os.Stat("./.cognitract"); os.IsNotExist(err) {
		os.Mkdir("./.cognitract", os.ModeDir|0755)
	}

	errWrite := ioutil.WriteFile("./.cognitract/abi.json", abi, 0644)
	if errWrite != nil {
		fmt.Println(errWrite)
	}

	fmt.Println("- ABI created successfully.")
}
