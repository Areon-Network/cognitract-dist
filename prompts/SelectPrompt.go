package prompts

import "github.com/manifoldco/promptui"

func SelectPrompt[T interface{}](promptLabel string, items []T) string {
	prompt := promptui.Select{Label: promptLabel, Items: items}
	_, solcVersion, err := prompt.Run()

	if err != nil {
		panic(err)
	}

	return solcVersion
}
