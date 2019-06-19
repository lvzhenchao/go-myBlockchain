package demochain

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

//一个区块包括区块头和区块体
type Block struct {
	Index         int64  //区块编号，代表区块的位置
	Timestamp     int64  //区块时间戳
	PrevBlockHash string //上一个区块的哈希值
	Hash          string //当前区块的哈希值

	Data string //区块数据
}

func calculateHash(b Block) string {
	blockData := string(b.Index) + string(b.Timestamp) + b.PrevBlockHash + b.Data

	hashInBytes := sha256.Sum256([]byte(blockData))

	return hex.EncodeToString(hashInBytes[:])

}

func GenerateNewBlock(preBlock Block, data string) Block {
	newBlock := Block{}
	newBlock.Index = preBlock.Index + 1
	newBlock.Timestamp = time.Now().Unix()
	newBlock.PrevBlockHash = preBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock
}

//创世区块
func GenerateGenesisBlock() Block {
	preBlock := Block{}
	preBlock.Index = -1
	preBlock.Hash = ""
	return GenerateNewBlock(preBlock, "Genesis Block")

}
