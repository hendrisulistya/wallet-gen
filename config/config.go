package config

// DefaultWalletPath is the default file path for saving the wallet.
const DefaultWalletPath = "wallet.json"

// DefaultHDPath is the default HD path for wallet generation.
const DefaultHDPath = "m/44'/60'/0'/0/0"

// MnemonicLengths holds the available mnemonic lengths for selection.
var MnemonicLengths = map[int]string{
	1: "12 words",
	2: "24 words",
}