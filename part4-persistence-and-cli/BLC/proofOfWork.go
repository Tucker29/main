package BLC

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

var (
	maxNonce = math.MaxInt64
)

const targetBit = 12

type ProofOfWork struct {
	block  *Block
	target *big.Int //区块难度
}

func (pow *ProofOfWork) PrepareData(Nonce int) []byte {
	Data := bytes.Join([][]byte{
		IntToHex(int64(pow.block.Timestamp)),
		pow.block.PreHash,
		pow.block.Data,
		IntToHex(int64(targetBit)),
		IntToHex(int64(pow.block.Nonce)),
	}, []byte{})
	return Data
}

func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, 256-targetBit)
	pow := &ProofOfWork{block, target}
	return pow
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var HashInt big.Int
	var hash [32]byte
	nonce := 0
	for nonce < maxNonce {
		hash := sha256.Sum256(pow.PrepareData(nonce))
		fmt.Printf("\r%x", hash)
		HashInt.SetBytes(hash[:])
		if HashInt.Cmp(pow.target) == -1 {
			break
		}

		nonce++
	}
	fmt.Printf("\n")
	return nonce, hash[:]
}
