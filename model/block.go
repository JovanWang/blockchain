package model

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const difficulty = 1

type Block struct {
	Index         int    // 区块编号
	Timestamp     string // 时间戳
	PrevBlockHash string // 上一个区块的hash值
	Hash          string // 当前区块的hash

	Data string // 区块数据

	Difficulty int    // 难度
	Nonce      string // 凭证
}

func isHashValid(hash string, difficulty int) bool {
	prefix := strings.Repeat("0", difficulty)
	return strings.HasPrefix(hash, prefix)
}

// 计算hash值
func calculateHash(b Block) string {
	blockData := strconv.Itoa(b.Index) + b.Timestamp + b.PrevBlockHash + b.Data + b.Nonce
	hashInBytes := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(hashInBytes[:])
}

// 产生新的区块
func GenerateNewBlock(prevBlock Block, data string) Block {
	NewBlock := Block{}
	NewBlock.Index = prevBlock.Index + 1
	NewBlock.Timestamp = time.Now().String()
	NewBlock.PrevBlockHash = prevBlock.Hash
	NewBlock.Data = data
	NewBlock.Difficulty = difficulty
	for i := 0; ; i++ {
		hex := fmt.Sprintf("%x", i)
		NewBlock.Nonce = hex
		letBlockHash := calculateHash(NewBlock)
		if !isHashValid(letBlockHash, NewBlock.Difficulty) {
			fmt.Println(letBlockHash, " do more work!!")
			continue
		} else {
			fmt.Println(letBlockHash, " work done!!")
			NewBlock.Hash = letBlockHash
			break
		}
	}
	return NewBlock
}

// 产生创世区块
func GenerateGenesisBlock() Block {
	prevBlock := Block{}
	prevBlock.Index = -1
	prevBlock.PrevBlockHash = ""
	return GenerateNewBlock(prevBlock, "Genesis Block")
}
