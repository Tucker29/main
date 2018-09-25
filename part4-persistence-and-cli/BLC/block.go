package BLC

import (
	"time"
)

type Block struct {
	Timestamp int64
	PreHash   []byte
	Data      []byte
	Hash      []byte
	Nonce     int
}

func NewBlock(data string, PreHash []byte) *Block {
	block := &Block{time.Now().Unix(), PreHash, []byte(data), []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Nonce = nonce
	block.Hash = hash
	return block
}

// 创建创世区块，并返回创世区块
func NewGenesisBlock() *Block {
	return NewBlock("Genenis Block", []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}
