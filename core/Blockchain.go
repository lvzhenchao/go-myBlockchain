package core

import (
	"fmt"
	"log"
)

type Blockchain struct {
	Blocks []*Block
}

//创建新的区块链
func NewBlockchain() *Blockchain{
	genesisBlock := GenerateGenesisBlock()
	blockchain := Blockchain{}
	blockchain.ApendBlock(&genesisBlock)
	return &blockchain
}

//发送数据
func (bc *Blockchain) SendData(data string) {
	preBlock := bc.Blocks[len(bc.Blocks) -1]
	newBlock := GenerateNewBlock(*preBlock, data)
	bc.ApendBlock(&newBlock)
}

//先往区块添加数据
func (bc *Blockchain) ApendBlock(newBlock *Block){
	if len(bc.Blocks) == 0 {
		bc.Blocks = append(bc.Blocks, newBlock)
		return
	}
	if isValid(*newBlock, *bc.Blocks[len(bc.Blocks) - 1]) {
		bc.Blocks = append(bc.Blocks, newBlock)
	} else {
		log.Fatal("invalid block")
	}

}


//验证
func isValid(newBlock Block, oldBlock Block) bool {
	if newBlock.Index - 1 != oldBlock.Index{
		return false
	}
	if newBlock.PrevBlockHash != oldBlock.Hash{
		return false
	}
	if calculateHash(newBlock) != newBlock.Hash {
		return false

	}
	return true
}

//打印
func(bc *Blockchain) Print(){
	for _, block := range bc.Blocks {
		fmt.Printf("Index： %d\n", block.Index)
		fmt.Printf("Prev.Hash： %s\n", block.PrevBlockHash)
		fmt.Printf("Curr.Hash： %s\n", block.Hash)
		fmt.Printf("Data： %s\n", block.Data)
		fmt.Printf("Timestamp： %d\n", block.Timestamp)
		fmt.Println()
	}
}
