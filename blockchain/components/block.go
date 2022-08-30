package components

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
	"math"
)

type Block struct {
	Index        int    `json:"Index"`
	TimeStamp    string `json:"TimeStamp"`
	Proof        int64  `json:"Proof"`
	PreviousHash string `json:"PreviousHash"`
}

// Proof of work
func (block *Block) PoW(previousProof int64) int64 {
	var newProof int64 = 1
	checkProof := false

	for !checkProof {
		hashOperation := sha256.New()
		hashComplexity := rune(int64(math.Pow(float64(newProof), 2) - math.Pow(float64(previousProof), 2)))
		_, err := hashOperation.Write([]byte(string(hashComplexity)))
		if err != nil {
			panic(err)
		}

		hexHashSum := hex.EncodeToString(hashOperation.Sum(nil))
		if hexHashSum[:4] == "0000" {
			log.Println("Block mined!")
			checkProof = true
		} else {
			newProof++
		}

		log.Println(newProof)
	}

	return newProof
}

func (block *Block) Hash() string {
	jsonBlock, jsonErr := json.Marshal(block)
	if jsonErr != nil {
		panic(jsonErr)
	}

	hash := sha256.New()
	_, err := hash.Write([]byte(string(jsonBlock)))
	if err != nil {
		panic(err)
	}

	hexHashSum := hex.EncodeToString(hash.Sum(nil))
	return hexHashSum
}
