package demochain

import "log"

type Blockchain struct {
	Blocks []*Block
}

func NewBlockchain() *Blockchain{
	genesisBlock := GenerateGenesisBlock()
	blockchain := Blockchain{}
	blockchain.ApendBlock(&genesisBlock)
	return &blockchain
}

func (bc *Blockchain) sendData(data string) {
	preBlock := bc.Blocks[len(bc.Blocks) -1 ]
	newBlock := GenerateNewBlock(*preBlock, data)
	bc.ApendBlock(&newBlock)
}

func (bc *Blockchain) ApendBlock(newBlock *Block){
	if (isValid(*newBlock, bc.Blocks[len(bc.Blocks) - 1])) {
		bc.Blocks = append(bc.Blocks, newBlock)
	} else {
		log.Fatal("invalid block")
	}

}

func isValid(newBlock Block, oldBlock Block){
	if newBlock.Index - 1 != oldBlock.Index{
		return false
	}
	if newBlock.PrevBlockHash != newBlock.Hash{
		return false
	}
	if calculateHash(newBlock) != newBlock.Hash {
		return false

	}
	return true
}
