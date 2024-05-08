package core

var ActiveNetwork string = "mainnet"

func SetActiveNetwork(newValue string) {
	ActiveNetwork = newValue
}

func GetActiveNetwork() string {
	return ActiveNetwork
}
