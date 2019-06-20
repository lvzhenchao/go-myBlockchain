package main

import "mychain/core"

func main(){
	bc := core.NewBlockchain()
	bc.SendData("Send 1 BTC to lzc")
	bc.SendData("Send ! EOC to Jack")
	bc.Prinf()
}
