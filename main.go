package main

import (
	"encoding/hex"
	"fmt"
	"vanitygo/mnemonicgenerator"
)

func main() {
	// mnemonic := mnemonicgenerator.GetMnemonic()
	mnemonic := "gather absent remind casual spray buddy raise faith useless load require lucky truly venture chaos angle immune lava bubble rose trophy large miracle churn"
	fmt.Println(mnemonic)
	pubkey := mnemonicgenerator.GetMasterKey(mnemonic)
	// fmt.Println(pubkey)
	address := mnemonicgenerator.GetAddress(pubkey)
	fmt.Println(address.PublicKey())
	ether := hex.EncodeToString(address.Serialize())
	fmt.Println(ether)
}
