package main

import (
	components "alt-coin/blockchain/components"
	"log"
	"net/http"
)

func main() {
	blockchain := components.Chain{
		Chain: []components.Block{},
	}
	blockchain.CreateBlock(1, "0") // Genesis block

	http.HandleFunc("/mine", blockchain.HttpMineBlock)
	http.HandleFunc("/chain", blockchain.HttpGetChain)
	http.HandleFunc("/isvalid", blockchain.HttpIsChainValid)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
