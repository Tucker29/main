package BLC

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

var (
	// 定义Nonce最大值
	maxNonce = math.MaxInt64
)

const targetBits = 12

// 0000000000000000000001000000000000000000000000000000000000000000000000000000000000000001
// 00000001
// 00100000
// 00010000
// 8 - 3 = 5

type ProofOfWork struct {
	block  *Block   // 当前需要验证的区块
	target *big.Int // 大数存储,区块难度
}

// 数据拼接，返回字节数组
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

// ProofOfWork对象的方法
func (pow *ProofOfWork) Run() (int, []byte) {

	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)

	for nonce < maxNonce {

		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])
		// pow.target hash
		// hashInt < pow.target -1
		// hashInt = pow.target 0
		// hashInt > pow.target 1
		// 如果hashInt < pow.target ，break
		// 000010000000
		// 000001000000
		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}

	fmt.Println("\n\n")

	return nonce, hash[:]
}

// 工厂方法
func NewProofOfWork(block *Block) *ProofOfWork {
	// 0000000000000000000100000000000000
	// 00fadadasdasdsfdsf0080808080582452
	// 000000001111
	target := big.NewInt(1)
	//fmt.Println("----------")
	//fmt.Println(target)
	target.Lsh(target, uint(256-targetBits))
	// 0000
	//fmt.Println(target)

	pow := &ProofOfWork{block, target}

	return pow
}

// 验证当前的工作量证明的有效性
func (pow *ProofOfWork) Validate() bool {

	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}
