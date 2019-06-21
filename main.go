package main

import "mychain/core"

func main(){
	bc := core.NewBlockchain()
	bc.SendData("Send 1 BTC to lzc")
	bc.SendData("Send 2 EOC to Jack")
	bc.Print()
}
