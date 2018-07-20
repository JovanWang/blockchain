package model

import (
	"fmt"
	"log"
)

type Blockchain struct {
	Blocks []*Block
}

func NewBlockchain() *Blockchain {
	genesisBlock := GenerateGenesisBlock()
	blockchain := Blockchain{}
	blockchain.AppendBlock(genesisBlock)
	return &blockchain
}

func (bc *Blockchain) SendData(data string) {
	preBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := GenerateNewBlock(*preBlock, data)
	bc.AppendBlock(newBlock)
}

func (bc *Blockchain) replaceChain(newBlocks []*Block) {
	if len(newBlocks) > len(bc.Blocks) {
		bc.Blocks = newBlocks
	}
}

func (bc *Blockchain) AppendBlock(newBlock Block) {
	if len(bc.Blocks) == 0 {
		bc.Blocks = append(bc.Blocks, &newBlock)
		return
	}
	if isValid(newBlock, *bc.Blocks[len(bc.Blocks)-1]) {
		newBlocks := append(bc.Blocks, &newBlock)
		bc.replaceChain(newBlocks)
	} else {
		log.Fatal("invalid block!")
	}

}

func isValid(newBlock Block, oldBlock Block) bool {
	if newBlock.Index != oldBlock.Index+1 {
		return false
	}
	if newBlock.PrevBlockHash != oldBlock.Hash {
		return false
	}
	if newBlock.Hash != calculateHash(newBlock) {
		return false
	}
	return true
}

func (bc *Blockchain) Print() {
	for _, block := range bc.Blocks {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Printf("PrevBlockHash: %s\n", block.PrevBlockHash)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("Data: %s\n\n", block.Data)
	}
}
