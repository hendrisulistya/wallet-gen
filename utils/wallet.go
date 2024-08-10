package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip39"
)

type Wallet struct {
	Address string `json:"address"`
	Path    string `json:"path"`
}

// SaveWallet saves the wallet information to a specified file.
func SaveWallet(wallet Wallet) {
	filePath := filepath.Join(".", wallet.Path)
	data, err := json.MarshalIndent(wallet, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal wallet data: %v", err)
	}

	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		log.Fatalf("Failed to save wallet: %v", err)
	}

	fmt.Printf("Wallet saved to %s\n", filePath)
}

// GenerateNewWallet generates a new wallet and saves its information.
func GenerateNewWallet() {
	mnemonic, _ := PromptAndGenerateMnemonic()

	// Print the generated mnemonic phrase
	fmt.Printf("Generated mnemonic: %s\n", mnemonic)

	// Generate keypair from mnemonic
	seed := bip39.NewSeed(mnemonic, "")
	privateKey, err := crypto.ToECDSA(seed[:32])
	if err != nil {
		log.Fatalf("Failed to generate private key: %v", err)
	}

	address := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()
	fmt.Printf("Generated address: %s\n", address)

	// Save wallet information
	wallet := Wallet{Address: address, Path: "wallet.json"}
	SaveWallet(wallet)
}

// RecoverWallet recovers a wallet from a mnemonic.
func RecoverWallet() {
	mnemonic := GetMnemonic()

	// Generate keypair from mnemonic
	seed := bip39.NewSeed(mnemonic, "")
	privateKey, err := crypto.ToECDSA(seed[:32])
	if err != nil {
		log.Fatalf("Failed to generate private key: %v", err)
	}

	address := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()
	fmt.Printf("Recovered address: %s\n", address)

	// Save wallet information
	wallet := Wallet{Address: address, Path: "wallet.json"}
	SaveWallet(wallet)
}