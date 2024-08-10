package utils

import (
	"github.com/manifoldco/promptui"
)

// GetUserChoice prompts the user to select an option from a list.
func GetUserChoice(label string, items []string) (string, error) {
	prompt := promptui.Select{
		Label: label,
		Items: items,
	}

	_, result, err := prompt.Run()
	if err != nil {
		return "", err
	}
	return result, nil
}

// GetPromptInput prompts the user for input with a label.
func GetPromptInput(label string, defaultValue string) (string, error) {
	prompt := promptui.Prompt{
		Label:   label,
		Default: defaultValue,
	}

	result, err := prompt.Run()
	if err != nil {
		return "", err
	}
	return result, nil
}