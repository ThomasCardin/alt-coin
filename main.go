package main

import (
	blockchain "alt-coin/pkg/blockchain"
	"log"
	"net/http"
)

func main() {
	blockchain := blockchain.Chain{
		Chain: []blockchain.Block{},
	}
	blockchain.CreateBlock(1, "0") // Genesis block

	http.HandleFunc("/mine", blockchain.HttpMineBlock)
	http.HandleFunc("/chain", blockchain.HttpGetChain)
	http.HandleFunc("/isvalid", blockchain.HttpIsChainValid)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
