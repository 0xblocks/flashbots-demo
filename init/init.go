package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	signingKey, err := crypto.GenerateKey()
	signingKeyBytes := crypto.FromECDSA(signingKey)

	fmt.Printf("BOT_PRIVATE_KEY=%s\n", hexutil.Encode(privateKeyBytes)[2:])
	fmt.Printf("BOT_ADDRESS=%s\n", fromAddress)
	fmt.Printf("FLASHBOTS_SIGNING_KEY=%s\n", hexutil.Encode(signingKeyBytes)[2:])
	fmt.Println("PROVIDER_URL=https://goerli.infura.io/v3/[yourkey]")
}
