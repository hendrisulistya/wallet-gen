package model

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Wallet represents an Ethereum wallet.
type Wallet struct {
	Address string `json:"address"`
	Path    string `json:"path"`
}

// Save saves the wallet information to a specified file.
func (w *Wallet) Save() error {
	filePath := filepath.Join(".", w.Path)
	data, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal wallet data: %v", err)
	}

	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to save wallet: %v", err)
	}

	fmt.Printf("Wallet saved to %s\n", filePath)
	return nil
}