// Functions and variables are not exported
// on purpose as this is a sample

package core

import "fmt"

var areonScanEndpoint string

func initializeAreonScanEndpoint() {
	AreonScanEndpoint = fmt.Sprintf("", GetActiveNetwork())
}

func getAreonScanEndpoint() string {
	return AreonScanEndpoint
}
