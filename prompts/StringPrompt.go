package prompts

import (
	"errors"

	"github.com/manifoldco/promptui"
)

func StringPrompt(promptLabel string) string {
	pathPrompt := promptui.Prompt{Label: promptLabel, Validate: func(input string) error {
		if len(input) < 4 {
			return errors.New("Invalid value")
		}
		return nil
	}}

	result, err := pathPrompt.Run()
	if err != nil {
		panic(err)
	}

	return result
}
