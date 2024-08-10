package main

import (
	"flag"
	"os"
	"strings"

	"github.com/hendrisulistya/wallet-gen/utils"
	"github.com/manifoldco/promptui"
)

func main() {
	// Display welcome ASCII art
	utils.WelcomeArt()

	// Check for command-line arguments
	interactiveMode := flag.Bool("it", false, "Enter interactive mode")
	flag.Parse()

	// If interactive mode flag is set, enter interactive mode
	if *interactiveMode {
		enterInteractiveMode()
	}

	// Handle command-line arguments for generate or recover
	args := flag.Args()
	if len(args) > 0 {
		command := strings.ToLower(args[0])
		switch command {
		case "generate":
			utils.GenerateNewWallet()
		case "recover":
			utils.RecoverWallet()
		default:
			println("Invalid command. Use 'generate' or 'recover'.")
			os.Exit(1)
		}
	} else {
		// No command provided, enter interactive mode
		enterInteractiveMode()
	}
}

func enterInteractiveMode() {
	// Create a list of options
	options := []string{"Generate new wallet", "Recover from mnemonic", "Exit"}

	// Create a new prompt for selecting an option
	prompt := promptui.Select{
		Label: "Choose an option",
		Items: options,
	}

	_, result, err := prompt.Run()
	if err != nil {
		os.Exit(1)
	}

	switch result {
	case "Generate new wallet":
		utils.GenerateNewWallet()
	case "Recover from mnemonic":
		utils.RecoverWallet()
	case "Exit":
		os.Exit(0)
	}
}