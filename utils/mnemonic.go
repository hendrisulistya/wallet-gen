package utils

import (
	"log"
	"strings"

	"github.com/tyler-smith/go-bip39"
)

// PromptAndGenerateMnemonic prompts the user to choose mnemonic length and generates a mnemonic.
func PromptAndGenerateMnemonic() (string, []byte) {
	mnemonicLength, err := GetUserChoice("Choose mnemonic length", []string{"12 words", "24 words"})
	if err != nil {
		log.Fatalf("Prompt failed: %v", err)
	}

	var entropy []byte
	var bitSize int

	if mnemonicLength == "12 words" {
		bitSize = 128 // 128 bits for 12-word mnemonic
	} else if mnemonicLength == "24 words" {
		bitSize = 256 // 256 bits for 24-word mnemonic
	} else {
		log.Fatalf("Invalid choice")
	}

	entropy, err = bip39.NewEntropy(bitSize)
	if err != nil {
		log.Fatalf("Failed to generate entropy: %v", err)
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		log.Fatalf("Failed to generate mnemonic: %v", err)
	}

	return mnemonic, bip39.NewSeed(mnemonic, "")
}

// GetMnemonic prompts the user to enter a mnemonic and validates it.
func GetMnemonic() string {
	mnemonic, err := GetPromptInput("Enter your mnemonic", "")
	if err != nil {
		log.Fatalf("Prompt failed: %v", err)
	}
	mnemonic = strings.TrimSpace(mnemonic)

	if !bip39.IsMnemonicValid(mnemonic) {
		log.Fatalf("Invalid mnemonic")
	}
	return mnemonic
}

// GetHDPath prompts the user for an HD path, returning the default if none is provided.
func GetHDPath() string {
	hardPath, err := GetPromptInput("Enter HD path (default: m/44'/60'/0'/0/0)", "m/44'/60'/0'/0/0")
	if err != nil {
		log.Fatalf("Prompt failed: %v", err)
	}
	return hardPath
}