package mnemonicgenerator

import (
	"github.com/FactomProject/go-bip32"
	"github.com/FactomProject/go-bip39"
)

func GetMnemonic() string {
	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)

	return mnemonic
}

func GetMasterKey(mnemonic string) *bip32.Key {
	seed := bip39.NewSeed(mnemonic, "")
	masterKey, _ := bip32.NewMasterKey(seed)
	return masterKey
}

func GetAddress(masterKey *bip32.Key) *bip32.Key {
	child, _ := masterKey.NewChildKey(bip32.FirstHardenedChild + 44)
	child, _ = child.NewChildKey(bip32.FirstHardenedChild + 60)
	child, _ = child.NewChildKey(bip32.FirstHardenedChild + 0)
	child, _ = child.NewChildKey(0)
	child, _ = child.NewChildKey(0)
	return child
}
