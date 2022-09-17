package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"
)

type Chain struct {
	Chain []Block `json:"Chain"`
}

func (chain *Chain) CreateBlock(proof uint32, previousHash string) Block {
	block := Block{
		Index:        len(chain.Chain) + 1,
		TimeStamp:    fmt.Sprintf("%d", time.Now().UnixNano()),
		Proof:        proof,
		PreviousHash: previousHash,
	}

	chain.Chain = append(chain.Chain, block)
	return block
}

func (chain *Chain) GetPreviousBlock() Block {
	return chain.Chain[len(chain.Chain)-1]
}

func (chain *Chain) IsChainValid() bool {
	blockIndex := 1
	previousBlock := chain.Chain[0]
	for blockIndex < len(chain.Chain) {
		block := chain.Chain[blockIndex]
		if block.PreviousHash != previousBlock.Hash() {
			return false
		}

		hashOperation := sha256.New()
		hashComplexity := rune(int64(math.Pow(float64(block.Proof), 2) - math.Pow(float64(previousBlock.Proof), 2)))
		_, err := hashOperation.Write([]byte(string(hashComplexity)))
		if err != nil {
			panic(err)
		}

		hexHashSum := hex.EncodeToString(hashOperation.Sum(nil))
		if hexHashSum[:4] != "0000" {
			return false
		}

		previousBlock = block
		blockIndex++
	}

	return true
}

func (chain *Chain) HttpMineBlock(w http.ResponseWriter, r *http.Request) {
	previousBlock := chain.GetPreviousBlock()
	previousProof := previousBlock.Proof
	proof := previousBlock.PoW(previousProof)
	previousHash := previousBlock.Hash()

	block := chain.CreateBlock(proof, previousHash)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(block)
}

func (chain *Chain) HttpGetChain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(chain.Chain)
}

func (chain *Chain) HttpIsChainValid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	w.Write([]byte(strconv.FormatBool(chain.IsChainValid())))
}
