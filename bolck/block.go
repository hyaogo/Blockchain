package bolck

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp    int64
	Data         []byte
	PreBlockHash []byte
	Hash         []byte
}

type Blockchain struct {
	Blocks []*Block
}

func (block *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))
	headers := bytes.Join([][]byte{block.PreBlockHash, block.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	block.Hash = hash[:]
}

func NewBlock(data string, preBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), preBlockHash, []byte{}}
	block.SetHash()
	return block
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func (bc *Blockchain) AddBlock(data string) {
	preBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, preBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func NewBlockChain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
