package compilation

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/areon-network/cognitract/constants"
)

func checkBinaryPresence(filepath string) bool {
	if _, err := os.Stat(filepath); err == nil {
		fmt.Println("Compiler exists!")
		return true
	} else if os.IsNotExist(err) {
		fmt.Println("Compiler does not exist.")
		return false
	} else {
		fmt.Println("Error at checkBinaryPresence:", err)
		return false
	}
}

func DownloadCompiler(version string) {
	compilerDirectoryPrefix := constants.BinaryDirectory + "/soljson-"
	isCompilerAvailable := checkBinaryPresence(compilerDirectoryPrefix + version + ".bin")

	if isCompilerAvailable {
		return
	}

	fmt.Printf("📁 Downloading 'soljson-%s'\n", version)
	fmt.Printf("This might take a while, hold on...")

	url := fmt.Sprintf("https://binaries.soliditylang.org/bin/soljson-%s.js", version)
	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	ioutil.WriteFile(compilerDirectoryPrefix+version+".bin", body, 0644)
	defer response.Body.Close()
}
