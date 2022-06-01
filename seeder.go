package main

import (
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

func GetMnemonic(bitSize int) string {
	entropy, _ := bip39.NewEntropy(bitSize)
	mnemonic, _ := bip39.NewMnemonic(entropy)

	return mnemonic
}

func getMasterKey(mnemonic string) *bip32.Key {
	seed := bip39.NewSeed(mnemonic, "")
	masterKey, _ := bip32.NewMasterKey(seed)
	return masterKey
}

func getAddress(masterKey *bip32.Key) *bip32.Key {
	// m/44'/60'/0'/0/0
	child, _ := masterKey.NewChildKey(bip32.FirstHardenedChild + 44)
	child, _ = child.NewChildKey(bip32.FirstHardenedChild + 60)
	child, _ = child.NewChildKey(bip32.FirstHardenedChild + 0)
	child, _ = child.NewChildKey(0)
	child, _ = child.NewChildKey(0)
	return child
}
